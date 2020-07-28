package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/huichen/sego"
	"golang.org/x/crypto/bcrypt"
)

var Segmenter sego.Segmenter

func init() {
	// 载入词典
	// 存到内存里
	//go func() {
	//	if Segmenter.Dictionary() != nil {
	//		return
	//	}
	//	cdfilepath := "../dictionary.txt"
	//	if _, err := os.Stat("../dictionary.txt"); err == nil {
	//		cdfilepath = "../dictionary.txt"
	//	}
	//	if _, err := os.Stat("./dictionary.txt"); err == nil {
	//		cdfilepath = "./dictionary.txt"
	//	}
	//	Segmenter.LoadDictionary(cdfilepath)
	//}()
}

// 没有就拼接
func GetNORepeatKey(keepKey, needKey string) string {
	if strings.Contains(keepKey, needKey) == true {
		return keepKey
	} else {
		return keepKey + " " + needKey
	}
}

// 随机打乱[]struct数组
func RandomStruct(SlicePointer interface{}) {
	strings2 := gconv.Interfaces(SlicePointer)
	if len(strings2) <= 0 {
		return
	}

	for i := len(strings2) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		(strings2)[i], (strings2)[num] = (strings2)[num], (strings2)[i]
	}

	err := gconv.Structs(strings2, SlicePointer)
	if err != nil {
		log.Println(err.Error())
	}

	return
}

// 随机打乱[]string数组
func RandomStrings(strings_ *[]string) {
	if len(*strings_) <= 0 {
		return
	}

	for i := len(*strings_) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		(*strings_)[i], (*strings_)[num] = (*strings_)[num], (*strings_)[i]
	}
	return
}

// 是否为ＰＣ端
func IsPC(c *gin.Context) bool {
	//log.Println(strings.ToLower(r.Request.UserAgent()))
	keywords := []string{"windows phone", "mqqbrowser", "iphone", "ipod", "ipad", "android", "mobile", "blackberry", "webos", "incognito", "webmate", "bada", "nokia", "lg", "ucweb", "skyfire", "iso"}
	for i := 0; i < len(keywords); i++ {
		if strings.Contains(strings.ToLower(c.Request.UserAgent()), keywords[i]) {
			return false
		}
	}
	return true
}

// 获取当前域名
func GetDomain(c *gin.Context) string {
	host := c.Request.Host
	// 域名改为上级传入
	if domain, ok := c.GetQuery("domain"); ok && domain != "" {
		host = domain
	} else {
		return ""
	}

	//只要主域名
	hostAr := strings.Split(host, ".")
	if len(hostAr) > 2 {
		host = strings.Join(hostAr[len(hostAr)-1:], ".")
	}

	//host = strings.ReplaceAll(host, "www.", "")
	// 去掉端口
	if strings.Contains(host, ":") == true {
		return strings.Split(host, ":")[0]
	} else {
		return host
	}
}

// 获取分词 kr3为是否混合成３个词
func FenCi(word string, kr3 ...bool) []string {
	if word == "" {
		return nil
	}

	var re []string
	var ks []string
	if runtime.GOOS == "linux" {
		// x := gojieba.NewJieba()
		// defer x.Free()
		// ks = x.Cut(word, true)
	}

	if runtime.GOOS == "windows" || true {
		// 载入词典
		// var Segmenter sego.Segmenter

		if Segmenter.Dictionary() == nil { //　没加载上
			cdfilepath := "../dictionary.txt"
			if _, err := os.Stat("../dictionary.txt"); err == nil {
				cdfilepath = "../dictionary.txt"
			}
			if _, err := os.Stat("./dictionary.txt"); err == nil {
				cdfilepath = "./dictionary.txt"
			}
			Segmenter.LoadDictionary(cdfilepath)
		}

		// 分词
		text := []byte(word)
		segments := Segmenter.Segment(text)

		// 处理分词结果
		// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
		ksStr := sego.SegmentsToString(segments, true)
		fmt.Println(ksStr)

		ks_ := strings.Split(ksStr, " ")
		for _, v := range ks_ {
			if strings.Contains(v, "/") {
				ks = append(ks, strings.Split(v, "/")[0])
			}
		}
	}

	// x := gojieba.NewJieba()
	// defer x.Free()
	// re := []string{}
	// ks := x.Cut(word, true)
	// log.Println(ks)
	if len(kr3) > 0 && kr3[0] == true {
		for k1, v1 := range ks {
			for k2, v2 := range ks {
				if k2 <= k1 || len(v1) < 4 || len(v2) < 4 {
					continue
				}
				for k3, v3 := range ks {
					if k3 <= k2 || len(v3) < 4 {
						continue
					}
					if len(re) > 300 {
						goto GTO
					}
					re = append(re, v1+" "+v2+" "+v3)
				}
			}
		}
	} else {
		re = ks
	}
GTO:
	return RemoveReplicaSliceString(re)
}

// 获取远程html
func Getpage(url string) string {
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	if resp, err := client.Get(url); err != nil {
		if resp != nil {
			_ = resp.Body.Close()
		}
		log.Error("1st request failed - %s", err)
		return ""
	} else {
		result, err2 := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err2 != nil {
			log.Error("1st response read failed - %s", err2)
			return ""
		}
		// log.Printf("1st request - %s", result)
		return string(result)
	}
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

// 生成密码
func Password(len int, pwdO string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "888888"
	if pwdO != "" {
		defaultPwd = pwdO
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

// 生成32位MD5
// func MD5(text string) string{
//    ctx := md5.New()
//    ctx.Write([]byte(text))
//    return hex.EncodeToString(ctx.Sum(nil))
// }

// 生成hash密码
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证hash密码
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 废弃 模拟三元 max := If(a > b, a, b).(int)
// 改用 go get github.com/ymzuiku/hit
// func If(condition bool, trueVal, falseVal interface{}) interface{} {
// 	if condition {
// 		return trueVal
// 	}
// 	return falseVal
// }

type fn func(string) string

var (
	nameStrategyMap = map[string]fn{
		defaultNameStrategy:      snakeString,
		SnakeAcronymNameStrategy: snakeStringWithAcronym,
	}
	defaultNameStrategy      = "snakeString"
	SnakeAcronymNameStrategy = "snakeStringWithAcronym"
	nameStrategy             = defaultNameStrategy
)

// StrTo is the target string
type StrTo string

// Set string
func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

// Clear string
func (f *StrTo) Clear() {
	*f = StrTo(0x1E)
}

// Exist check string exist
func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

// Bool string to bool
func (f StrTo) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

// Float32 string to float32
func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

// Float64 string to float64
func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

// Int string to int
func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

// Int8 string to int8
func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

// Int16 string to int16
func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

// Int32 string to int32
func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

// Int64 string to int64
func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10) // octal
		if !ok {
			return v, err
		}
		return ni.Int64(), nil
	}
	return v, err
}

// Uint string to uint
func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

// Uint8 string to uint8
func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// Uint16 string to uint16
func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

// Uint32 string to uint31
func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

// Uint64 string to uint64
func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10)
		if !ok {
			return v, err
		}
		return ni.Uint64(), nil
	}
	return v, err
}

// String string to string
func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

// ToStr interface to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

// ToInt64 interface to int64
func ToInt64(value interface{}) (d int64) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		panic(fmt.Errorf("ToInt64 need numeric not `%T`", value))
	}
	return
}

func snakeStringWithAcronym(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		before := false
		after := false
		if i > 0 {
			before = s[i-1] >= 'a' && s[i-1] <= 'z'
		}
		if i+1 < num {
			after = s[i+1] >= 'a' && s[i+1] <= 'z'
		}
		if i > 0 && d >= 'A' && d <= 'Z' && (before || after) {
			data = append(data, '_')
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// snake string, XxYy to xx_yy , XxYY to xx_y_y
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// SetNameStrategy set different name strategy
func SetNameStrategy(s string) {
	if SnakeAcronymNameStrategy != s {
		nameStrategy = defaultNameStrategy
	}
	nameStrategy = s
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data[:])
}

type argString []string

// get string by index from string slice
func (a argString) Get(i int, args ...string) (r string) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

// parse time to string with location
// func timeParse(dateString, format string) (time.Time, error) {
// 	tp, err := time.ParseInLocation(format, dateString, DefaultTimeLoc)
// 	return tp, err
// }

// get pointer indirect type
func indirectType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Ptr:
		return indirectType(v.Elem())
	default:
		return v
	}
}

// 生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

const (
	base64Table = "0123456789ABCDEFGHIjklmnopqrstuvwxyzabcdefghiJKLMNOPQRSTUVWXYZ+-"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src string) string {
	return coder.EncodeToString([]byte(src))
}
func Base64Decode(src string) string {
	str, err := coder.DecodeString(src)
	strS := string(str)
	if err != nil {
		strS = ""
	}
	return strS
}
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// slice(string类型)元素去重
func RemoveReplicaSliceString(slc []string) []string {
	result := make([]string, 0)
	tempMap := make(map[string]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

// slice(int类型)元素去重
func RemoveReplicaSliceInt(slc []int) []int {
	result := make([]int, 0)
	tempMap := make(map[int]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

func JoinAny(array interface{}, sep string) string {
	argsMap := gconv.MapStrStr(array)
	ar := make([]string, len(argsMap))
	for _, v := range argsMap {
		ar = append(ar, gconv.String(v))
	}
	return strings.Join(ar, sep)
}

// 获取正在运行的函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

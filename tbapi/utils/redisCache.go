package utils

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func TestRedis() {
	_, _ = g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().Do("GET", "k")
	fmt.Println(gconv.String(v))
}

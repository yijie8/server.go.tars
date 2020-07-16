package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	"runtime"
	"ZserverWebApiNoAuth/Zserver"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func ServerInfo(c *gin.Context) {
	osDic, memDic, cpuDic, diskDic, err := _ServerInfo()
	tools.HasError(err, "", 500)
	app.Custum(c, gin.H{
		"code": 200,
		"os":   osDic,
		"mem":  memDic,
		"cpu":  cpuDic,
		"disk": diskDic,
	})
}
func ServerInfo_tars(res *Zserver.Monitor_server_res) error {
	osDic, memDic, cpuDic, diskDic, err := _ServerInfo()
	if err != nil {
		return err
	}
	res.Code = 200
	res.Cpu.CpuNum = int32(cpuDic["cpuNum"].(int))
	res.Disk.Free = int32(diskDic["free"].(int))
	res.Disk.Total = int32(diskDic["total"].(int))
	res.Mem.Total = int32(memDic["total"].(int))
	res.Mem.Free = int32(memDic["free"].(int))
	res.Mem.Usage = int32(memDic["usage"].(int))
	res.Mem.Used = int32(memDic["used"].(int))

	res.Os.GoOs = osDic["goOs"].(string)
	res.Os.Arch = osDic["arch"].(string)
	res.Os.Mem = int32(osDic["mem"].(int))
	res.Os.Compiler = osDic["compiler"].(string)
	res.Os.Version = osDic["version"].(string)
	res.Os.NumGoroutine = int32(osDic["numGoroutine"].(int))
	return nil
}

func _ServerInfo() (osDic map[string]interface{}, memDic map[string]interface{}, cpuDic map[string]interface{}, diskDic map[string]interface{}, err error) {
	osDic = make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	dis, err := disk.Usage("/")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic = make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic = make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent
	cpuDic = make(map[string]interface{}, 0)
	cpuDic["cpuNum"], err = cpu.Counts(false)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return
}

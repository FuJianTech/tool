package net

import (
	"log"
	"os"
	"runtime"
)

type version struct {
	hostname   string `json:"hostname"`
	goVersion  string	`json:"go_version"`
	sysVersion string	`json:"sys_version"`
	cpuArch    string	`json:"cpu_arch"`
	memTotal   string	`json:"mem_total"`
	cpu int	`json:"cpu"`
	cpuNum int	`json:"cpu_num"`

}

func GetHost() version {
	//var myPC version

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("获取失败")
	}
	//myPC.hostname = hostname

	sysVersion := runtime.GOOS
	cpuArch := runtime.GOARCH
	goVersion := runtime.Version()
	cpu := runtime.NumCPU()
	cpunum := runtime.GOMAXPROCS(0)
	//myPC.cpuArch = cpuArch

	myPC := version{
		hostname: hostname,
		sysVersion: sysVersion,
		goVersion: goVersion,
		cpuArch: cpuArch,
		cpu: cpu,
		cpuNum: cpunum,
	}
	//fmt.Printf("操作系统为：%s \n", goos)
	//fmt.Printf("CPU架构为：%s \n", goarch)
	//fmt.Printf("go编译版本：%s \n", s)

	return myPC
}

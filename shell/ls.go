package shell

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Ls(path string)  {
	bi, err := exec.LookPath("ls")// 查看 命令的位置
	if  err != nil {
		log.Fatal("没有找到对应的命令")
	}
	env := os.Environ()//本机环境变量
	//goos := runtime.GOOS
	arg := []string{"ls", path}      // 真正的命令
	 syscall.Exec(bi, arg, env) // 开始调用
}
func date()  {

}
func Find(year, month string)  {
	bi, err := exec.LookPath("find")// 查看 命令的位置
	if  err != nil {
		log.Fatal("没有找到对应的命令")
	}
	env := os.Environ()//本机环境变量
	//goos := runtime.GOOS
	arg := []string{"find", year+month+"*" , "-name", "*.mp4"}      // 真正的命令
//	fmt.Println(arg)
	syscall.Exec(bi, arg, env) // 开始调用
}
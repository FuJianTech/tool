/*
Copyright © 2021 author name hanzhenyong

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/Hzhenyong/go/tool/net"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看系统型号，系统架构，go版本",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}

func version()  {
	//goos := runtime.GOOS
	//goarch := runtime.GOARCH
	//s := runtime.Version()
	//fmt.Printf("操作系统为：%s \n", goos)
	//fmt.Printf("CPU架构为：%s \n", goarch)
	//fmt.Printf("go编译版本：%s \n", s)
	host := net.GetHost()
	fmt.Println(host)


}
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
	"github.com/Hzhenyong/go/tool/shell"
	"github.com/spf13/cobra"
	"os"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell  /home ",
	Short: "输入系统命令调用，暂时没有卵用",
	Long: `输入一个目录查看，目录下都有什么，`,
	Args: cobra.RangeArgs(1, 3),
	Run: func(cmd *cobra.Command, args []string) {
		//ls()
		if len(args) == 0 {
			fmt.Println("请输入一个命令。 eg: ls")
			os.Exit(0)
		}
		switch args[0] {
		case "ls":
			if len(args) > 1 {
				shell.Ls(args[1])
			}else {
				fmt.Printf("请输入要查看的路径. \n")
			}
		case "find":
			if len(args) > 2 {
				shell.Find(args[1], args[2])
			}
			fmt.Printf("请输入年份和月份. \n")
		default:
			fmt.Printf("此命令暂时未支持. \n")
		}
		
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

}


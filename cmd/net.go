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
	"os"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "集成网络相关的工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net called")

		if len(args) == 0 {
			fmt.Println("请输入一个命令。 eg: ")
			os.Exit(0)
		}
		switch args[0] {
		case "curl":
			net.Curl(args[1])
		case "ping":
			net.Ping(args)
		default:
			fmt.Printf("暂未支持")
		}
	},
}

func init() {
	rootCmd.AddCommand(netCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

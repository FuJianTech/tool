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

	"github.com/spf13/cobra"
)

var name string
// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "数据库相关的操作工具",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sql called")
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)
	rootCmd.Flags().StringVarP(&name, name, "n", "", "")



}

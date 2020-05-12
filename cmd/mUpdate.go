/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fso/internal/npmrepo"
	"fso/internal/updater"
	"log"

	"github.com/spf13/cobra"
)

var moduleNameFlag string

// mUpdateCmd represents the mUpdate command
var mUpdateCmd = &cobra.Command{
	Use:   "mUpdate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		npmrepo.UpdateNodeModules()

		config := GetConfig()

		u := updater.NewUpdater("Module")

		if moduleNameFlag == "all" {
			for moduleName := range config.Modules {
				fmt.Println(moduleName)
				if err := u.Copy(moduleName, config.Modules[moduleName].PathFrom, config.Modules[moduleName].PathIn); err == nil {
					fmt.Println("Обновление модуля прошло успешно")
					fmt.Println()
				} else {
					log.Fatal(err)
				}
			}
		} else {
			fmt.Println(moduleNameFlag)
			if err := u.Copy(moduleNameFlag, config.Modules[moduleNameFlag].PathFrom, config.Modules[moduleNameFlag].PathIn); err == nil {
				fmt.Println("Обновление модуля прошло успешно")
			} else {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mUpdateCmd)

	mUpdateCmd.Flags().StringVarP(&moduleNameFlag, "module-name", "m", "", "Имя обновляемого модуля, или all, если нужно обновить все модули")
	mUpdateCmd.MarkFlagRequired("module-name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

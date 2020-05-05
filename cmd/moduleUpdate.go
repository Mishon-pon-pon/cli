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
	"fso/internal/modules"
	"fso/internal/npmrepo"
	"log"

	"github.com/spf13/cobra"
)

var moduleNameFlag string

// moduleUpdateCmd represents the moduleUpdate command
var moduleUpdateCmd = &cobra.Command{
	Use:   "module-update",
	Short: "Обновляет модуль проекта. Флаг -m обязателен",
	Long: `
	fso module-update -m имя_модуля/all

	Вы вводите команду fso module-update -m имя_модуля.
	cli обновляет npm пакеты(указанные в package.json), затем идет в node_modules по имя_модуля.
	Берет от туда файлы и переносит их в папку которую вы добавили в fso_config.json в поле pathIn.
	Если указать после флага -m значение all, то обновятся все модули указанные в fso_config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		npmrepo.UpdateNodeModules()

		m := modules.NewModule()

		config := GetConfig()

		if moduleNameFlag == "all" {
			for moduleName := range config.Modules {
				fmt.Println(moduleName)
				if err := m.CopyModule(moduleName, config.Modules); err == nil {
					fmt.Println("Обновление модуля прошло успешно")
					fmt.Println()
				} else {
					log.Fatal(err)
				}
			}
		} else {
			fmt.Println(moduleNameFlag)
			if err := m.CopyModule(moduleNameFlag, config.Modules); err == nil {
				fmt.Println("Обновление модуля прошло успешно")
			} else {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(moduleUpdateCmd)

	moduleUpdateCmd.Flags().StringVarP(&moduleNameFlag, "module-name", "m", "", "Имя обновляемого модуля, или all, если нужно обновить все модули")
	moduleUpdateCmd.MarkFlagRequired("module-name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

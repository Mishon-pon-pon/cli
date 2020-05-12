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

var serviceNameFlag string

// moduleUpdateCmd represents the moduleUpdate command
var serviceUpdateCmd = &cobra.Command{
	Use:   "s-update",
	Short: "Обновляет модуль проекта. Флаг -n обязателен",
	Long: `
	fso module-update -m имя_модуля/all

	Вы вводите команду fso module-update -m имя_модуля.
	cli обновляет npm пакеты(указанные в package.json), затем идет в node_modules по имя_модуля.
	Берет от туда файлы и переносит их в папку которую вы добавили в fso_config.json в поле pathIn.
	Если указать после флага -m значение all, то обновятся все модули указанные в fso_config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		npmrepo.UpdateNodeModules()

		config := GetConfig()

		u := updater.NewUpdater("Service")

		if serviceNameFlag == "all" {
			for moduleName := range config.Services {
				fmt.Println(moduleName)
				if err := u.Copy(moduleName, config.Services[moduleName].PathFrom, config.Services[moduleName].PathIn); err == nil {
					fmt.Println("Обновление модуля прошло успешно")
					fmt.Println()
				} else {
					log.Fatal(err)
				}
			}
		} else {
			fmt.Println(serviceNameFlag)
			if err := u.Copy(serviceNameFlag, config.Services[serviceNameFlag].PathFrom, config.Services[serviceNameFlag].PathIn); err == nil {
				fmt.Println("Обновление модуля прошло успешно")
			} else {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(serviceUpdateCmd)

	serviceUpdateCmd.Flags().StringVarP(&serviceNameFlag, "service-name", "s", "all", "Имя обновляемого сервиса, или all, если нужно обновить все модули")
	serviceUpdateCmd.MarkFlagRequired("service-name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

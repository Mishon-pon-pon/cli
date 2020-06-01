/*
Updateright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a Update of the License at

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

	"github.com/spf13/cobra"
)

var moduleNameFlag string

// mUpdateCmd represents the mUpdate command
var mUpdateCmd = &cobra.Command{
	Use:   "m-update",
	Short: "обновляет модуль проекта. Флаг -m обязателен.",
	Long: `
	fso m-update -m имя_модуля/all

	Вы вводите команду fso m-update -m имя_модуля.
	cli обновляет npm пакеты(указанные в package.json), затем идет в node_modules по имя_модуля.
	Берет от туда файлы и переносит их в папку которую вы добавили в fso_config.json в поле pathIn.
	Если указать после флага -m значение all, то обновятся все модули указанные в fso_config.json
	
	При обновлении модуля, программа мёрджит содержимое папки указанной в pathIn модуля. Если файл
	уже существовал программа обновит его содержимое, если файла небыло, то добавит. Если файл есть
	локально, но его нет в pathFrom то программа ничего с ним не сделает.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config := GetConfig()

		success := npmrepo.UpdateNodeModules(config.NpmRepository)
		if success != true {
			return
		}

		u := updater.NewUpdater("Module")

		if moduleNameFlag == "all" {
			for moduleName := range config.Modules {
				fmt.Println(moduleName)
				_, ok := config.Modules[moduleName]
				if ok {
					if err := u.Update(moduleName, config.Modules[moduleName].PathFrom, config.Modules[moduleName].PathIn); err == nil {
						fmt.Println("Обновление модуля прошло успешно")
						fmt.Println()
					} else {
						fmt.Println(err)
						fmt.Println()
					}
				} else {
					fmt.Print("нет такого модуля")
				}

			}
		} else {
			fmt.Println(moduleNameFlag)
			if err := u.Update(moduleNameFlag, config.Modules[moduleNameFlag].PathFrom, config.Modules[moduleNameFlag].PathIn); err == nil {
				fmt.Println("Обновление модуля прошло успешно")
			} else {
				fmt.Println(err)
				fmt.Println()
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

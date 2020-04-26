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
	"log"

	"github.com/spf13/cobra"
)

var moduleName string

// moduleUpdateCmd represents the moduleUpdate command
var moduleUpdateCmd = &cobra.Command{
	Use:   "module-update",
	Short: "Обновляет модуль проекта. Флаг -m обязателен",
	Long: `Вы вводите команду fso module-update -m имя_модуля.
	cli обновляет npm пакет, затем идет в этот пакет по имя_модуля.
	Берет от туда файлы и переносит их в папку которую вы добавили в fso_config в поле pathIn`,
	Run: func(cmd *cobra.Command, args []string) {
		m := modules.NewModule()
		m.UpdateNodeModules()

		fmt.Println("node_modules обновлены")
		fmt.Println("\nОбновление модуля проекта...")

		config := GetConfig()
		if err := m.FindModule(moduleName, config.Modules); err == nil {
			m.DeleteModule(moduleName, config.Modules)
			err := m.CopyModule(config.Modules[moduleName].PathFrom, config.Modules[moduleName].PathIn, moduleName, config.Modules)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\nОбновление модуля прошло успешно")
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(moduleUpdateCmd)

	moduleUpdateCmd.Flags().StringVarP(&moduleName, "module-name", "m", "", "module name")
	moduleUpdateCmd.MarkFlagRequired("module-name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

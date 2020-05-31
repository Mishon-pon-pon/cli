/*.Updateright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a.Update of the License at

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
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var serviceNameFlag string

// moduleUpdateCmd represents the moduleUpdate command
var serviceUpdateCmd = &cobra.Command{
	Use:   "s-update",
	Short: "Обновляет сервис проекта. Флаг -s обязателен.",
	Long: `
	fso s-update -s имя_сервиса/all

	Вы вводите команду fso s-update -s имя_сервиса.
	cli обновляет npm пакеты(указанные в package.json), затем идет в node_modules по имя_сервиса.
	Берет от туда файлы и переносит их в папку которую вы добавили в fso_config.json в поле pathIn.
	Если указать после флага -s значение all, то обновятся все модули указанные в fso_config.json
	
	При обновлении сервиса, программа зачищает папку которую вы указали в pathIn. Удаляется всё, кроме
	файлов с расширением .config
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config := GetConfig()

		success := npmrepo.UpdateNodeModules(config.NpmRepository)
		if success != true {
			return
		}

		u := updater.NewUpdater("Service")

		if serviceNameFlag == "all" {
			for moduleName := range config.Services {
				fmt.Println(moduleName)
				if err := u.Update(moduleName, config.Services[moduleName].PathFrom, config.Services[moduleName].PathIn); err == nil {
					fmt.Println("Обновление сервиса прошло успешно")
					fmt.Println()
				} else {
					log.Fatal(err)
				}
			}
		} else {
			fmt.Println(serviceNameFlag)
			_, ok := config.Services[serviceNameFlag]
			if ok {
				s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
				s.Start()
				if err := u.Update(serviceNameFlag, config.Services[serviceNameFlag].PathFrom, config.Services[serviceNameFlag].PathIn); err == nil {
					s.Stop()
					fmt.Println("Обновление сервиса прошло успешно")
				} else {
					log.Fatal(err)
				}

			} else {
				fmt.Print("нет такого сервиса")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(serviceUpdateCmd)

	serviceUpdateCmd.Flags().StringVarP(&serviceNameFlag, "service-name", "s", "", "Имя обновляемого сервиса, или all, если нужно обновить все сервисы")
	serviceUpdateCmd.MarkFlagRequired("service-name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

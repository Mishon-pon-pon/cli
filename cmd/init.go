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
	"fso/internal/init/initjson"
	"fso/internal/version"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Создает файл fso_config.json в котором нужно прописать настройки",
	Long: `
	Создает файл fso_config.json в котором нужно прописать настройки
	Основа там есть, нужно только прописать конкретику.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("init called")
		initConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfiguration() {
	if _, err := os.Stat("fso_config.json"); os.IsNotExist(err) {
		fmt.Println(version.GetVersion("banner3-D"))
		file, err := os.Create("fso_config.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file.WriteString(initjson.DefaultJSON)

		fmt.Println("\nИнициализация прошла успешно! \nСоздан файл fso_config.json. \nЕго необходимо настроить.")
	} else {
		// conf := GetConfig()
		// newJSON := strings.Replace(initjson.DefaultJSON, "server=IP_сервера;user id=Имя_Пользователя;password=Пароль;database=База_Данных", conf.DataBase.DataBaseURL, -1)
		// newJSON = strings.Replace(newJSON, "mssql", conf.DataBase.DBManager, -1)
		// newJSON = strings.Replace(newJSON, "полный путь до удаленного репозитория", conf.Repository.RemotePath, -1)
		// newJSON = strings.Replace(newJSON, "полный путь до dev репозитория", conf.Repository.DevPath, -1)
		// newJSON = strings.Replace(newJSON, "полный путь до test репозитория", conf.Repository.TestPath, -1)

		// file, err := os.Create("fso_config.json")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file.Close()

		// file.WriteString(newJSON)
		fmt.Println("\nВ этой папке уже была инициализация. Смотрите файл fso_config.json")
	}

}

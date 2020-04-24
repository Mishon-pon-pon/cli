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
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
		UpdateNodeModules()

		fmt.Println("node_modules обновлены")
		fmt.Println("\nОбновление модуля проекта...")

		config := GetConfig()
		if _, ok := config.Modules[moduleName]; ok {
			if config.Modules[moduleName].PathFrom == "" {
				found := false
				err := filepath.Walk("node_modules/", func(path string, info os.FileInfo, err error) error {
					if err != nil {
						fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
						return err
					}
					path = strings.Replace(path, `\`, "/", -1)
					if info.IsDir() {
						if strings.Contains(path, moduleName) && !found {
							config.Modules[moduleName].PathFrom = path
							found = true
						}
					}
					return nil
				})
				if err != nil {
					log.Fatal(err)
				}
			}
			err := Copy(config.Modules[moduleName].PathFrom, config.Modules[moduleName].PathIn)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\nОбновление модуля прошло успешно")
		} else {
			fmt.Println("\nНет такого модуля или он не указан в fso_config.json")
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

// UpdateNodeModules ...
func UpdateNodeModules() {
	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Обновление node_modules...")
	c := exec.Command(ps, "npm update")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	c.Run()
	// o, err := c.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if string(o) != "" {
	// 	fmt.Println(string(o))
	// } else {
	// 	// fmt.Println("Нет изменений в текущей версии модуля")
	// }
}

// CopyError ...
type CopyError struct {
	err error
	msg string
}

func (c *CopyError) Error() string { return c.msg }

// Copy ...
func Copy(from string, in string) error {
	return filepath.Walk(from, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, `\`, "/", -1)
		if !info.IsDir() {
			fileFrom, err := os.Open(path)
			defer fileFrom.Close()
			if err != nil {
				return &CopyError{err: err, msg: "\nНе удалось открыть файл " + path}
			}

			path = strings.Replace(path, from, in, -1)

			fileIn, err := os.Create(path)
			defer fileIn.Close()
			if err != nil {
				return &CopyError{err: err, msg: "\nНе удалось создать файл " + path}
			}
			_, err = io.Copy(fileIn, fileFrom)
			if err != nil {
				return &CopyError{err: err, msg: "\nНе удалось скопировать файл " + path}
			}
		} else {
			path = strings.Replace(path, from, in, -1)
			os.Mkdir(path, 0666)
		}

		return nil
	})
}

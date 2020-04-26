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
	"errors"
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
		m := NewModule()
		m.UpdateNodeModules()

		fmt.Println("node_modules обновлены")
		fmt.Println("\nОбновление модуля проекта...")

		config := GetConfig()
		if err := m.FindModule(moduleName, config); err == nil {
			m.DeleteModule(moduleName, config)
			err := m.CopyModule(config.Modules[moduleName].PathFrom, config.Modules[moduleName].PathIn)
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

// Module ...
type Module struct {
	ConfigFiles map[string][]string
}

func NewModule() *Module {
	return &Module{}
}

// UpdateNodeModules ...
func (m *Module) UpdateNodeModules() {
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

// FindModule ...
func (m *Module) FindModule(moduleName string, config *Config) error {
	if _, ok := config.Modules[moduleName]; ok {
		if config.Modules[moduleName].PathFrom == "" {
			found := false
			return filepath.Walk("node_modules/", func(path string, info os.FileInfo, err error) error {
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
		}
		return nil
	}
	return errors.New("\nНет такого модуля, либо он не указан в fso_config.json")
}

// CopyError ...
type CopyError struct {
	err error
	msg string
}

func (c *CopyError) Error() string { return c.msg }

func copyFile(from, in string) error {
	fileFrom, err := os.Open(from)
	defer fileFrom.Close()
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось открыть файл " + from}
	}

	fileIn, err := os.Create(in)
	defer fileIn.Close()
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось создать файл " + from}
	}
	_, err = io.Copy(fileIn, fileFrom)
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось скопировать файл " + from}
	}
	return nil
}

// CopyModule ...
func (m *Module) CopyModule(from string, in string) error {

	return filepath.Walk(from, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, `\`, "/", -1)
		fromFile := path
		inFile := strings.Replace(path, from, in, -1)
		if !info.IsDir() {
			if !strings.Contains(path, ".config") {
				err = copyFile(fromFile, inFile)
				if err != nil {
					return err
				}
			} else {
				toCopy := true
				confFrom := strings.Replace(path, from, "", -1)
				for i := 0; i < len(m.ConfigFiles[moduleName]); i++ {
					confIn := m.ConfigFiles[moduleName][i]
					var comfirmPath string
					if len(confIn) > len(confFrom) {
						comfirmPath = strings.Replace(confIn, confFrom, "", -1)
					} else {
						comfirmPath = strings.Replace(confFrom, confIn, "", -1)
					}
					if comfirmPath == `` || comfirmPath == `/` || comfirmPath == `\` {
						toCopy = false
					}
				}
				if toCopy {
					copyFile(fromFile, inFile)
				}

			}

		} else {
			path = strings.Replace(path, from, in, -1)
			os.Mkdir(path, 0666)
		}

		return nil
	})
}

// DeleteModule ...
func (m *Module) DeleteModule(moduleName string, config *Config) error {
	m.ConfigFiles = map[string][]string{}
	in := config.Modules[moduleName].PathIn
	if in != "" {
		removeDir := make(map[string]bool)
		allDir := []string{}
		filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
			path = strings.Replace(path, `\`, "/", -1)
			allDir = append(allDir, path)
			if strings.Contains(path, ".config") {
				confIn := strings.Replace(path, in, "", -1)
				m.ConfigFiles[moduleName] = append(m.ConfigFiles[moduleName], confIn)
				pathArr := strings.Split(path, "/")
				for _, v := range pathArr {
					removeDir[v] = true
				}
			}
			return nil
		})
		for _, path := range allDir {
			pathDirArr := strings.Split(path, "/")
			for _, dir := range pathDirArr {
				if dir != "" {
					if _, ok := removeDir[dir]; ok {

					} else {
						os.RemoveAll(path)
					}
				}

			}
		}
	}

	return nil
}

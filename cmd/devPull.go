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
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// devPullCmd represents the devPull command
var devPullCmd = &cobra.Command{
	Use:   "dev-pull",
	Short: "Эта команда делает пулл на дев стенде",
	Long: `
	Для того чтобы команда сделала pull на дев стенде нужно, 
	чтобы в файле fso_config.json в поле Repository.devPath
	был прописан путь до папки с гитом на дев стенде. 
	Так как это json и так как это сетевой путь windows
	то нужно экранировать симовлы. Например:

	"Repository": {
		"devPath": "\\\\10.1.12.87\\c$\\Foresight\\mtk17"
	}

`,
	Run: func(cmd *cobra.Command, args []string) {
		DevPull()
	},
}

func init() {
	rootCmd.AddCommand(devPullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devPullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devPullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// DevPull ...
func DevPull() {
	ps, _ := exec.LookPath("powershell.exe")

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	conf := GetConfig()
	err = os.Chdir(conf.GitRepository.DevPath)
	if err != nil {
		log.Fatal(err)
	}

	p, _ := os.Getwd()
	fmt.Println("Обновляем экземпляр репозитория который находится по адресу", p)
	cmd := exec.Command(ps, "git pull")

	fmt.Println("Основной репозиторий:")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	var userAccept string
	fmt.Print(`Хотите обновить субмодуль? y(да)\n(нет)`)
	fmt.Fscan(os.Stdin, &userAccept)

	if userAccept == "y" {
		cmd = exec.Command(ps, "git submodule update --init")
		fmt.Println("Субмодуль:")
		out, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		if string(out) == "" {
			fmt.Println("Нет изменений в субмодуле.")
		} else {
			fmt.Println(string(out))
		}
	} else {

	}

	fmt.Println("\nЗавершено")

	err = os.Chdir(path)
	if err != nil {
		log.Fatal(err)
	}
}

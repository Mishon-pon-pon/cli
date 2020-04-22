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
	"encoding/json"
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
	чтобы в файле fso_config.json в поле dev_repository.path
	был прописан путь до папки с гитом на дев стенде. 
	Так как это json и так как это сетевой путь 
	то нужно экранировать симовлы. Например:

	"dev_repository": {
		"path": "\\\\10.1.12.87\\c$\\Foresight\\mtk17"
	}

`,
	Run: func(cmd *cobra.Command, args []string) {
		DevPull()
	},
}

type Repository struct {
	Dev_repository struct {
		Path string `json:"path"`
	}
}

var dev_repository Repository

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

func DevPull() {
	ps, _ := exec.LookPath("powershell.exe")
	var userAccept string

	config, err := os.Open(`fso_config.json`)
	if err != nil {
		log.Fatal(err)
	}

	defer config.Close()

	decoder := json.NewDecoder(config)
	err = decoder.Decode(&dev_repository)
	if err != nil {
		log.Fatal(err)
	}
	comand := "cd " + dev_repository.Dev_repository.Path
	fmt.Println(comand)

	cmd := exec.Command(ps, "cd", dev_repository.Dev_repository.Path, "\ngit pull")

	fmt.Println("Основной репозиторий:")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	fmt.Print(`Хотите обновить субмодуль? y(да)\n(нет)`)
	fmt.Fscan(os.Stdin, &userAccept)

	if userAccept == "y" {
		cmd = exec.Command(ps, `
		cd \\10.1.12.87\c$\Foresight\kzk2018
		git submodule update --init
	`)
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
	}
}

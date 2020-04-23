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
	"strings"

	"github.com/spf13/cobra"
)

// fsoUpdateCmd represents the fsoUpdate command
var fsoUpdateCmd = &cobra.Command{
	Use:   "fso-update",
	Short: "Обновляет эту программу, если есть обновления.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Обновление...")
		FSOUpdate()
	},
}

func init() {
	rootCmd.AddCommand(fsoUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fsoUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fsoUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// FSOUpdate ...
func FSOUpdate() {
	cmd, _ := exec.LookPath(`powershell.exe`)
	dir, _ := os.Executable()
	arr := strings.Split(dir, `\`)
	arr = arr[:len(arr)-1]
	dir = strings.Join(arr, "/")
	comand := exec.Command(cmd, "cd", dir, "\ngit pull")
	out, err := comand.Output()
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(out), "Already up to date") {
		out = []byte("У вас самая свежая версия программы.")
	} else {
		out = []byte("Программа обновлена.")
	}
	fmt.Println(string(out))

}

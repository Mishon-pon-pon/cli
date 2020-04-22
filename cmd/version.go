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

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Версию показывает и картинку",
	Long:  `будет описание, но поже`,
	Run: func(cmd *cobra.Command, args []string) {
		hello := `
    ___________ ____ 
   / ____/ ___// __ \
  / /_   \__ \/ / / /
 / __/  ___/ / /_/ / 
/_/    /____/\____/ 
    ____  _______    __________    ____  ____  __________  _____
   / __ \/ ____/ |  / / ____/ /   / __ \/ __ \/ ____/ __ \/ ___/
  / / / / __/  | | / / __/ / /   / / / / /_/ / __/ / /_/ /\__ \ 
 / /_/ / /___  | |/ / /___/ /___/ /_/ / ____/ /___/ _, _/___/ / 
/_____/_____/  |___/_____/_____/\____/_/   /_____/_/ |_|/____/  
        ____   ____   ___
 _   __/ __ \ / __ \ <  /
| | / / / / // / / / / / 
| |/ / /_/ // /_/ / / /  
|___/\____(_)____(_)_/  
`
		fmt.Println(hello)

		fmt.Println("v.0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

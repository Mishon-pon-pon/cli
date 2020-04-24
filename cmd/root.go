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
	"fso/internal/db"
	"fso/internal/modules"
	"fso/internal/repo"
	"fso/internal/version"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var config *Config

// Config ...
type Config struct {
	DataBase   *db.Config
	Repository *repo.Config
	Modules    map[string]*modules.Config
}

// GetConfig ...
func GetConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{
		DataBase:   db.NewConfig(),
		Repository: repo.NewConfig(),
	}
	return config
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fso",
	Short: "CLI для FSO прогеров",
	Long: version.Version + `

	CLI для FSO прогеров.

	Эта потрясающая программа пока не делает ничего, но что бУдет делать!
	
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is fso_config.json)")

	// if cfgFile == "" {
	// 	cfgFile = "fso_config.json"
	// }
	configs := GetConfig()
	file, err := os.Open("fso_config.json")
	if err != nil {
		fmt.Println("fso не проинициализированно. Выполните fso init")
	} else {
		jsonConf := json.NewDecoder(file)
		err = jsonConf.Decode(configs)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// // Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// // Search config in home directory with name ".fso" (without extension).
		// viper.AddConfigPath(home)
		// viper.SetConfigName(".fso")
		viper.SetConfigFile("fso_config.json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Используется конфигаруционный файл:", viper.ConfigFileUsed())
	}
}
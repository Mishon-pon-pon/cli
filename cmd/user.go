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
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "users",
	Short: "Предоставляет пароли и логины пользователей.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		User()
		fmt.Println("user called")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// User ...
func User() {
	type user struct {
		Account  string
		Password string
	}

	config := NewConfig()
	database, err := sql.Open(config.DataBase.DBManager, config.DataBase.DataBaseURL)
	rows, err := database.Query(`SELECT u.Password,
									    a.Account
									FROM Account a 
									JOIN [User] u ON u.UserId = a.AccountId
									WHERE u.[Password] is not NULL 
									AND a.Account is not NULL`,
	)
	if err != nil {
		log.Fatal("Не удалось сделать запрос к базе данных", "\nОшибка:", err)
	}
	defer rows.Close()

	users := make([]*user, 0)
	for rows.Next() {
		u := new(user)
		rows.Scan(&u.Password, &u.Account)

		users = append(users, u)
	}
	for _, user := range users {
		fmt.Printf("Login %s | Password %s\n", user.Account, user.Password)
	}
}

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
	"encoding/json"
	"fmt"
	"fso/internal/db"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/cobra"
)

var PageElementName string

// mpemCmd represents the mpem command
var mpemCmd = &cobra.Command{
	Use:   "mpem",
	Short: "Пока в разработке. Будет делать миграцию js кода PageElement в файл на диске",
	Long:  `Пока в разработке команда`,
	Run: func(cmd *cobra.Command, args []string) {
		// metaPageElementMigrate()
		fmt.Println("Пока в разработке")
	},
}

func init() {
	rootCmd.AddCommand(mpemCmd)
	rootCmd.PersistentFlags().StringVarP(&PageElementName, "page-element", "p", "", "page element name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mpemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mpemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func metaPageElementMigrate() {
	var database db.DB
	file, err := os.Open(`fso_config.json`)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&database)
	if err != nil {
		log.Fatal(err)
	}

	type Test struct {
		PageElementId string
		DirName       string
		FileName      string
		Content       string
	}

	db, err := sql.Open("mssql", database.Connection.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r, err := db.Query(`
		SELECT cast(PageElementId as nvarchar(255)) as PageElementId,  mp.Name as Dir, mpe.Name + '.js' as [File], mpe.Content
		FROM MetaPageElement mpe
		JOIN MetaPage mp ON mp.PageId = mpe.PageId
		WHERE mpe.[Name] = ?
	`, PageElementName)
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	output.Close()

	for r.Next() {
		var x Test

		r.Scan(&x.PageElementId, &x.DirName, &x.FileName, &x.Content)

		var str = regexp.MustCompile(`<script>(.|\n|\r\n)*?</script>`)
		scriptPath := `<sciprt type="text/javascript" src="/asyst/libs/page/` + x.DirName + `/` + x.FileName + `" />`
		scriptTag := string(str.Find([]byte(x.Content)))
		fmt.Println("ScriptTag", scriptTag)
		fmt.Println("ScriptPath", scriptPath)
		fmt.Println(x.Content)
		if scriptTag != "" {
			newContent := strings.Replace(x.Content, scriptTag, scriptPath, -1)
			fmt.Println(newContent)

			output, err := os.OpenFile("output.html", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				log.Fatal(err)
			}
			output.WriteString(x.PageElementId + newContent + `

			<!--====================================================================================================-->

			`)
			output.Close()

			tmpD, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			err = os.Mkdir(tmpD+`\`+x.DirName, 0777)
			file, err := os.Create(tmpD + `\` + x.DirName + `\` + x.FileName)
			if err != nil {
				log.Fatal(err)
			}
			scriptJs := strings.Replace(scriptTag, "<script>", "", -1)
			scriptJs = strings.Replace(scriptJs, "</script>", "", -1)
			defer file.Close()
			file.WriteString(scriptJs)

			// db.Exec(`
			// 	UPDATE MetaPageElement
			// 	SET Content = ?
			// 	WHERE PageElementId = ?
			// `, newContent, x.PageElementId)
		}

	}
	fmt.Println("done!")
}

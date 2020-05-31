package npmrepo

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// stdOutput - создан для того чтобы перенаправить вывод в нашу переменную
type stdOutput struct {
	msg string
}

func (t *stdOutput) Write(p []byte) (int, error) {
	t.msg = string(p)
	return 0, nil
}

type stdInput struct {
}

type stdError struct {
	errMsg string
}

func (e *stdError) Write(m []byte) (int, error) {
	e.errMsg = string(m)
	return -1, nil
}

// Config - npm repository configuration
type Config struct {
	Registry string `json:"registry"`
}

// NewConfig - constructor npm repository configuration
func NewConfig() *Config {
	return &Config{}
}

// UpdateNodeModules ...
func UpdateNodeModules(config *Config) bool {
	// var input *stdInput = &stdInput{}
	var output *stdOutput = &stdOutput{}
	var errorput *stdError = &stdError{}

	var userInput string
	// cmd, err := exec.LookPath("cmd")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c := exec.Command(cmd, "npm config get registry")
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr
	// c.Run()
	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		log.Fatal(err)
	}

	/*Смотрим какой npm регистри установлен в данный момент у пользователя*/
	c := exec.Command(ps, "npm config get registry")
	c.Stdout = output
	c.Run()

	/*сравниваем регистри пользователя с регистри подтянутым из конфига*/
	if strings.TrimSpace(config.Registry) == strings.TrimSpace(output.msg) {
		/*просто обновляем npm пакеты*/
		fmt.Println("Обновление node_modules...")
		c = exec.Command(ps, "npm update")
		c.Stdout = os.Stdout
		c.Stderr = errorput

		c.Run()

		if errorput.errMsg != "" {
			fmt.Println(errorput.errMsg)
			fmt.Println("node_modules не обновлены, продолжить?...")
			fmt.Fscan(os.Stdin, &userInput)
			if strings.ToLower(userInput) != "y" {
				return false
			}
		}

		fmt.Println("node_modules обновлены\n")

	} else { // если текущий регистри пользователя не совпадает с регистри из конфига
		/*смотрим в c:/Users/пользователь/.npmrc и проверяем ходил ли пользователь на этот регистри*/
		homePath := os.Getenv("HOMEPATH")
		// homePath = strings.Replace(homePath, "\\", "/", -1)
		// fmt.Println("C:" + homePath + "\\.npmrc")
		file, err := os.Open("C:" + homePath + "\\.npmrc")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		b := make([]byte, 64)
		var fileContent string

		for {
			n, err := file.Read(b)
			if err == io.EOF {
				break
			}
			fileContent += string(b[:n])
		}
		/*если в файле есть строка совпадающая с регистри из конфига, то устанавливем его в качестве текущего*/
		if strings.Contains(fileContent, strings.Replace(config.Registry, "http:", "", -1)) {
			c = exec.Command(ps, "npm config set registry "+config.Registry)
			c.Run()
			fmt.Println("Обновление node_modules...")
			c = exec.Command(ps, "npm update")
			c.Stdout = os.Stdout
			c.Stderr = errorput

			c.Run()

			if errorput.errMsg != "" {
				fmt.Println(errorput.errMsg)
				fmt.Println("node_modules не обновлены, продолжить?...(y/n)")
				fmt.Fscan(os.Stdin, &userInput)
				if strings.ToLower(userInput) != "y" {
					return false
				}
			}
			fmt.Println("node_modules обновлены\n")
		} else { // если в файле нет строки регистри из конфига то добавляем его пользователю
			fmt.Println("Вы еще не разу не логинились в npm-репозитории " + config.Registry)
			fmt.Println("Сделайте это.")
			example := `
			Пример:
			Username: ваше имя в корпаротивной системе.
			Password: пароль с которым вы входите в виндоуз.
			Email: корпаративный емаил(ваш_логин@pmpractice.ru).
			`
			fmt.Println(example)
			c = exec.Command(ps, "npm adduser --registry="+config.Registry+" --always-auth")
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = errorput
			c.Run()

			if errorput.errMsg != "" {
				fmt.Println(errorput.errMsg)
				fmt.Println("node_modules не обновлены, продолжить?...(y/n)")
				fmt.Fscan(os.Stdin, &userInput)
				if strings.ToLower(userInput) != "y" {
					return false
				}
			}
		}
		/*возвращаем регистри пользователю что был до выполения функции*/
		c = exec.Command(ps, "npm config set registry "+output.msg)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	}

	// o, err := c.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if string(o) != "" {
	// 	fmt.Println(string(o))
	// } else {
	// 	// fmt.Println("Нет изменений в текущей версии модуля")
	// }
	return true
}

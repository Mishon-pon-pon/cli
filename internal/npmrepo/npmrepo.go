package npmrepo

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type stdOutput struct {
	currentNpmRegistry string
}

func (t *stdOutput) Write(p []byte) (int, error) {
	t.currentNpmRegistry = string(p)
	return 0, nil
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
func UpdateNodeModules(config *Config) {
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
	c := exec.Command(ps, "npm config get registry")
	out := &stdOutput{}
	c.Stdout = out
	c.Run()

	if strings.TrimSpace(config.Registry) == strings.TrimSpace(out.currentNpmRegistry) {
		fmt.Println("Обновление node_modules...")
		c = exec.Command(ps, "npm update")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		c.Run()
		fmt.Println("node_modules обновлены\n")
	} else {
		homePath := os.Getenv("HOMEPATH")
		file, err := os.Open(homePath + "/.npmrc")
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
		if strings.Contains(fileContent, strings.Replace(config.Registry, "http:", "", -1)) {
			c = exec.Command(ps, "npm config set registry "+config.Registry)
			c.Run()
			fmt.Println("Обновление node_modules...")
			c = exec.Command(ps, "npm update")
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			c.Run()
			fmt.Println("node_modules обновлены\n")
		} else {
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
			c.Stderr = os.Stderr
			c.Run()
		}
		c = exec.Command(ps, "npm config set registry "+out.currentNpmRegistry)
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
}

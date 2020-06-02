package npmrepo

import (
	"errors"
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
	if config.Registry == "" || strings.Contains(config.Registry, "вашего") {
		log.Fatal(errors.New("Вы не указали адрес npm репозитория(файл fso_config.json)."))
	}

	var output *stdOutput = &stdOutput{}
	var npmrc *os.File
	defer npmrc.Close()

	homePath := os.Getenv("USERPROFILE")
	_, err := os.Stat(homePath + "\\.npmrc")
	if err != nil {
		npmrc, err = os.Create(homePath + "\\.npmrc")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		npmrc, err = os.Open(homePath + "\\.npmrc")
		if err != nil {
			log.Fatal(err)
		}
	}

	buf := make([]byte, 64)
	var npmrcContent string

	for {
		bundle, err := npmrc.Read(buf)
		if err == io.EOF {
			break
		}
		npmrcContent += string(buf[:bundle])
	}

	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		fmt.Println("У вас не установлен powershell")
		log.Fatal(err)
	}
	c := exec.Command(ps, "npm config get registry")
	c.Stdout = output
	c.Run()
	userRegistry := output.msg

	if strings.Contains(npmrcContent, strings.Replace(config.Registry, "http:", "", -1)+":_authToken") {
		c := exec.Command(ps, "npm config set registry="+config.Registry)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		c = exec.Command(ps, "npm update")
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		c = exec.Command(ps, "npm config delete registry")
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		if userRegistry == "" {
			c = exec.Command(ps, "npm config delete registry")
		} else {
			c = exec.Command(ps, "npm config set registry="+userRegistry)
		}
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		return true
	} else {
		fmt.Println("Вы не залогированы в", config.Registry)
		fmt.Println("Для начала залогинтесь. Далее повторите попытку.")
		example := `
			Пример:
			Username: ваше имя в корпаротивной системе.
			Password: пароль с которым вы входите в виндоуз.
			Email: корпаративный емаил(ваш_логин@pmpractice.ru).
			`
		fmt.Println(example)
		c := exec.Command(ps, "npm adduser --registry="+config.Registry+" --always-auth")
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		c.Run()
	}

	return false
}

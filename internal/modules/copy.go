package modules

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func copyFile(from, in string) error {
	fileFrom, err := os.Open(from)
	defer fileFrom.Close()
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось открыть файл " + from}
	}

	fileIn, err := os.Create(in)
	defer fileIn.Close()
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось создать файл " + in}
	}
	_, err = io.Copy(fileIn, fileFrom)
	if err != nil {
		return &CopyError{err: err, msg: "\nНе удалось скопировать файл " + from}
	}
	return nil
}

// CopyModule ...
func (m *Module) CopyModule(moduleName string, moduleNames map[string]*Config) error {
	in, from, err := m.findModule(moduleName, moduleNames)
	if err != nil {
		return err
	}

	m.updateNodeModules()
	fmt.Println("node_modules обновлены")
	fmt.Println("\nОбновление модуля проекта...")

	err = m.deleteModule(moduleName, moduleNames)
	if err != nil {
		return err
	}

	err = os.MkdirAll(*in, 0777)
	if err != nil {
		return err
	}
	return filepath.Walk(*from, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, `\`, "/", -1)
		fromFile := path
		inFile := strings.Replace(path, *from, *in, -1)
		if !info.IsDir() {
			if !strings.Contains(path, ".config") {
				err = copyFile(fromFile, inFile)
				if err != nil {
					return err
				}
			} else {
				toCopy := true
				confFrom := strings.Replace(path, *from, "", -1)
				for i := 0; i < len(m.folderThatContainConfig[moduleName]); i++ {
					confIn := m.folderThatContainConfig[moduleName][i]
					var comfirmPath string
					if len(confIn) > len(confFrom) {
						comfirmPath = strings.Replace(confIn, confFrom, "", -1)
					} else {
						comfirmPath = strings.Replace(confFrom, confIn, "", -1)
					}
					if comfirmPath == `` || comfirmPath == `/` || comfirmPath == `\` {
						toCopy = false
					}
				}
				if toCopy {
					copyFile(fromFile, inFile)
				}

			}

		} else {
			path = strings.Replace(path, *from, *in, -1)
			os.Mkdir(path, 0777)
		}

		return nil
	})
}

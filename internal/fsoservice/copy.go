package fsoservice

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Update ...
func (m *Service) Update(serviceName, from, in string) error {
	// _, _, err := m.findModule(serviceName, serviceNames)
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Обновление сервиса проекта...")

	err := m.deleteService(serviceName, from, in)
	if err != nil {
		return err
	}

	if err = m.copyServiceInternal(serviceName, from, in); err != nil {
		return err
	}

	return nil
}

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

func (m *Service) copyServiceInternal(serviceName, from, in string) error {

	return filepath.Walk(from, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, `\`, "/", -1)
		fromFile := path
		inFile := strings.Replace(path, from, in, -1)
		if info != nil && !info.IsDir() {
			if !strings.Contains(path, ".config") {
				err = copyFile(fromFile, inFile)
				if err != nil {
					return err
				}
			} else {
				toCopy := true
				confFrom := strings.Replace(path, from, "", -1)
				for i := 0; i < len(m.folderThatContainConfig[serviceName]); i++ {
					confIn := m.folderThatContainConfig[serviceName][i]
					var comfirmPath string
					/* сделать ревью кода ниже */
					if len(confIn) > len(confFrom) {
						comfirmPath = strings.Replace(confIn, confFrom, "", -1)
					} else {
						comfirmPath = strings.Replace(confFrom, confIn, "", -1)
					}
					if comfirmPath == `` || comfirmPath == `/` || comfirmPath == `\` {
						toCopy = false
					}
					/****************************/
				}
				if toCopy {
					copyFile(fromFile, inFile)
				}

			}

		} else {
			path = strings.Replace(path, from, in, -1)
			os.MkdirAll(path, 0777)
		}

		return nil
	})
}

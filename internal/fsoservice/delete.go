package fsoservice

import (
	"os"
	"path/filepath"
	"strings"
)

// DeleteModule ...
func (m *Service) deleteService(serviceName, from, in string) error {
	if in != "" {
		savePath := make(map[string]bool)
		allDir := []string{}
		filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
			path = strings.Replace(path, `\`, "/", -1)
			allDir = append(allDir, path)
			if strings.Contains(path, ".config") {
				confIn := strings.Replace(path, in, "", -1)
				m.folderThatContainConfig[serviceName] = append(m.folderThatContainConfig[serviceName], confIn)
				pathArr := strings.Split(path, "/")
				for _, v := range pathArr {
					savePath[v] = true
				}
			}
			return nil
		})
		for _, path := range allDir {
			pathDirArr := strings.Split(path, "/")
			for _, dir := range pathDirArr {
				if dir != "" {
					if _, ok := savePath[dir]; ok {
						continue
					} else {
						os.RemoveAll(path)
					}
				}

			}
		}
	}

	return nil
}

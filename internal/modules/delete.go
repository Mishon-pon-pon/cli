package modules

import (
	"os"
	"path/filepath"
	"strings"
)

// DeleteModule ...
func (m *Module) deleteModule(moduleName string, moduleNames map[string]*Config) error {
	in := moduleNames[moduleName].PathIn
	if in != "" {
		savePath := make(map[string]bool)
		allDir := []string{}
		filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
			path = strings.Replace(path, `\`, "/", -1)
			allDir = append(allDir, path)
			if strings.Contains(path, ".config") {
				confIn := strings.Replace(path, in, "", -1)
				m.folderThatContainConfig[moduleName] = append(m.folderThatContainConfig[moduleName], confIn)
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

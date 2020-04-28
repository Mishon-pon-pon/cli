package modules

import (
	"errors"
	"fmt"
)

// FindModule ...
func (m *Module) findModule(moduleName string, moduleNames map[string]*Config) error {
	if _, ok := moduleNames[moduleName]; ok {
		if moduleNames[moduleName].PathFrom == "" {
			return fmt.Errorf("Не указан путь(pathFrom) в fso_configs.json до модуля %s в node_modules", moduleName)
			/*Код ниже нужен для рекурсивного поиска в node_modules, директории с модулем(по имени)*/
			/*found := false
			return filepath.Walk("node_modules/", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
					return err
				}
				path = strings.Replace(path, `\`, "/", -1)
				if info.IsDir() {
					if strings.Contains(path, moduleName) && !found {
						moduleNames[moduleName].PathFrom = path
						found = true
					}
				}
				return nil
			})*/
		}
		return nil
	}
	return errors.New("\nНет такого модуля, либо он не указан в fso_config.json")
}

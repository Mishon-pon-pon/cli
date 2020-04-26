package modules

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// UpdateNodeModules ...
func (m *Module) UpdateNodeModules() {
	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Обновление node_modules...")
	c := exec.Command(ps, "npm update")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	c.Run()
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

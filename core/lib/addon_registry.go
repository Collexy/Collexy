package lib

import (
	"log"
)

var Modules []Module

// Register a Module
func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func main() {
	log.Println("addon_registry.go main()")
}

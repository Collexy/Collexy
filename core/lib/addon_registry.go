package lib

import (
	"log"
)

type ModulesSlice []Module

func (slice ModulesSlice) Len() int {
	return len(slice)
}

func (slice ModulesSlice) Less(i, j int) bool {
	return slice[i].Order < slice[j].Order
}

func (slice ModulesSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

var Modules ModulesSlice

// Register a Module
func RegisterModule(m Module) {
	Modules = append(Modules, m)
}

func main() {
	log.Println("addon_registry.go main()")
}

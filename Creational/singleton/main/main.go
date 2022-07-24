package main

import (
	"fmt"
	"singleton/idgenerator"
)

func main() {
	var idg idgenerator.IdGenerator
	idg = idgenerator.GetEagerInstance()
	for i := 0; i < 10; i++ {
		fmt.Println("id:", idg.GetId())
	}
	idg = idgenerator.GetLazyInstance()
	for i := 0; i < 10; i++ {
		fmt.Println("id:", idg.GetId())
	}
}

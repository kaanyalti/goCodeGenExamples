package main

import (
	"fmt"
	"goGenerateExample/usefulExample/noGeneration/types"
)

func main() {
	var pl types.PersonList
	pl = append(pl, &types.Person{Name: "Jane", Age: 32})
	pl = append(pl, &types.Person{Name: "Ed", Age: 27})

	pl2 := pl.Filter(func(p *types.Person) bool {
		return p.Age > 30
	})
	for _, p := range pl2 {
		fmt.Println(p)
	}
}
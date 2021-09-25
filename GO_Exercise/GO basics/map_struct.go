package main

import "fmt"

type Seeding struct {
	Revenue, GrossProfit float64
}

var m map[string]Seeding

func main() {
	n := map[string]float64{"kloud1": 76.6, "Accuknox": 67}
	m = make(map[string]Seeding)
	m["kloud1"] = Seeding{
		40.5, 74.6,
	} // map[string]Seeding{"kloud1": {40.5, 74.6}}
	fmt.Println(n, m["kloud1"]) // map[Accuknox:67 kloud1:76.6] {40.5 74.6}
}

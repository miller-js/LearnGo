package main

import (
	"fmt"
	"maps"
)

func runMaps() {

	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 13

	fmt.Println("map:", m)

	value1 := m["key1"]
	fmt.Println("value1:", value1)

	value3 := m["key3"]
	fmt.Println("value3:", value3)

	fmt.Println("len:", len(m))

	delete(m, "key2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}

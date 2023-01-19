package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	rucksacks := strings.Fields(string(b))

	var sum int
	for _, rucksack := range rucksacks {
		var common string
		for _, c := range rucksack {
			if strings.Contains(rucksack[:len(rucksack)/2], string(c)) && strings.Contains(rucksack[len(rucksack)/2:], string(c)) {
				common = string(c)
				break
			}
		}
		priority := 0
		if common >= "a" && common <= "z" {
			priority = int(common[0] - 'a' + 1)
		} else {
			priority = int(common[0] - 'A' + 27)
		}
		sum += priority
	}

	fmt.Println(sum)
}

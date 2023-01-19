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
	data := strings.Fields(string(b))

	var sum int
	for i := 0; i < len(data); i += 3 {
		var common string
		for _, c := range data[i] {
			if strings.Contains(data[i], string(c)) && strings.Contains(data[i+1], string(c)) && strings.Contains(data[i+2], string(c)) {
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

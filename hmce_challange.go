package main

import (
	"fmt"
)

func convert(kalimat string) string{
	var result string
	for _, char := range kalimat {
		if char == 32 || char == 33 || char == ','{
			char = char
		} else {
			char = char + 13
			if char > 122 {
				char = char - 26
			} else {
				char = char
			}
		}
		result = result + string(char)
	}
	return result
}

func main() {
	kalimat1 := "frznatng xhyvnualn, wnatna gryng znfhx xrynf!"
	kalimat2 := "lhx vxhg ybzon yntv erx!"

	fmt.Println("Kalimat 1 : " + convert(kalimat1))
	fmt.Println("Kalimat 2 : " + convert(kalimat2))
}

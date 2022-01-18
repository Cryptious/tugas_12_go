package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	// Test input = "JeRuk1 pis8ang9 anGgur8"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		var regex, err = regexp.Compile(`[A-Z]+\D`)
		if err != nil {
			fmt.Println(err.Error())
		}
		var hasil = regex.FindAllString(input, -1)
		fmt.Println(hasil)

		var index1 = regex.FindStringIndex(input)
		fmt.Println(index1)

		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}

}

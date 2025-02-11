package main

// First steps with Go :)

import (
	"basics/fileops"
	"errors"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func readIn(name string) (int, error) {
	var x = 0
	fmt.Printf("Give me %s= ", name)
	fmt.Scan(&x)
	if x <= 0 {
		return 0, errors.New("Values >=0!")
	}
	return x, nil
}

const fName = "results.txt"

func main() {
	var x = 0
	var y = 0
	presentHeader()
	fmt.Printf("Reach us 24h at: %s\n", randomdata.PhoneNumber())
	balance, err := fileops.GetStringFromFile(fName)
	if err == nil {
		fmt.Printf("Last result: %s\n", balance)
	} else {
		fmt.Println(err)
	}

	for {

		x, err = readIn("y")
		if err != nil {
			fmt.Print(err)
			return
		}

		y, err = readIn("y")
		if err != nil {
			fmt.Print(err)
			return
		}

		result1, result2 := calculate2(x, y)
		fmt.Println(result1)
		fmt.Println(result2)
		fileops.SaveTwoIntsToFile(fName, result1, result2)

		if result2 > 0 {
			fmt.Println((">0"))
		} else if result2 < 0 {
			fmt.Println(("<0"))
		} else {
			fmt.Println(("==0"))
		}

	}

}
func calculate2(a int, b int) (sum int, div int) {
	sum = a + b
	div = a - b
	return sum, div
}

// pointers and conditions
package main

import (
	"fmt"
	"errors"
)

const winePrice = 100

func main() {
	name := "Vasya"
	fmt.Println(name)

	setName(&name)
	fmt.Println(name)


	change, err := buyWine(11, 8)
	if err != nil {
		fmt.Println("Pokupka failed", err.Error())
	} else {
		fmt.Printf("Your sdacha - %d. Have a good day!", change)
	}
}

func setName(name *string) {
	*name = "Jikia"
}

func buyWine(age, moneyMount int) (int,error) {
	if age >= 10 && moneyMount >= winePrice {
		return moneyMount - winePrice, nil		
	} else if age < 18 {
		return moneyMount, errors.New("You have to go drink tea")
	} else {
		return moneyMount, errors.New("You do not have enough moeny")
	}
}
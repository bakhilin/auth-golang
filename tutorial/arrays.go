package main

import (
	"fmt"
)

func main(){
	contactsList := [3]string{"Kostya", "Oleg", "Serega"}
	contactsList[1] = "Leonid"

	fmt.Println("List of your contact:")
	for index, value := range contactsList {
		fmt.Printf("%d. %s\n", index, value)
	}

	for i:=0; i <= 10; i++ {
		fmt.Println(i)
	}

	//slices
	slice := []int{1,2,3,}
	slice = append(slice, 22,3,4)
	fmt.Println(slice)
}
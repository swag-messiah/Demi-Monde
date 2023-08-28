package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Ширина %0.2f не подходит", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("Высота %0.2f не подходит", height)
	}

	area := width * height
	return area / 10.0, nil
}

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f литров нужно\n", amount)
	}
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f литров нужно\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f литров нужно", total)
}

//Check Test Demi-Monde

/*func main() {

	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}*/

package main

import (
	"fmt"
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
	amount, err := paintNeeded(4.2, 3.0)
	fmt.Println(err)

	fmt.Printf("%0.2f литров нужно\n", amount)
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f литров нужно\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f литров нужно", total)
}

//Check Test Demi-Monde

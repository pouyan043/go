package main

import (
	"fmt"
	"time"
)

func main() {

	var balance float64 = 1000

	for {

		t := time.Now()

		switch {
		case t.Hour() < 12:
			fmt.Println("before noon")
		case t.Hour() < 18:
			fmt.Println("after noon")

		default:
			fmt.Println("evening")

		}

		fmt.Printf("welcome\nyour amount : %v\n", balance)
		fmt.Println("press 1 for withdraw")
		fmt.Println("press 2 for deposit")
		fmt.Println("press 3 for balance")
		fmt.Println("press 4 for exit")
		var UserChoice int
		fmt.Println("choose your option")
		fmt.Scanln(&UserChoice)
		if UserChoice == 1 {
			var withdrawAmount float64
			fmt.Print("withdraw Amount :")
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount <= balance {
				balance -= withdrawAmount
				fmt.Println(balance)
			} else {
				println("invalid amount !!")
			}

		} else if UserChoice == 2 {
			var depositAmount float64
			fmt.Print("depositAmount : ")
			fmt.Scanln(&depositAmount)

			balance += depositAmount
			fmt.Printf("balance : %v\n ", balance)

		} else if UserChoice == 3 {
			fmt.Printf("balance : %v\n ", balance)

		} else {
			break

		}
	}

}

// DRY = DON'T REPEAT YOURSELF
//  EXAMPLE :
//  func a (b, c int ) {
// fmt.Println(b+c)
// }
// func g () {
// a(2,4)
// }

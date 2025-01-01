package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Print(" input farenheit :")
	fmt.Scan(&farenheit)
	celsius := (farenheit - 32)*5 / 9
	fmt.Println(celsius)
}

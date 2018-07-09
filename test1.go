package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	newSlice := []int{7, 8}

	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}

package main

import "fmt"

func SliceTotal(Slice[]int)(int){
	SliceSum := 0
	for _, value := range Slice{
		SliceSum += value
	}
	return SliceSum
}
func main() {
	fmt.Println("Slice Total:", SliceTotal([]int{1,2,3,4,5}))
}

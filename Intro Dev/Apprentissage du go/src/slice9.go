package main

import "fmt"

func main() {
	var tab []int = []int{7,2,3,4,5}
	var t []int = make([]int, 2 ,4 )
	copy(t, tab)
	fmt.Println("tab:", tab, len(tab), cap(tab))
	fmt.Println("t:", t, len(t), cap(t))

	t[0] = 0
	fmt.Println("tab:", tab, len(tab), cap(tab))
	fmt.Println("t:", t, len(t), cap(t))
}

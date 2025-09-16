package main

import "fmt"


func main() {
	var Input []int = []int{1, 2, 3, 4, 5, 6 ,7 ,8 ,9 , 10}
	c := 0
	d := 0
	for i := 0; i < len(Input); i+=2 {
   		c = c + Input[i] 
		d = d + Input[i+1]
	}
	fmt.Println(c+d)
}

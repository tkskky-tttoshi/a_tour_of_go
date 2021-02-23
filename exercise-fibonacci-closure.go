package main

import "fmt"

func fibonacci() func() int{
	x, y := 0, 1

	return func() int{
		//now -> x -> y
		now := x
		x = y
		y = x + now
		return now
	}
}

func main(){
	f := fibonacci()
	for i := 0; i < 10; i++{
		fmt.Println(f())
	}
}

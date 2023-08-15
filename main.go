package main

import "fmt"
func main() {
		fmt.Println("Hello, World!")

		//decalre variable for specific type int8,int16,int32,int64
		//positive ints are uint8,uint16,uint32,uint64
		//float32,float64,string,bool
		// var x int8 = 5
		// var y string = "hey"
    // fmt.Println(x, y)

		//short hand declaration
		x,y:=5,"hey"
		fmt.Println(x, y)

		//printf formated way
		fmt.Printf("x is of type %T\n", x)
		fmt.Printf("y is %v and x is %v",x, y)

} 
package main

// import "fmt"
import ("fmt" 
       "strings"
			)

func main() {
	fmt.Println("Hello, World!")

	//decalre variable for specific type int8,int16,int32,int64
	//positive ints are uint8,uint16,uint32,uint64
	//float32,float64,string,bool
	// var x int8 = 5
	// var y string = "hey"
	// fmt.Println(x, y)

	//short hand declaration
	x, y := 5, "hey"
	fmt.Println(x, y)

	//printf formated way
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is %v and x is %v\n", x, y)

	// //input from user can be done by pointer
	// var input string
	// fmt.Println("Enter your name: ")
	// fmt.Scanln(&input)

	// 	//uint is for positive numbers
	//  totalticket:=5
	//  userticketbuy:=3 //user can't buy -ve tickets
	//  remainingticket := totalticket-userticketbuy
	//  fmt.Println(remainingticket)
	//  fmt.Printf("remaining ticket is %v\n",remainingticket)

	//arrays
	ar := [5]int{1, 2, 3, 4, 5} //st
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(ar)
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	//normally when value not known
	var arr3 [5]int
	arr3[0] = 1
	fmt.Println(arr3) //[1 0 0 0 0]

	//slices when size not known
	sli := []int{} //st
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(sli)
	//slice can be created from array
	var slice2 = arr[1:3]
	fmt.Println(slice2)
	//append to slice
	var slice3 []int
	slice3 = append(slice3, 1)
	// slice3[5]=2//it will give error
	fmt.Println(slice3) //[2]





	//for loop, no while loop in go
	// for{} //infinte loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	months := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	for i := range months {
		fmt.Println(months[i])
	}

	//while loop as For loop
  i:=0
	for i<20{
		fmt.Println(i)
		if i==10{
			break
		}
		i++
	}


	//goto statement
	// goto label
	for i:=0;i<10;i++{
		fmt.Println(i)
		if i==5{
			goto label
		}
	}
	label:fmt.Println("goto statement")


	// fields to separate the string into substrings
	str := "Welcome to the online portal of GeeksforGeeks"
	// Split the string around spaces.
	result := strings.Fields(str)
	// Display each word.
	for i := range result {
		fmt.Println(result[i])
	}



	

	

}

////slide 4 : problem solving - problem 5
//package main
//
//import "fmt"
//
//func main() {
//	var x, y int
//	fmt.Println("please type a Number :")
//	fmt.Scanln(&x)
//	fmt.Println("please type another one :")
//	fmt.Scanln(&y)
//	findMax(x, y)
//}
//func findMax(x, y int) {
//	if x > y {
//		fmt.Println("the biggest number is ", x)
//
//	} else {
//		fmt.Println("the biggest number is ", y)
//	}
//}
//-------------------------------------------------------------------------------------------------

////slide 4 : problem solving - problem 8
//package main
//
//import "fmt"
//
//func main() {
//	var input int64
//	fmt.Println("please input your number")
//	fmt.Scanln(&input)
//
//	for {
//		if input%5 == 0 && input%11 == 0 {
//			fmt.Println("your number is divisible by 5 and 11 ")
//			break
//		} else {
//			fmt.Println("your number isn't divisible , take another one please")
//			fmt.Println("enter a new one : ")
//			fmt.Scanln(&input)
//		}
//	}
//}
////-------------------------------------------------------------------------------------------------

////slide 4 : problem solving - problem 9
//package main
//
//import "fmt"
//
//func main() {
//	var input int
//	fmt.Println("plz enter a number")
//	fmt.Scanln(&input)
//
//	//#1      is not true  - it has bug
//	for input >= 1 {
//		fmt.Println(input - 1)
//		input--
//	}
//
//	//#2        is true
//	for i := input - 1; i > 0; i-- {
//		fmt.Println(i)
//	}
//
//}
//------------------------------------------------------------------------------------------

////slide 4 : problem solving - problem 10
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var input int32
//	fmt.Println("enter a number : ")
//	fmt.Scanln(&input)
//
//	fmt.Println("-------------------------------------------------------------------------")
//	recognizeEvenNumber(input)
//}
//
//func recognizeEvenNumber(x int32) {
//
//	for i := x - 1; i >= 0; i-- {
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	}
//}
////---------------------------------------------------------------------------------------------------

////slide 4 : problem solving - problem 11
//package main
//
//import "fmt"
//
//func primeFunc(primNum int) {
//	var primCount int = 0
//	for i := 2; i < primNum; i++ {
//		if primNum%i == 0 {
//			primCount++
//		}
//	}
//	if primCount == 0 && primNum != 1 {
//		fmt.Println(primNum)
//	}
//}
//
//func main() {
//
//	var primNum int
//
//	fmt.Print("Enter the Number to find the Prime Numbers = ")
//	fmt.Scanln(&primNum)
//
//	for X := primNum; X > 0; X-- {
//		primeFunc(X)
//
//	}
//
//}
////-------------------------------------------------------------------------------------------------

//slide 4 : problem solving - problem 12
//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	fmt.Printf("Enter size of your array: ")
//	var size int
//	fmt.Scanln(&size)
//
//	var arr = make([]int, size)
//	for i := 0; i < size; i++ {
//		fmt.Println("Enter your element : ")
//		fmt.Scanln(&arr[i])
//	}
//
//	sort.Ints(arr[:])
//	fmt.Println("Your sorted array is: ", arr)
//	fmt.Println("The smallest number is :  ", arr[0])
//	fmt.Println("The median number is :  ", arr[size/2])
//	fmt.Println("The largest number is :  ", arr[size-1])
//
//	arrSum := 0
//	for i := 0; i < size; i++ {
//		arrSum += (arr[i])
//	}
//	average := arrSum / size
//	fmt.Printf("The Sum of Array Items = %d\n", arrSum)
//	fmt.Printf("The Average of Array Items = %d", average)
//
//}
//--------------------------------------------------------------------------------------------------------

////slide 4 : problem solving - problem 14 - 15
//package main
//
//import "fmt"
//
//func main() {
//	var n int
//	fmt.Print("Enter the number: ")
//	fmt.Scanln(&n)
//	num := numberReverse(n)
//	count := 0
//	for n >= 1 {
//		n = n / 10
//		count++
//	}
//	fmt.Printf("The number of digits in the given number is: %d\n", count)
//	fmt.Printf("and The reversed number is: %d\n", num)
//
//}
//
//func numberReverse(n int) int {
//	new_int := 0
//	for n > 0 {
//		remainder := n % 10
//		new_int *= 10
//		new_int += remainder
//		n /= 10
//	}
//	return new_int
//}
////-----------------------------------------------------------------------------------------------------


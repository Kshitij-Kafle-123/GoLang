package main

import (
	"fmt"
)

var abcd string

func main() {
	fmt.Println("Hello GoLang")
	//fmt.Print("hello")

	var anyThing string = "Hello"
	var student = "student"
	abc := 235
	fmt.Println(abc)
	fmt.Println(anyThing)
	fmt.Println(student)
	abcd = "abc"
	fmt.Println(abcd)
	var a string
	var b int
	var c bool
	var d float32
	var e float64

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	var x, y, z int = 14, 15, 16
	var w, m, n = 500, "hello", 3.14
	fmt.Println(w, m, n)
	fmt.Println(x, y, z)
	println(x)
	const PI float32 = 3.14
	fmt.Println(PI)

	const (
		Q   = 4.5
		AQ  = 8.9
		WWW = "World Wide Web"
	)

	fmt.Println(WWW)

	var name string = "Kshitij Kafle."
	var age int = 23

	fmt.Printf("My Name is: %v and My Age is: %v", name, age)

	var test = 100.25
	var test1 = "anything testing"

	fmt.Printf("Value: %v \n", test)
	fmt.Printf("Go %#v \n", test1)
	fmt.Printf("Type: %T\n", test1)
	fmt.Printf("Percent Sign: %v%%\n", test)

	var testArray = [3]int{1, 2, 3}
	fmt.Println(testArray)
	var inferedArray = [...]int{4, 5, 6, 7, 8}
	fmt.Println(inferedArray)
	arrayIsArray := [...]int{5, 7, 9}
	fmt.Println(arrayIsArray)
	var arrayString = [...]string{"abc", "def", "ghi"}
	fmt.Println(arrayString)

	arr1 := [...]int{1, 2}
	arr2 := [3]int{}
	arr3 := [3]int{2, 3}

	fmt.Println(arr1, arr2, arr3)
	if arr2[2] == 0 {
		fmt.Println("In IF CONDITION")
	}
	//This initialize position 2nd and 5th (1 and 4)
	amazingArray := [5]int{1: 15, 4: 58}
	fmt.Println(amazingArray)
	fmt.Println("Length of the array: ", len(amazingArray))

	newSlice := []int{1, 2, 3}
	fmt.Println("Slice in GOLANG: ", newSlice)

	//Make a slice from an Array

	var arrBeforeSlice = [7]int{4, 2, 6, 5, 5, 6, 8}
	sliceFromArr := arrBeforeSlice[2:5]
	fmt.Println("Slice created from Array: ", sliceFromArr, "Length: ", len(sliceFromArr), "Capacity: ", cap(sliceFromArr))
	anotherSlice := arrBeforeSlice[:5]
	fmt.Println("New Slice: ", anotherSlice)

	//Creating a Slice from a make()  function

	newSliceforMake := make([]int, 6, 10)
	fmt.Println("Slice from make function ", newSliceforMake)

	//appending element in a slice

	appendSlice := []int{7, 5, 5, 8}
	fmt.Println("Append Slice before appending: ", appendSlice)
	appendSlice = append(appendSlice, 89, 56)
	fmt.Println("Append Slice after appending: ", appendSlice)
	sliceAnother := append(appendSlice, appendSlice...)
	fmt.Println("Two slices: ", sliceAnother)

	//Memory efficiency
	heavySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 5, 62, 216, 35, 8, 35, 225}
	fmt.Println("Memory Consuming Slice: ", heavySlice)
	neededElemOfHeavySlice := heavySlice[1 : len(heavySlice)-5]
	copyOfSlice := make([]int, len(neededElemOfHeavySlice))
	copy(copyOfSlice, neededElemOfHeavySlice)
	fmt.Println("Copy Of Heavy Slice: ", copyOfSlice)

	var operatorr = 10
	operatorr += 520
	fmt.Println("After Addition of 520: ", operatorr)

	var less = 5
	var more = 10
	fmt.Println(less < more)

	//LOGIC AND OR
	var binX = 9
	var binY = 10
	fmt.Printf("x in binary or base2 : %b\n", binX)
	fmt.Printf("y in binary or base2 : %b\n", binY)

	//Use AND-Gate (Returns 1 if both bits are 1 otherwise returns 0)
	fmt.Printf("binX AND binY (binX & binY): %b\n", binX&binY)
	//Returns 1 if one of the bit is 1
	fmt.Printf("binX OR binY (binX | binY): %b\n", binX|binY)

	fmt.Printf("XOR: (binX ^ binY): %b\n", binX^binY)

	//shift
	fmt.Printf("Shift by 2: binX<<2: %b\n", binX<<2)
	fmt.Printf("Shift by 2: binX>>2: %b\n", binX>>2)

	//switch
	dayOftheWeek := 3
	switch dayOftheWeek {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	default:
		fmt.Println("Not a day")
	}
	//Switch MultiCase
	validNums := 9
	switch validNums {
	case 1, 2, 3, 4, 5:
		fmt.Println("Valid Number")
	case 8, 9:
		fmt.Println("Invalid Number")
	}

	//for loops
	for i := 0; i < 5; i++ {
		fmt.Printf("For loop index in Binary: %b\n", i)
	}
	//range keyWord in looping
	data := []string{"hello", "learn", "golang", "better", "way"}
	for index, value := range data {
		fmt.Printf("index: %v\t Value: %v\n", index, value)
	}
	for _, val := range data {
		fmt.Printf("index: is N/A \t value is: %v\n", val)
	}
	newFunction(256)
	errOrFunction(fmt.Errorf("error: 404"))
	var result = sumOfTwoInt(10, 14)
	fmt.Println("output from sum  function: ", result)
	fmt.Println("Calling achhamma ko function: ", achhamma(45, 4.5))

	fmt.Println(iLoveThisFunction(10, "Hello"))
	nums, txt := iLoveThisFunction(5, "hello")
	fmt.Println("values: ", nums, txt)

	//fmt.Println("Calling a recursive function: ", RecursionFunc(1))

	var man1 Man
	var man2 Man
	man1.asian = true
	man1.name = "kshitj"
	man1.age = 23

	man2.age = 44
	man2.name = "koii"
	man2.asian = false

	//fmt.Println("name: ", man1.name, "age: ", man1.age, "isAsian: ", man1.asian)

	//man1
	printMan(man1)
	//man2
	printMan(man2)

	//Maps Of GOLANG

	var mapA = map[string]int{"a": 1, "b": 2}
	fmt.Println("Map values: ", mapA["b"], mapA["a"])

	mapB := make(map[int]int)
	mapB[1] = 56
	mapB[2] = 567

	fmt.Println("MapB values are: ", mapB[1], mapB[2])

	var maapp = make(map[int]int)
	var maaappp map[string]string
	fmt.Printf("The type and values are: %T\t%v\n", maapp, maapp)
	fmt.Printf("The type and values are: %T\t%v\n", maaappp, maaappp)
	fmt.Println(maapp == nil)
	fmt.Println(maaappp == nil)
	//Checking value in a map
	val, ok := mapB[1]
	fmt.Println(val, ok)

	for k, v := range mapB {
		fmt.Println(k, v)
	}

}
func newFunction(abc int) {
	fmt.Println("i'm in new Function: Thank You for Calling Me !! the value you passed to me as a gift is : ", abc)
}
func errOrFunction(err error) {
	fmt.Println("What the Error: ", err)
}
func sumOfTwoInt(x int, y int) int {
	return x + y
}
func achhamma(x int, y float32) (kHoYesto float32) {
	kHoYesto = float32(x) + y
	return
}
func iLoveThisFunction(x int, y string) (num int, txt string) {
	num = x + x
	txt = y + " World!!"
	return
}

//Recursion
func RecursionFunc(x int) int {
	if x > 15 {
		return x
	}
	fmt.Println(x)
	return RecursionFunc(x + 1)
}

type Man struct {
	name  string
	age   int
	asian bool
}

func printMan(man Man) {
	fmt.Println("Name: ", man.name, "Age: ", man.age, "IsAsianMan: ", man.asian)
}

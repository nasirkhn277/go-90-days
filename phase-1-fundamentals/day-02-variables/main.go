package main

import (
	"fmt"
)

/*func main() {

	//1. var with explicit type
	var age int = 25

	//2. var with infered type (Go figure out it type by value)
	var name = "Nasir"

	//3. self declaration (most common inside the function and it only work inside the function)
	city := "Delhi"

	//4. multiple varibale at once
	var x, y int = 10, 20
	a, b := "hello", true

	fmt.Println(age, name, city, x, y, a, b)

} */

/*

var isActive bool     = true
var count    int      = 42        // platform-dependent size (usually 64-bit)
var price    float64  = 19.99
var initial  byte     = 'A'       // alias for uint8
var letter   rune     = '❤'       // alias for int32, represents a Unicode code point
var text     string   = "Go"

int, int8/16/32/64		whole numbers
uint, uint8/16/32/64	unsigned (non-negative) whole numbers
float32, float64		decimals (default to float64)
bool 					true/false
string 					text(immutable, UTF-8 by default)
byte					alias for uint8 — raw bytes
rune					alias for int32 — a single Unicode character

Note: There is no null concept in Go

*/

func main() {

	var name string = "Nasir"
	age := 30
	var height float32 = 13.09
	var weight float64 = 90.12
	isLearning := true

	fmt.Printf("%v is a type %T\n", name, name)
	fmt.Printf("%v is a type %T\n", age, age)
	fmt.Printf("%v is type of %T\n", name, name)
	fmt.Printf("%v is type of %T\n", weight, weight)
	fmt.Printf("%v is type of %T\n", height, height)
	fmt.Printf("%v is type of %T\n", isLearning, isLearning)

}

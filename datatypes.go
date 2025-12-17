package main

import (
	"fmt"
	"unicode/utf8"
)

func datatypes() {
	fmt.Println("Integers")
	var Intnum int = 236598
	fmt.Println(Intnum + 1)
	// float is either float32 or float64
	var floatNum1 float64 = 73.5
	fmt.Println(floatNum1)
	var floatNum2 float32 = 88.7
	fmt.Println(floatNum2)
	var result float64 = floatNum1 + float64(floatNum2)
	fmt.Println(result)

	var myStr string = "Hello" + " " + " Bro"
	fmt.Println(myStr)
	var str2 string = `First line 
	second line `
	fmt.Println(str2)

	fmt.Println(utf8.RuneCountInString("heyyy")) // gives no of chrs

	var myrune rune = 'A' //ascii value
	fmt.Println(myrune)

	var mybool bool = false
	fmt.Println(mybool)

	var defaultNum rune
	fmt.Println(defaultNum) // 0 by default

	var name string
	fmt.Println(name) // " " empty str by default

	name1 := "ghyan"
	fmt.Println(name1)
	a, b := 1, 2
	fmt.Println(a, b)

	const pi float64 = 3.14
	fmt.Println(pi)
}

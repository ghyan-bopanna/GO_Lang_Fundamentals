## for new project initialise a module

`go init mod go_tutorials`

## name of the package should be the same for all files within the folder

- use the fmt package to print something (print,Println)

## two ways to run

- create a build file then run `build main.go`, creates a binary file you can run
- or use the run command to directly create a binary file and run it `go run main.go`

## Datatypes

- int int16 int32 int64

- just 'int' > becomes 32/64 based on your sys architecture
- uint uint16 unsigned int is used to store same integers but only positive values ,
- int 16 max value it can store = 32767

- int8 : (-128,127)
- uint8 : (0,255) > hense lets us use integers twice as large in same amount of memory.

- float 32/64

- ALWAYS MAKE SURE TO USE THE DATATYPES ACCORDING TO YOUR NEEDS 8/16/32/unsigned

## Ooperations

- Cant do arithematic operations on different types
- Integer division rounds DOWN the result. 3/2 =1 , not 1.5

## Strings

- Use " ",` `

```
     var myStr string = "Hello" + " " + " Bro"
      fmt.Println(myStr)
      var str2 string = `First line
      second line `
      fmt.Println(str2)

  fmt.Println(len("hello")) // gives no of bytes , not no of chrs
```

- import "unicode/utf8" // to count len of chrs

- fmt.Println(utf8.RuneCountInString("heyyy")) // gives no of chrs

## Runes // for chrs

-     var myrune rune = 'A' //ascii value
  fmt.Println(myrune)

## declaring usig type infering

var1, var2 := 2,3

# control flow

## functions

- functions can return multiple variables using parethesis (int, int)

```
func main() { // start the bracket in the same line in function
	var printvalue string = "Yoooo"
	var1, var2 := 1, 2
	printme(printvalue)
	var sum, dif int = Oprof(var1, var2)                                                    //storing the return values into 2 diff variables
	fmt.Println("The Sum of ", var1, "and ", var2, " is", sum)                              //using println
	fmt.Printf("The Sum of the variables is %v and the difference of them is %v", sum, dif) //using printf easier formatting

}

func printme(printvalue string) { //void function
	fmt.Println(printvalue)
}

func Oprof(var1 int, var2 int) (int, int) { // int fun , returns 2 int values
	var sum int = var1 + var2
	var dif int = var1 - var2
	return sum, dif
}

```

## if statement

```
if a > b && a > c {
		max = a
	} else if b > a && b > c { // else starts with the end bracket of if
		max = b
	} else {
		max = c
	}
	return max


```

&& -> and opr
|| -> or opr

## handeling custom errors using type error

```
import (
	"errors"
	"fmt"
)

func main() {
	var a, b, c int = -10, 6, 2
	var maxi, err = findMax(a, b, c)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("The maximum value is %v ", maxi)
}

func findMax(a int, b int, c int) (int, error) {
	var err error // default inital value is `nil`
	if a < 0 || b < 0 || c < 0 { // custom error type to mke sure user dosnt use -ve numbers
		err = errors.New("Input cannot be negative")
		return 0, err
	}
	var max int
	if a > b && a > c {
		max = a
	} else if b > a && b > c {
		max = b
	} else {
		max = c
	}
	return max, err
}
```

## Conditional Switch

switch remainder {
case
}

## conditional switch

```
// use switch for a particular variable in the func , ie:remainder
switch reminder {
	case 0:
		fmt.Printf("the division was exact")
	case 1,2 :
		fmt.Printf("The division was close")
	default :
		fmt.Printf("The division was not close")
	}

```

## [] Arrays

- Fixed Length
- Same Type
- Indexable
- Contiguous in Memeory

`
func main() {
var intArr [3]int32 = [3]int32{1, 2, 3} //can be assigner like this
intArr2 := [...]int{1, 2, 3, 4, 5} // array size can be inferred
intArr[1] = 88
fmt.Println(intArr)
fmt.Println(intArr2)
fmt.Println(&intArr[0]) // print array address
fmt.Println(&intArr[1])
fmt.Println(&intArr[2])

}

`

## Slices

- Slices are wrappers of arrays with additional functionality
- very similar to arrays
- functions in slices
  - append
  - len
  - cap (capacity)
- when you append a slice , you create a whole new array with larger capacity and copy elements from old array.

`
func main() {
var intSlice []int32 = []int32{1, 2, 3}
intSlice2 := []int32{1, 2, 3, 4, 5}
intSlice[1] = 88
fmt.Printf("Slice 1 has the length %v and capacity %v", len(intSlice), cap(intSlice))
intSlice = append(intSlice, intSlice2...)
fmt.Println(" ")
fmt.Printf("Slice 1 has the length %v and capacity %v", len(intSlice), cap(intSlice))
fmt.Println()
fmt.Println(intSlice)
fmt.Println("size 5 capacity 5", intSlice2)
fmt.Println(cap(intSlice2))
intSlice2 = append(intSlice2, 2)
fmt.Println(cap(intSlice2))
fmt.Println("size 6 capacity 12", intSlice2)
fmt.Println(&intSlice[0])
fmt.Println(&intSlice[1])
fmt.Println(&intSlice[2])

}

`

### output

```
Slice 1 has  the length 3 and capacity 3
Slice 1 has  the length 8 and capacity 8
[1 88 3 1 2 3 4 5]
size 5 capacity 5 [1 2 3 4 5]
5
12
size 6 capacity 12 [1 2 3 4 5 2]
0xc0000141c0
0xc0000141c4
0xc0000141c8
```

## Maps

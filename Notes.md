## for new project initialise a module

`go init mod go_tutorials`

## name of the package shoulld be the same for all files within the folder

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
func main() {
	var printvalue string = "Yoooo"
	var1, var2 := 1, 2
	printme(printvalue)
	var sum, dif int = Oprof(var1, var2) //storing the return values into 2 diff variables
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
if 1==1 && 2==2{
	...

}else{
	...
}

```

&& and opr
|| or opr

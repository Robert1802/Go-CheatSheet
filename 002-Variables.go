// Go Values
/*
	In Go, values can be unnamed or named. Unnamed values are literals such as 3.14, true, and "Codecademy".
	Named values have a name attached to the value and they can either be unchangeable as constants or changeable as variables once defined.
*/

// literal unnamed value
fmt.Println("PI = ", 3.14159)

// constant named value
const pi = 3.14159

// variable named value
var radius = 6

// ------------------------------------------------

// Go Data Types
/* 
	In Go, values have a data type.
	The data type determines what type of information is being stored and how much space is needed to store it.
	Go has basic data types such as:
*/

- string
- bool
- numeric types:
	- int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
	- float32, float64
	- complex64, complex128

// ------------------------------------------------

// Go Variables
/*
	A Go variable has a name attached to a value but unlike a Go constant, a variable’s value can be changed after it has been defined.
	There are four ways to declare and assign a Go variable:
*/

// use the var keyword followed by a name and its data type. This variable can be assigned later in the program. For example:
var fruit string
fruit = "apple"

// use the var keyword followed by a name, data type, = and value.
var fruit string = "apple"

// use the var keyword, followed by a name, = and value. Ignore the data type and let the compiler infer its type.
var fruit = "apple"

// skip the var keyword, define a name followed by := and value and let the compiler infer its type.
fruit := "apple"

// ------------------------------------------------

// Go Errors
/*
	In Go, errors are raised when the compiler doesn’t recognize the code as valid.
	The error message is printed to the terminal and contains the following information:

	- The filename
	- The line that raises the error
	- The number of characters from the left side that raises the error
	- The type of error and reason for raising the error

	For example:

	./Main.go:11:3: undefined: dinner
	This particular error occurs in the file main.go at line 11, 3 characters into the line, and its error type and reason is "undefined: dinner".

*/
// ------------------------------------------------

// Go Strings
/*
	A Go string is a data type that stores text or a sequence of characters in any length in double-quoted form.
	To concatenate two strings, use the + operator.
*/

var firstName string = "Abe"
var lastName string = "Lincoln"

// prints "Abe Lincoln"
fmt.Println(firstName + " " + lastName)

// ------------------------------------------------

// Go Zero Values
/*
	In Go, when a variable is declared without initializing a value, it has a default value.
	The default value is known as the zero value.

	Different zero values exist for different data types:

	- Type    Zero Value
	- ints     0
	- floats   0
	- string   "" (empty string)
	- boolean  false

*/
// ------------------------------------------------

// Go Inferred Int Type
/*
	When we declare a Go variable without specifying its data type and assign the variable (using := or var =) to a whole number,
	the Go compiler automatically infers the variable data type as an int. For example:
*/

score := 85
var temperature = 60

// ------------------------------------------------

// Go Updating Variables
// Unlike constants, Go variables can change their values if we reassign new values to them. For example:

var zipcode = "02134"
zipcode = "03035"

// Go supports additional assignment operators that updates a variable by performing an operation such as addition, subtraction, multiplication or division to iself.

// sum = sum + value
sum += value

// total = total - value  
total -= value   

// average = average / quantity  
average /= quantity

// price = price * quantity
price *= quantity  

// ------------------------------------------------

// Go Multiple Variable Declaration
/*
	Multiple Go variables can be declared and initialized on the same line delimited with a comma.
	If they are of the same type, the type can be optionally declared after the variable names before the assignment operator. For example:
*/

var x, y int = -1, 5
a, b := 7, 2
fmt.Println(x, y, a, b) // -1, 5, 7, 2

// If the variables are of different types, they can also be declared on the same line without the type designation.

found, answer := true, "yes"
var name, age = "Steve", 35
fmt.Println(found, answer, name, age) // true, "yes", "Steve", 35
package main

import "fmt"

func main(){
	// printing formating varb
	// %v = print default formating
	// %#v = print the go syntax
	// %% = print % 
	// %T = print value type
	fmt.Println("------------------------------------------------------")

	numberTest := 15.5
	stringTest := "Shuvo"
	fmt.Printf("%v\n",numberTest) // showing only value
	fmt.Printf("%T\n", numberTest) // showing type

	fmt.Printf("%#v\n", stringTest) // showing with ""
	fmt.Printf("%v\n", stringTest) // showing only text

	fmt.Println("------------------------------------------------------")
	// integer verb
	testingNumber := 15
	fmt.Printf("%b\n", testingNumber) // with base2
	fmt.Printf("%d\n", testingNumber) // with base 10
	fmt.Printf("%+d\n", testingNumber) // base 10 and show sign
	fmt.Printf("%o\n", testingNumber) // base 8
	fmt.Printf("%O\n", testingNumber) // base8 with leading 0
	fmt.Printf("%x\n", testingNumber) // base16 lowercase
	fmt.Printf("%X\n", testingNumber) // base16 upercase
	fmt.Printf("%#x\n", testingNumber) // base16 with leading 0x
	fmt.Printf("%4d\n", testingNumber) // pad with space
	fmt.Printf("%-4d\n", testingNumber) // pad with left spacing
	fmt.Printf("%04d\n", testingNumber) // pad woih zero

	fmt.Println("------------------------------------------------------")
	// string verb
	testingString := "Shuvo Halder"
	fmt.Printf("%s\n", testingString) // printing plain text
	fmt.Printf("%q\n", testingString) // print quota text
	fmt.Printf("%8s\n", testingString) // print plain 8 justify right
	fmt.Printf("%-8s\n", testingString) // print plain 8 justify left
	fmt.Printf("%x\n", testingString) // print the value of hex dump of byte volume
	fmt.Printf("% x\n", testingString) // printing the value of hex dump with space

	fmt.Println("------------------------------------------------------")
	// boolean verb
	var i = true
	var j = false
	fmt.Printf("%t\n", i) // for print true
	fmt.Printf("%t\n", j) // for print false

	fmt.Println("------------------------------------------------------")
	// float verb
	piValue := 3.1416
	fmt.Printf("%e\n", piValue) // notation with 'e' as expoint
	fmt.Printf("%f\n", piValue) // decimal point no expoint
	fmt.Printf("%2f\n", piValue) // default with precission 2
	fmt.Printf("%6.2f\n", piValue) // width 6, precission 2
	fmt.Printf("%g\n", piValue) // exponent as need, only nessary digit

	fmt.Println("------------------------------------------------------")
}
package main

import "fmt"

func main(){
	/* array define main stracture */
	/* array name := [length]dataType{values} */


	/* array define with defined array length */
	arrayStore := [5]int{10,20,30,50,4:100}
	fmt.Println("Array Print: ",arrayStore)
	fmt.Println("access 4: ",arrayStore[4])
	fmt.Println("---------------------------------------")


	/* array defined with inferred data length */
	inferredArray := [...]int{101,203,5}
	fmt.Println("Inferred array: ", inferredArray)
	fmt.Println("---------------------------------------")

	/* access element of an array */
	fmt.Println(arrayStore[1]);
	fmt.Println(inferredArray[2])
	fmt.Println("---------------------------------------")

	/* change element of an array */
	arrayStore[3] = 70
	fmt.Println("change number is: ",arrayStore[3])

	fmt.Println("---------------------------------------")

	/* array have initilazation as like partially , fully, not initilize */
	array1 := [5]string{} /* not Initialization */
	array2 := [5]int{10,20} /* partially init */
	array3 := [3]int{10,20,30} /* fully init */ 
	fmt.Println("array1: ",array1)
	fmt.Println("array2:",array2)
	fmt.Println("array3:",array3)

	fmt.Println("---------------------------------------")

	/* find the length of an array */
	fmt.Println("array length:",len(array1))
}
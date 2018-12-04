package main

import "fmt"

//This function intSeq returns another function,
// which we define anonymously in the body of intSeq.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	//This function value captures its own i value,
	// which will be updated each time we call nextInt.
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	newInts := intSeq()
	fmt.Println(newInts()) //1
}

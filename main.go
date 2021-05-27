/*
21.05.23
Learning GO language with nomad code

Using simple print function  (Hello World!)


add something package (21.05.23)


learning const variable & nomral variable (21.05.23)

add comment for github third (21.05.23)


21.05.27 Functions part One
--Naked Function
-- Defer

21.05.27 for loop
- range

*/

package main //package name (Necessary)
import (
	"fmt"
	"strings"
)

//func multiply(a int, b int) int { // a type, b type, result type (total 3 things)
func multiply(a, b int) int { // a, b type both int, result type (total 3 things)
	return a * b
}

//func functionName(parameter) (return value,...)
// func lenAndUpper(name string) (int, string) {

// 	return len(name), strings.ToUpper(name)
// }

//Already return variable is setting (like parameter setting)
//Naked return
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // After function close
	length = len(name)
	uppercase = strings.ToUpper(name)
	return

}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {

	//range give index
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	//

	//:= declared
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	total := 0
	for _, number := range numbers { //only get value, ignore index
		total += number //total = total + number
	}
	defer fmt.Println((total))
	return total
}

//var name bool = false
func main() {
	total := superAdd(1, 2, 3, 4, 5, 8)

	fmt.Println("Total: ", total)
	//for loop

}

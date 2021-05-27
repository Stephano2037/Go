/*
21.05.23
Learning GO language with nomad code

Using simple print function  (Hello World!)


add something package (21.05.23)


learning const variable & nomral variable (21.05.23)

add comment for github third (21.05.23)


21.05.27 Functions part One
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
func lenAndUpper(name string) (int, string) {

	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

//var name bool = false
func main() {

	//If you don't use your variable in Go, then compiler makes error
	//totalLength, _ := lenAndUpper("stephano") //ignore value with '_'
	//totalLength, uppperName := lenAndUpper("stephano")
	//fmt.Println(totalLength, uppperName)
	//fmt.Println(totalLength)
	//fmt.Println(multiply(2, 2))

	repeatMe("nico", "stephano", "dal", "marcy")

}

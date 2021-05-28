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


21.05.29 if else

*/

package main //package name (Necessary)
import (
	"fmt"
)

func canIDrink(age int) bool {
	//koreanAge := age + 2

	//if koreanAge < 18
	//can add variable in if
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {

	}

	return true

} //end of canIdrinkk function

//var name bool = false
func main() {
	fmt.Println(canIDrink(16))
	//for loop

}

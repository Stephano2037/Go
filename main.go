/*
21.05.23
Learning GO language with nomad code

Using simple print function  (Hello World!)


add something package (21.05.23)


learning const variable & nomral variable (21.05.23)

add comment for github third (21.05.23)
*/

package main //package name (Necessary)
import (
	"chanhyun/something"
	"fmt"
)

//var name bool = false
func main() {
	const name = "stephano" //
	//name = "cool"

	//	var names string = "nico"
	names := "nico" //find appropriate type itself ,축약형 대입
	//	names = "lynn" //can change the variable content
	fmt.Println(name)

	fmt.Println(names)
	something.SayHello()
}

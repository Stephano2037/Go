package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputValue() (a int, err error) {
	var n int
	_, error1 := fmt.Scanln(&n)
	if error1 != nil {
		stdin.ReadString('\n')
	}

	return n, error1
}

const standard_num int = 5

func main() {

	var Money int = 1000
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Start Game!!!")
	for {
		randvalue := (rand.Intn(standard_num) + 1)
		n, err := InputValue()
		if err != nil {
			fmt.Println("숫자를 입력하세요")

		} else if n < 1 || n > standard_num {
			fmt.Println("1과 5사이의 값으로 입력해주세요!")
		} else {
			if n == randvalue {
				Money += 500
				fmt.Println("축하합니다. 500원 추가!", Money)
			} else {
				Money -= 100
				fmt.Println("안타깝네요. 100원 차감", Money)
			}

			if Money <= 0 || Money >= 5000 {
				fmt.Println("게임 종료 ")
				break
			}
		}

	} //end of for
}

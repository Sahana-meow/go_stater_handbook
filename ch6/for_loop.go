package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	i = 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// 死循环
	i = 0
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
		if i > 10 {
			break
		}
	}

	fmt.Println("continue & break")
	// continue break 用于提前进入下一轮或者提前终止
	i = 0
	for {
		i++
		if i == 5 {
			continue
		}
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}

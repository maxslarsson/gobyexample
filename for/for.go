package _for

import "fmt"

func For() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	for j := ""; len(j) < 4; j += "a" {
		fmt.Println(j)
	}

	k := 4
	for {
		k -= 1
		fmt.Println(k)
		if k < 0 {
			break
		}
	}
}

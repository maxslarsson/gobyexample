package values

import "fmt"

func Values() {
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)
	fmt.Println("8.5 / 4 =", 8.5/4)

	fmt.Println("(2 + 4i) *  (3 + 2i) =", (2+4i)*(3+2i))

	fmt.Println("'A' + 'B' =", 'A'+'B')

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

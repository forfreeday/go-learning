package P01var

import "fmt"

func ForLoopTest() {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch j {
			case 4:
				fmt.Println(i, j)
				//break OuterLoop
				break
			case 5:
				fmt.Println(i, j)
				//break OuterLoop
				break OuterLoop
			}
			fmt.Println("forj")
		}
		fmt.Println("fori")
	}
	fmt.Println("out")
}

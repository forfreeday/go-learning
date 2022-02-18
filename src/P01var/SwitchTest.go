package p01var

import "fmt"

func SwitchTest() {
	/* 定义局部变量 */
	var grade string = "D"
	//var marks int = 90
	//
	//switch marks {
	//case 90:
	//	grade = "A"
	//case 80:
	//	grade = "B"
	//case 50, 60, 70:
	//	grade = "C"
	//default:
	//	grade = "D"
	//}

	for i := 0; i < 10; i++ {
		switch {
		case grade == "A":
			fmt.Printf("优秀!\n")
		case grade == "B", grade == "C":
			fmt.Printf("良好\n")
		case grade == "D":
			fmt.Printf("及格\n")
			//continue
			//break
			//return
		case grade == "F":
			fmt.Printf("不及格\n")
		default:
			fmt.Printf("差\n")
		}
		fmt.Printf("你的等级是 %s\n", grade)
	}

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {

}

func SwitchTest2(x interface{}) {
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}
}

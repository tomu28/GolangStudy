package Basics_Packages_variables_and_functions

import "fmt"

const Pi = 3.14

// 定数constは変数varのように、:=　を使って宣言できない
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

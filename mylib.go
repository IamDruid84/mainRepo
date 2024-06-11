package myLib

import "github.com/IamDruid84/myLib"

func main() {
	DoSomething(0, 1)
}

func DoSomething(do int, do2 int) int {
	myLib.DoSomething(1, 2)
	return do2 + do
}

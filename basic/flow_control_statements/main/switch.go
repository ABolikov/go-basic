package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switchExample()
	switchTrueExample()
}

// switch - так же как и в java/kotlin, за исключением того, что Go выполняет только выбранный случай, а не все последующие случаи
func switchExample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchTrueExample() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

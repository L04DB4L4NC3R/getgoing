package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	flag := true
	// continue, break, switch case,
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		} else if i == 1 {
			continue
		}
	}
	fmt.Println(flag)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}

}

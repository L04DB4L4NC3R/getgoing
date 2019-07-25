package main

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

// func main() {
// 	m1, m2 := 2, 3
// 	fmt.Println(m1, m2)
// 	swap(&m1, &m2)
// 	fmt.Println(m1, m2)
// }

// func main() {
// 	m1 := "my name"
// 	m2 := "name"
// 	fmt.Println(strings.Split(m1, " "), m1+m2)
// }

// func main() {

// 	m1 := 2
// 	m2 := 3
// 	fmt.Println(m1 + m2)
// }

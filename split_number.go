package main

// e.g: Split(17) return 7 . 10
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// func main() {
// 	fmt.Println(Split(17))
// }

package main

func SlicesAppend() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5, 6)
	printSlice(s)

	// use make
	z := make([]int, 0, 3)
	printSlice(z)

	z = append(z, 44)
	printSlice(z)

	z = append(z, 94, 12)
	printSlice(z)

	z = append(z, 1, 2, 3)
	printSlice(z)

	z = append(z, 1, 2, 3)
	printSlice(z)

	z = append(z, 99, 26, 87, 54, 11)
	printSlice(z)
}

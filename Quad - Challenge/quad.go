package piscine

import "fmt"

func makeLine(x int, first string, middle string, last string) {
	fmt.Print(first)
	for i := 1; i < x-1; i++ {
		fmt.Print(middle)
	}

	fmt.Println(last)
}

func QuadA(x int, y int) {
	if x == 1 && y == 1 {
		fmt.Print("o")
		return
	}

	// If x is one
	if x == 1 {
		fmt.Println("o")
		for i := 0; i < y-2; i++ {
			fmt.Println("|")
		}
		fmt.Println("o")
		return
	}

	// if y is one

	if y == 1 {
		makeLine(x, "o", "-", "o")
		return
	}

	makeLine(x, "o", "-", "o")

	for i := 2; i < y; i++ {
		makeLine(x, "|", " ", "|")

	}
	makeLine(x, "o", "-", "o")
}

func QuadB(x int, y int) {
	if x == 1 && y == 1 {
		fmt.Print("/")
		return
	}

	// If x is one, prints a line down
	if x == 1 {
		fmt.Println("/")
		for i := 0; i < y-2; i++ {
			fmt.Println("*")
		}
		fmt.Println("\\")
		return
	}

	// if y is one, prints a line across

	if y == 1 {
		makeLine(x, "/", "*", "\\")
		return
	}
	// Main stuff

	makeLine(x, "/", "*", "\\")

	for i := 2; i < y; i++ {
		makeLine(x, "*", " ", "*")

	}
	makeLine(x, "\\", "*", "/")
}

func QuadC(x int, y int) {
	if x == 1 && y == 1 {
		fmt.Print("A")
		return
	}

	// If x is one, prints a line down
	if x == 1 {
		fmt.Println("A")
		for i := 0; i < y-2; i++ {
			fmt.Println("B")
		}
		fmt.Println("C")
		return
	}

	// if y is one, prints a line across

	if y == 1 {
		makeLine(x, "A", "B", "A")
		return
	}
	// Main stuff

	makeLine(x, "A", "B", "A")

	for i := 2; i < y; i++ {
		makeLine(x, "B", " ", "B")

	}
	makeLine(x, "C", "B", "C")
}

func QuadD(x int, y int) {
	if x == 1 && y == 1 {
		fmt.Print("A")
		return
	}

	// If x is one, prints a line down
	if x == 1 {
		fmt.Println("A")
		for i := 0; i < y-2; i++ {
			fmt.Println("B")
		}
		fmt.Println("A")
		return
	}

	// if y is one, prints a line across

	if y == 1 {
		makeLine(x, "A", "B", "C")
		return
	}
	// Main stuff

	makeLine(x, "A", "B", "C")

	for i := 2; i < y; i++ {
		makeLine(x, "B", " ", "B")

	}
	makeLine(x, "A", "B", "C")
}

func QuadE(x int, y int) {
	if x == 1 && y == 1 {
		fmt.Print("A")
		return
	}

	// If x is one, prints a line down
	if x == 1 {
		fmt.Println("A")
		for i := 0; i < y-2; i++ {
			fmt.Println("B")
		}
		fmt.Println("C")
		return
	}

	// if y is one, prints a line across

	if y == 1 {
		makeLine(x, "A", "B", "C")
		return
	}
	// Main stuff

	makeLine(x, "A", "B", "C")

	for i := 2; i < y; i++ {
		makeLine(x, "B", " ", "B")

	}
	makeLine(x, "C", "B", "A")
}

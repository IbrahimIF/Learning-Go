package main

import (
	"fmt"
	"piscine"
	//"github.com/01-edu/z01"
)

func main() {

	fmt.Println("\n--- UltimatePointOne ---")
	var val1 int = 0
	var ptr1 *int = &val1
	var ptr2 **int = &ptr1
	var ptr3 ***int = &ptr2
	fmt.Printf("Before UltimatePointOne: val1 = %d\n", val1)
	piscine.UltimatePointOne(ptr3)
	fmt.Printf("After UltimatePointOne: val1 = %d (Expected: 1)\n", val1)

	fmt.Println("\n--- UltimateDivMod ---")
	var a, b int
	a = 10
	b = 3
	fmt.Printf("Before UltimateDivMod: a = %d, b = %d\n", a, b)
	piscine.UltimateDivMod(&a, &b)
	fmt.Printf("After UltimateDivMod: a (quotient) = %d (Expected: 3), b (remainder) = %d (Expected: 1)\n", a, b)

	a = 25
	b = 5
	fmt.Printf("Before UltimateDivMod: a = %d, b = %d\n", a, b)
	piscine.UltimateDivMod(&a, &b)
	fmt.Printf("After UltimateDivMod: a (quotient) = %d (Expected: 5), b (remainder) = %d (Expected: 0)\n", a, b)

	fmt.Println("\n--- Swap ---")
	var x, y int
	x = 100
	y = 200
	fmt.Printf("Before Swap: x = %d, y = %d\n", x, y)
	piscine.Swap(&x, &y)
	fmt.Printf("After Swap: x = %d (Expected: 200), y = %d (Expected: 100)\n", x, y)

	fmt.Println("\n--- StrRev ---")
	testStrRev1 := "Hello"
	fmt.Printf("StrRev(\"%s\"): \"%s\" (Expected: olleH)\n", testStrRev1, piscine.StrRev(testStrRev1))
	testStrRev2 := "GoLang"
	fmt.Printf("StrRev(\"%s\"): \"%s\" (Expected: gnaLoG)\n", testStrRev2, piscine.StrRev(testStrRev2))
	testStrRev3 := ""
	fmt.Printf("StrRev(\"%s\"): \"%s\" (Expected: \"\")\n", testStrRev3, piscine.StrRev(testStrRev3))
	testStrRev4 := "a"
	fmt.Printf("StrRev(\"%s\"): \"%s\" (Expected: a)\n", testStrRev4, piscine.StrRev(testStrRev4))

	fmt.Println("\n--- StrLen ---")
	testStrLen1 := "Hello"
	fmt.Printf("StrLen(\"%s\"): %d (Expected: 5)\n", testStrLen1, piscine.StrLen(testStrLen1))
	testStrLen2 := "Golang is fun!"
	fmt.Printf("StrLen(\"%s\"): %d (Expected: 14)\n", testStrLen2, piscine.StrLen(testStrLen2))
	testStrLen3 := ""
	fmt.Printf("StrLen(\"%s\"): %d (Expected: 0)\n", testStrLen3, piscine.StrLen(testStrLen3))

	fmt.Println("\n--- RecursivePower ---")
	fmt.Printf("RecursivePower(2, 3): %d (Expected: 8)\n", piscine.RecursivePower(2, 3))
	fmt.Printf("RecursivePower(5, 0): %d (Expected: 1)\n", piscine.RecursivePower(5, 0))
	fmt.Printf("RecursivePower(4, 1): %d (Expected: 4)\n", piscine.RecursivePower(4, 1))
	fmt.Printf("RecursivePower(2, -1): %d (Expected: 0)\n", piscine.RecursivePower(2, -1))
	fmt.Printf("RecursivePower(-3, 2): %d (Expected: 9)\n", piscine.RecursivePower(-3, 2))
	fmt.Printf("RecursivePower(-2, 3): %d (Expected: -8)\n", piscine.RecursivePower(-2, 3))

	fmt.Println("\n--- RecursiveFactorial ---")
	fmt.Printf("RecursiveFactorial(4): %d (Expected: 24)\n", piscine.RecursiveFactorial(4))
	fmt.Printf("RecursiveFactorial(0): %d (Expected: 1)\n", piscine.RecursiveFactorial(0))
	fmt.Printf("RecursiveFactorial(1): %d (Expected: 1)\n", piscine.RecursiveFactorial(1))
	fmt.Printf("RecursiveFactorial(-5): %d (Expected: 0)\n", piscine.RecursiveFactorial(-5))
	fmt.Printf("RecursiveFactorial(22): %d (Expected: 0) (Due to > 21 constraint)\n", piscine.RecursiveFactorial(22))
	fmt.Printf("RecursiveFactorial(7): %d (Expected: 5040)\n", piscine.RecursiveFactorial(7))

	fmt.Println("\n--- PrintStr ---")
	fmt.Print("PrintStr(\"Hello World!\"): ")
	piscine.PrintStr("Hello World!")
	fmt.Print("PrintStr(\"Test\"): ")
	piscine.PrintStr("Test")
	fmt.Println()

	fmt.Println("\n--- PrintComb2 ---")
	fmt.Println("PrintComb2 (Output below):")
	piscine.PrintComb2()

	fmt.Println("\n--- PrintComb ---")
	fmt.Println("PrintComb (Output below):")
	piscine.PrintComb()

	fmt.Println("\n--- PointOne ---")
	var num int = 5
	fmt.Printf("Before PointOne: num = %d\n", num)
	piscine.PointOne(&num)
	fmt.Printf("After PointOne: num = %d (Expected: 1)\n", num)

	fmt.Println("\n--- IterativePower ---")
	fmt.Printf("IterativePower(2, 3): %d (Expected: 8)\n", piscine.IterativePower(2, 3))
	fmt.Printf("IterativePower(5, 0): %d (Expected: 1)\n", piscine.IterativePower(5, 0))
	fmt.Printf("IterativePower(4, 1): %d (Expected: 4)\n", piscine.IterativePower(4, 1))
	fmt.Printf("IterativePower(2, -1): %d (Expected: 0)\n", piscine.IterativePower(2, -1))
	fmt.Printf("IterativePower(-3, 2): %d (Expected: 9)\n", piscine.IterativePower(-3, 2))
	fmt.Printf("IterativePower(-2, 3): %d (Expected: -8)\n", piscine.IterativePower(-2, 3))

	fmt.Println("\n--- IterativeFactorial ---")
	fmt.Printf("IterativeFactorial(4): %d (Expected: 24)\n", piscine.IterativeFactorial(4))
	fmt.Printf("IterativeFactorial(0): %d (Expected: 1)\n", piscine.IterativeFactorial(0))
	fmt.Printf("IterativeFactorial(1): %d (Expected: 1)\n", piscine.IterativeFactorial(1))
	fmt.Printf("IterativeFactorial(-5): %d (Expected: 0)\n", piscine.IterativeFactorial(-5))
	fmt.Printf("IterativeFactorial(22): %d (Expected: 0) (Due to > 21 constraint)\n", piscine.IterativeFactorial(22))
	fmt.Printf("IterativeFactorial(7): %d (Expected: 5040)\n", piscine.IterativeFactorial(7))

	fmt.Println("\n--- IsNegative ---")
	fmt.Print("IsNegative(5): ")
	piscine.IsNegative(5)
	fmt.Print("IsNegative(-5): ")
	piscine.IsNegative(-5)
	fmt.Print("IsNegative(0): ")
	piscine.IsNegative(0)

	fmt.Println("\n--- DivMod ---")
	var valA, valB, divResult, modResult int
	valA = 10
	valB = 3
	fmt.Printf("Before DivMod: valA = %d, valB = %d\n", valA, valB)
	piscine.DivMod(valA, valB, &divResult, &modResult)
	fmt.Printf("After DivMod: div = %d (Expected: 3), mod = %d (Expected: 1)\n", divResult, modResult)

	valA = 25
	valB = 5
	fmt.Printf("Before DivMod: valA = %d, valB = %d\n", valA, valB)
	piscine.DivMod(valA, valB, &divResult, &modResult)
	fmt.Printf("After DivMod: div = %d (Expected: 5), mod = %d (Expected: 0)\n", divResult, modResult)
}
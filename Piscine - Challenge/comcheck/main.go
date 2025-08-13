package main

import (
	"fmt"
	"os"
)

func main() {
	arr := [...]string{"01", "galaxy", "galaxy 01"}

	for i := 1; i <= len(os.Args)-1; i++ {
		for j := 0; j <= len(arr)-1; j++ {
			if string(os.Args[i]) == arr[j] {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}

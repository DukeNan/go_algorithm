package main

import "fmt"

func SelectSort(data []int) {
	llen := len(data)
	for i := 0; i < llen; i++ {
		tmp := data[i]
		flag := i
		for j := i + 1; j < llen; j++ {
			if data[j] < tmp {
				tmp = data[j]
				flag = j
			}
		}
		if flag != i {
			data[flag] = data[i]
			data[i] = tmp
		}
	}
}
func main() {
	data := []int{5, 4, 8, 7, 6, 0, 1, 3, 2}
	SelectSort(data)
	fmt.Printf("data: %v\n", data)

}

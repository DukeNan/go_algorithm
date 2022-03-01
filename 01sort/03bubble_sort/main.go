/*
 * @Author: shaun
 * @Date: 2022-03-01 16:50:55
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-01 17:28:25

  冒泡排序

  基本思想：

比较相邻的元素。如果第一个比第二个大（升序），就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

*/
package main

import "fmt"

func BubbleSort(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}

	}
}

func main() {
	data := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}
	// data := []int{5}
	BubbleSort(data)
	fmt.Printf("data: %v\n", data)
}

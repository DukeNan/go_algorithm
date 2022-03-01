/*
 * @Author: shaun
 * @Date: 2022-03-01 15:13:36
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-01 15:35:55

选择排序算法

基本思想：

首先在未排序的数列中找到最小(or最大)元素，
然后将其存放到数列的起始位置；
接着，再从剩余未排序的元素中继续寻找最小(or最大)元素，
然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
*/

package main

import "fmt"

func SelectSort(data []int) {
	llen := len(data)
	for i := 0; i < llen; i++ {
		// 初始最小值
		tmp := data[i]
		// 记录剩余未排序部分最小元素的下标
		flag := i
		// 交换剩余未排序部分最小元素的位置
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
		fmt.Printf("i: %v, data:%v\n", i, data)
	}

}
func SelectSort_1(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}
func main() {
	data := []int{13, 65, 97, 76, 38, 27, 49}
	// data := []int{20}
	SelectSort_1(data)
	fmt.Printf("data: %v\n", data)

}

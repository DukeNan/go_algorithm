/*
 * @Author: shaun
 * @Date: 2022-03-01 16:11:33
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-01 16:46:49

 插入排序

基本思想：
它的工作原理是通过构建有序序列，
对于未排序数据，在已排序序列中从后向前扫描，
找到相应位置并插入。插入排序在实现上，
在从后向前扫描过程中，
需要反复把已排序元素逐步向后挪位，
为最新元素提供插入空间。

*/
package main

import "fmt"

func InsertSort(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		// 存放当前需要排序的元素值
		back := data[i]
		j := i - 1
		// 已排序部分每次向左移动，小于back向右移动一位
		for j >= 0 && back < data[j] {
			data[j+1] = data[j]
			j--
		}
		// 将back值放到已排序部分最小位置
		data[j+1] = back
	}
}

func InsertSort_1(data []int) {
	length := len(data)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

func main() {
	data := []int{7, 3, 19, 40, 4, 7, 1}
	InsertSort_1(data)
	fmt.Printf("data: %v\n", data)
}

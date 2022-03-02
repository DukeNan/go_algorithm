/*
 * @Author: shaun
 * @Date: 2022-03-02 15:44:46
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-02 20:15:15

 快速排序
1.从数列中挑出一个元素，称为"基准"（pivot）

2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。

3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

双指针法，填坑交换位置，初始化基准位置为坑位。

*/
package main

import "fmt"

func quickSort(array []int, left, right int) {
	// 递归退出条件
	if left >= right {
		return
	}
	// 设置基准为起始位置
	pivot := array[left]
	//i：从左向右移动的指针
	// j: 从右向左移动的指针
	var i, j = left, right
	for i < j {
		//  如果i与j未重合，j指向的元素不比基准元素小，则j向左移动
		for i < j && array[j] >= pivot {
			j--
		}
		// 将j指向的元素放到i的位置上
		array[i] = array[j]

		// 如果i与j未重合，i指向的元素比基准元素小，则i向右移动
		for i < j && array[i] < pivot {
			i++
		}
		// 将i指向的元素放到j的位置上
		array[j] = array[i]
	}
	// 退出循环后，i与j重合，此时所指位置为基准元素的正确位置
	// 将基准元素放到该位置
	array[i] = pivot

	// 对基准元素左边的子序列进行快速排序
	quickSort(array, left, i-1)
	// 对基准元素右边的子序列进行快速排序
	quickSort(array, i+1, right)
}

func main() {
	array := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)

}

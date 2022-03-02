/*
 * @Author: shaun
 * @Date: 2022-03-02 15:04:16
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-02 15:38:07
  归并排序

  基本思想：
将数组分解最小之后，然后合并两个有序数组，
基本思路是比较两个数组的最前面的数，谁小就先取谁，取了后相应的指针就往后移一位。
然后再比较，直至一个数组为空，最后把另一个数组的剩余部分复制过来即可。

算法步骤：
对于给定的一组记录（假设有n个记录），首先将每两个相邻的长度为1的子序列进行归并，得到n/2（向上取整）个长度为2或1的有序子序列，
再将其归并，反复执行过程，直到得到一个有序序列为止。
*/
package main

import "fmt"

// 自顶向下归并排序，排序范围在[begin, end]数组
func mergeSort(array []int, begin, end int) {
	if end-begin <= 1 {
		return
	}
	// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
	mid := begin + (end-begin+1)/2
	// 先将左边排序好
	mergeSort(array, begin, mid)
	// 再将右边排序好
	mergeSort(array, mid, end)
	// 两个有序数组进行合并
	merge(array, begin, mid, end)

}

func merge(array []int, begin, mid, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

func main() {
	array := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}
	mergeSort(array, 0, len(array))
	fmt.Printf("array: %v\n", array)
}

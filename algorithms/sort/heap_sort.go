// package sort
// @file: heap_sort.go
// @description: 堆排序
// @author: SaltFish
// @date: 2020/9/12
package sort

func adjustHeap(i, lef int, nums []int) {
	tmp := nums[i]
	for k := i*2 + 1; k < lef; k = k*2 + 1 {
		if k+1 < lef && nums[k] < nums[k+1] {
			k++
		}
		if nums[k] > tmp {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = tmp
}

func heapSort(nums []int) {
	// 构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(i, len(nums), nums)
	}
	// 调整堆结构，交换堆顶元素与末尾元素
	for j := len(nums) - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		adjustHeap(0, j, nums)
	}
}

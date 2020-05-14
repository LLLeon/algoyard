/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-05-14
+************************************************************************/

package heap

import (
	"testing"
)

// 求一个无序数组中的中位数.
// 首先将数组的前 (n+1)／2 个元素建立一个最小堆.
//
// 然后, 对于下一个元素, 和堆顶的元素比较, 如果小于等于, 丢弃之, 接着看下一个元素.
// 如果大于, 则用该元素取代堆顶, 再调整堆, 接着看下一个元素. 重复这个步骤, 直到数组为空.
//
// 当数组都遍历完了, 那么, 堆顶的元素即是中位数.
//
// 可以看出, 长度为 (n＋1)／2 的最小堆是解决方案的精华之处.

// 这里用大顶堆实现, 注意边界条件即可.
func TestNewHeap(t *testing.T) {
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
	list := []int{2, 4, 3, 1, 6, 7, 5, 9, 8}
	heap := NewHeap(10)

	// 计算要插入堆中的元素数量
	n := len(list)/2 + 1

	// 把数组前半部分插入小顶堆
	for i := 0; i < n; i++ {
		heap.Insert(list[i])
	}

	for j := n; j < len(list); j++ {
		if list[j] < heap.items[1] {
			heap.RemoveMax()
			heap.Insert(list[j])
		}
	}

	// 取出堆顶元素即数组中位数.
	median := heap.items[1]
	if median != 5 {
		t.Error("wrong result")
	} else {
		t.Log("median:", median)
	}
}

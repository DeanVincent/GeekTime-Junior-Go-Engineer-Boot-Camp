package slice

import (
	"fmt"
)

// DeleteAt 删除数组中指定索引的元素
// 参数：源数组，需要删除元素的索引
// 返回值：删除指定元素后的数组，被删除的元素的值，错误
func DeleteAt[T any](src []T, idx int) ([]T, T, error) {
	if idx < 0 || idx >= len(src) {
		var zero T
		return nil, zero, newErrIndexOutOfRange(len(src), idx)
	}
	deleted := src[idx]
	for i := idx; i+1 < len(src); i++ {
		src[i] = src[i+1]
	}
	src = src[:len(src)-1]
	src = MaybeShrink(src)
	return src, deleted, nil
}

func newErrIndexOutOfRange(len, idx int) error {
	return fmt.Errorf("索引超出范围，长度 %v，索引 %v\n", len, idx)
}

func MaybeShrink[T any](src []T) []T {
	srcLen, srcCap := len(src), cap(src)
	dstCap, change := calCapacity(srcLen, srcCap)
	if !change {
		return src
	}
	dst := make([]T, srcLen, dstCap)
	copy(dst, src)
	return dst
}

func calCapacity(len, cap int) (int, bool) {
	if cap <= 64 {
		return len, false
	}
	if cap >= 1280 && float64(cap)/float64(len) >= float64(25)/16 {
		return cap * 4 / 5, true
	}
	if cap < 1280 && cap/len >= 4 {
		return cap / 2, true
	}
	return cap, false
}

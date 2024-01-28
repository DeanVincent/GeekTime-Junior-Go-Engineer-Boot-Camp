package test

import (
	"fmt"
	slice "github/GeekTime-Junior-Go-Engineer-Boot-Camp/homework/week1"
	"testing"
)

func TestNormal(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4}
	fmt.Printf("删除前 %v 长度 %v 容量 %v \n", s1, len(s1), cap(s1))
	if s2, _, err := slice.DeleteAt(s1, 2); err != nil {
		t.Fatalf("删除出错 %v \n", err)
	} else {
		fmt.Printf("删除后 %v 长度 %v 容量 %v \n", s2, len(s2), cap(s2))
	}
}

func TestIndexOutOfRange(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4}
	if _, _, err := slice.DeleteAt(s1, len(s1)); err != nil {
		return
	}
	t.Fatalf("数组 %v 删除 %v 位置的元素，本该出的 ErrIndexOutOfRange 没有出现", s1, len(s1))
}

func TestShrink(t *testing.T) {
	s1 := make([]int, 3125, 3125)
	for i := 0; i < len(s1); i++ {
		s1[i] = i
	}
	oldCap := cap(s1)
	fmt.Printf("在删除前，数组的长度 %v 容量 %v 地址 %p \n", len(s1), cap(s1), s1)

	for len(s1) > 0 {
		s1, _, _ = slice.DeleteAt(s1, len(s1)-1)
		//fmt.Printf("Delete happended, len %v cap %v address %p \n", len(s1), cap(s1), s1)
		s1 = slice.MaybeShrink(s1)
		if cap(s1) != oldCap {
			fmt.Printf("数组在删除时发生了缩容, 缩容后的数组的长度 %v 容量 %v 地址 %p \n", len(s1), cap(s1), s1)
			oldCap = cap(s1)
		}
	}
}

package main

import (
	"strings"
)

// string从后遍历，用cnt记录#的个数，当遇到正常节点时，#的个数-2，并将该节点转化成#，cnt+1，整体即为cnt-1
// 当出现#的个数不足2时，即false，最终也须保证cnt为1。
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")

	length := len(nodes)
	cnt := 0
	for i := length - 1; i >= 0; i-- {
		if nodes[i] == "#" {
			cnt++
		} else {
			if cnt >= 2 {
				cnt--
			} else {
				return false
			}
		}
	}

	return cnt == 1
}

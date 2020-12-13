package treesort

import "strconv"

// 二叉树 struct
type tree struct {
	value       int
	left, right *tree
}

// 排序 binary tree
func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

// 对 binary tree 上所有元素 进行递归排序
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left) // 递归后 []int 不具有引用 作用
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

// 向 binary tree 添加元素
func add(t *tree, val int) *tree {
	if t == nil {
		t = new(tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = add(t.left, val) // 大于root node 上 num, 添加到left node
	} else {
		t.right = add(t.right, val) // 小于root node 上 num, 添加到left node
	}

	return t
}

// 将节点上的number 转为 string 的String方法
func (t *tree) String() string {
	var s string

	if t == nil {
		return s
	}

	s += t.left.String()
	s += strconv.Itoa(t.value)
	s += t.right.String()

	return s
}

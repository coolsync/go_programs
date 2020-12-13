package treesort

// uses a binary tree to implement an insertion sort
// 1. define binary tree struct
type tree struct {
	value       int
	left, right *tree
}

// 2. sort value in tree struct
func Sort(s []int) {
	var root *tree // 根节点
	for _, v := range s {
		root = add(root, v)
	}

	appendValues(s[:0], root)
}

// 3. 将slice的值 添加到二叉树上
func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if t.value > v {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

// 4. 对二叉树递归排序, 获取所有值追加到slice
func appendValues(s []int, t *tree) []int {
	if t != nil {
		s = appendValues(s, t.left)
		s = append(s, t.value)
		s = appendValues(s, t.right)
	}
	return s
}

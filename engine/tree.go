package engine

import "strings"

type GoodNode struct {
	pattern string
	part    string
	child   []*GoodNode
	isWild  bool
}

func NewGoodNode() *GoodNode {
	child := make([]*GoodNode, 0)
	return &GoodNode{
		child: child,
	}
}

func NewRoot() *GoodNode {
	child := make([]*GoodNode, 0)
	return &GoodNode{
		child:   child,
		part:    "/",
		pattern: "/",
	}
}

// 插入函数
// 实现：node将part插入到nodechild中。
// 1. 如果存在，则不处理
// 2. 如果不存在,则新建节点
// 3. nodechild继续执行insert(next part)
func (n *GoodNode) insert(pattern string, paths []string, height int) {
	if len(paths) == 0 {
		return
	}
	if len(paths) == height {
		n.pattern = pattern
		return
	}

	part := paths[height]

	// 遍历子节点是否存在part
	isCreate := true
	var child *GoodNode
	for _, child = range n.child {
		if child.part == part {
			isCreate = false
			break
		}
	}
	if isCreate {
		child = NewGoodNode()
		child.part = part
		if part[0] == '*' || part[0] == ':' {
			child.isWild = true
		}
		n.child = append(n.child, child)

	}
	child.insert(pattern, paths, height+1)
}

func parsePath(param string) []string {
	out := make([]string, 0)
	if param == "" {
		return out
	}
	for _, item := range strings.Split(param, "/") {
		if item != "" {
			out = append(out, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return out
}

func searchChildALL(n *GoodNode, part string) []*GoodNode {
	out := make([]*GoodNode, 0)
	for _, child := range n.child {
		if child.isWild {
			out = append(out, child)
		} else {
			if child.part == part {
				out = append(out, child)
			}
		}
	}
	return out
}

// search
// 需求： 查询pattern对应的节点
// 实现： node调用search查找part对应的节点
// 1. 如何确定最终节点
// 2. 如何设计获取part
func (n *GoodNode) search(pattern string, paths []string, height int) *GoodNode {
	if len(paths) == 0 {
		return NewRoot()
	}
	// 最后的节点直接返回
	// 这里少考虑了一种情况就是*不用完全遍历path
	if height == len(paths) || n.part[0] == '*' {
		if n.pattern != "" {
			return n
		}
		return nil
	}

	// 寻找n的子节点中所有符合path的所有集合
	childs := searchChildALL(n, paths[height])
	for _, child := range childs {
		out := child.search(pattern, paths, height+1)
		if out != nil {
			return out
		}
	}
	return nil
}

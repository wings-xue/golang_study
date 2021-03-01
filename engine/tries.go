package engine

import (
	"strings"
)

// 路由支持两种匹配

// "/hello/:name"
// $ curl "http://localhost:9999/hello/geektutu"
// hello geektutu, you're at /hello/geektutu

// "/assets/*filepath"
// $ curl "http://localhost:9999/assets/css/geektutu.css"
// {"filepath":"css/geektutu.css"}

// 输入  /hello/:name   输出: /hello/geektutu
//

const (
	WildSearch = 0
	MATCH      = 1
)

type Node struct {
	part     string  // 路由中的一部分，例如 :lang
	children []*Node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true

}

type NodeInfo struct {
	Node  *Node
	Param map[string]string
}

func NewNode() *Node {
	return &Node{
		part:     "/",
		children: make([]*Node, 0),
		isWild:   false,
	}
}

func removeSplitSlice(s []string) []string {
	out := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			out = append(out, s[i])
		}
	}
	return out
}

func (n *Node) Insert(path string) {
	tree := n
	for _, part := range strings.Split(path, "/") {
		if part != "" {
			isAddChile := true
			// 直接从孩子开始，因为根节点已经是/
			for _, node := range tree.children {
				if part == node.part {
					tree = node
					isAddChile = false
					break
				}
			}
			if isAddChile {
				newNode := NewNode()
				newNode.part = part
				tree.children = append(tree.children, newNode)
				tree = newNode
			}

		}

	}
}

// 返回路由对应的节点
// 寻找path是否路过当前节点的子节点中
//
func (n *Node) Search(path string) *NodeInfo {

	if path == "" || path == "/" {
		return &NodeInfo{
			Node:  n,
			Param: make(map[string]string),
		}
	}
	for _, node := range n.children {
		if param, ok := isMatchNode(path, node); ok {
			p := removeSplitSlice(strings.Split(path, "/"))
			path = strings.Join(p[1:], "/")
			// 判断是否是最后一个节点，感觉还是没有解耦
			if path == "" || strings.Index(node.part, "*") != -1 {
				return &NodeInfo{
					Node:  node,
					Param: param,
				}
			}
			out := node.Search(path)
			if out != nil {
				return out
			}

		}
	}

	return nil
}

// 匹配路由和当前节点
// 匹配当前路由和当前节点是否存在映射关系
//
func isMatchNode(path string, node *Node) (map[string]string, bool) {
	out := map[string]string{}
	if path == "" || path == "/" {
		return out, true
	}

	part := node.part
	switch {
	case strings.Index(part, ":") != -1:
		if strings.Index("/", path) == -1 {
			out[part[1:]] = path
			return out, true
		}
	case strings.Index(part, "*") != -1:

		out[part[1:]] = path
		return out, true
	// 正常情况匹配节点
	default:
		part := removeSplitSlice(strings.Split(path, "/"))

		return out, part[0] == node.part

	}

	return out, false
}

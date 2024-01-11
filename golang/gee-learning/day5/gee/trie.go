package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string // 完整路由
	part     string // 路由当前待匹配部分
	isWild   bool
	children []*node
}

func (n *node) String() string {
	return fmt.Sprintf(
		"node{pattern=%s, part=%s, isWild=%t}",
		n.pattern, n.part, n.isWild)
}

func (n *node) search(parts []string, height int) *node {
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		// 标明调用该方法的节点是一个路由的叶节点
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			// pattern 为空，说明此处不（一定）是叶节点
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// ---- Helper functions ----
func (n *node) matchChild(part string) *node {
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			return ch
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	res := make([]*node, 0)
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			res = append(res, ch)
		}
	}
	return res
}

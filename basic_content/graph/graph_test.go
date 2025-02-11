package graph

import (
	"fmt"
	"testing"
)

func TestGraph_bfs(t *testing.T) {
	// 创建一个包含 6 个节点的图
	g := NewGraph(6)

	// 添加边（无向图）
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 5)

	// 预期路径: 0 --> 2 --> 3 --> 4 --> 5
	//expected := "0 --> 2 --> 3 --> 4 --> 5 --> "

	// 捕获 BFS 输出
	g.bfs(0, 5)
	fmt.Println()
	g.dfs(0, 5)
}

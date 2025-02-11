package graph

import "fmt"

// Graph 无向图的邻接表实现与重要操作
type Graph struct {
	// 顶点个数
	v int
	// 邻接表
	adj []*LinkedList
	// 全局变量
	found bool
}

// 深度搜索算法
func (g *Graph) dfs(s, v int) {
	visited := make([]bool, g.v)
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}

	g.recurDfs(s, v, visited, prev)
	print(prev, s, v)
}

func (g *Graph) recurDfs(s int, v int, visited []bool, prev []int) {
	if g.found {
		return
	}

	visited[s] = true

	if s == v {
		// 访问到出发点
		g.found = true
		return
	}

	for i := 0; i < g.adj[s].size; i++ {
		q := g.adj[s].Get(i)
		if !visited[q] {
			prev[q] = s
			g.recurDfs(q, v, visited, prev)
		}
	}
}

// 广度搜索算法，打印两节点之间的路径
func (g *Graph) bfs(s, t int) {
	if s == t {
		return
	}

	visited := map[int]bool{}
	visited[s] = true
	queue := Queue{}
	queue.Enqueue(s)
	prev := make([]int, g.v)

	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}

	for queue.Size() > 0 {
		e := queue.Dequeue()
		if g.adj[e] == nil { // 防止访问 nil
			continue
		}

		for i := 0; i < g.adj[e].Size(); i++ {
			// 广度遍历
			q := g.adj[e].Get(i)
			if q == -1 {
				continue
			}

			if !visited[q] {
				prev[q] = e
				if q == t {
					print(prev, s, t)
					return
				}
				visited[q] = true
				queue.Enqueue(q)
			}

		}
	}

}

func print(prev []int, s, t int) {
	if t == -1 {
		return
	}
	if s != t {
		print(prev, s, prev[t]) // 递归打印路径
	}
	fmt.Printf("%d --> ", t) // 逐步输出路径
}

func (g *Graph) addEdge(s, t int) {
	g.adj[s].Add(t)
	g.adj[t].Add(s)
}

func NewGraph(v int) *Graph {
	g := &Graph{
		v:   v,
		adj: make([]*LinkedList, v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = &LinkedList{} // 这里初始化每个邻接表
	}
	return g
}

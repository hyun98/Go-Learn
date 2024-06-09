package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	M       int
	R       int
	graph   *Graph
	visited [100001]bool
	answer  [100001]int
	order   = 1
)

type Graph struct {
	adj [][]int
}

func NewGraph() *Graph {
	return &Graph{
		adj: make([][]int, N+1),
	}
}

func (g *Graph) AddEdge(src, dest int) {
	g.adj[src] = append(g.adj[src], dest)
}

func (g *Graph) SortList() {
	for i := range g.adj {
		sort.Slice(g.adj[i], func(q, w int) bool {
			return g.adj[i][q] < g.adj[i][w]
		})
	}
}

func (g *Graph) dfs(now int) {
	visited[now] = true
	answer[now] = order
	order++

	for i := range g.adj[now] {
		next := g.adj[now][i]
		if visited[next] {
			continue
		}
		g.dfs(next)
	}
}

func main() {
	defer writer.Flush()
	line, _, _ := reader.ReadLine()
	nums := strings.Fields(string(line))
	N, _ = strconv.Atoi(nums[0])
	M, _ = strconv.Atoi(nums[1])
	R, _ = strconv.Atoi(nums[2])

	graph = NewGraph()

	for i := 0; i < M; i++ {
		line, _, _ = reader.ReadLine()
		nums = strings.Fields(string(line))
		u, _ := strconv.Atoi(nums[0])
		v, _ := strconv.Atoi(nums[1])

		graph.AddEdge(u, v)
		graph.AddEdge(v, u)
	}

	graph.SortList()
	graph.dfs(R)

	for i := 1; i <= N; i++ {
		writer.WriteString(strconv.Itoa(answer[i]) + "\n")
	}
}

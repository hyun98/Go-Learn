package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type P struct {
	dst int
	wei int
}

type Item struct {
	v    int
	cost int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	V      int
	E      int
	K      int
	adj    [20001][]P
	dist   [20001]int
)

func dijkstra() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(&Item{K, 0})
	dist[K] = 0

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*Item)

		if dist[now.v] < now.cost {
			continue
		}

		for _, p := range adj[now.v] {
			next := p.dst
			nextCost := now.cost + p.wei
			if dist[next] > nextCost {
				dist[next] = nextCost
				heap.Push(&pq, &Item{next, nextCost})
			}
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &V, &E, &K)

	for i := 0; i < E; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		adj[u] = append(adj[u], P{v, w})
	}

	for i := 1; i <= V; i++ {
		dist[i] = int(1e9 + 7)
	}

	dijkstra()

	for i := 1; i <= V; i++ {
		if dist[i] == int(1e9+7) {
			writer.WriteString("INF")
		} else {
			writer.WriteString(strconv.Itoa(dist[i]))
		}
		writer.WriteByte('\n')
	}
}

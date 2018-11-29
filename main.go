package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/yourbasic/graph"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap. The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	V int
	E int
	g *graph.Mutable
}

type Result struct {
	x       []int
	k       int
	elapsed time.Duration
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createGraphFromFile(filepath string) (g Graph) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {
		case 'p':
			tokens := strings.Split(line, " ")
			g.V, err = strconv.Atoi(tokens[2])
			if err != nil {
				log.Fatal(err)
			}
			g.E, err = strconv.Atoi(tokens[3])
			if err != nil {
				log.Fatal(err)
			}
			g.g = graph.New(g.V)
			break
		case 'e':
			tokens := strings.Split(line, " ")
			e1, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal(err)
			}
			e2, err := strconv.Atoi(tokens[2])
			if err != nil {
				log.Fatal(err)
			}
			g.g.AddBoth(e1-1, e2-1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return g
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func canColor(g Graph, v int, c int, x []int) bool {
	return !g.g.Visit(v, func(n int, w int64) (skip bool) {
		if x[n] == c {
			return true
		}
		return
	})
}

func count(x []int) map[int]int {
	m := make(map[int]int)
	for _, c := range x {
		if c != 0 {
			m[c]++
		}
	}
	return m
}

func deepcopy(x []int) []int {
	xcopy := make([]int, len(x))
	for i, xi := range x {
		xcopy[i] = xi
	}
	return xcopy
}

func backtrack(g Graph, v int, x []int, k int) {
	// for c := 1; c <= g.V; c++ {
	for c := 1; c <= k+1; c++ {
		if canColor(g, v, c, x) {
			x[v] = c
			k = len(count(x))
			// fmt.Println(x)
			if v+1 < g.V {
				// fmt.Printf("v=%d c=%d k=%d min=%d\n", v, c, k, kMin)
				if k > kMin {
					return
				}
				backtrack(g, v+1, deepcopy(x), k)
			} else {
				pkMin := kMin
				kMin = min(kMin, k)
				exploredSolutions++
				if pkMin != kMin || exploredSolutions%100000 == 0 {
					fmt.Printf("k=%d min=%d exploredSolutions=%d\n", k, kMin, exploredSolutions)
					// time.Sleep(100 * time.Millisecond)
				}
				sol = x
				return
			}
		}
	}
	// return x
}

func runBacktrack(g Graph) {
	sol = make([]int, g.V)
	x := make([]int, g.V)

	start := time.Now()

	backtrack(g, 0, deepcopy(x), 0)

	elapsed := time.Since(start)

	fmt.Printf("sol=%v min=%d\n", sol, kMin)
	fmt.Printf("exploredSolutions=%d\n", exploredSolutions)
	fmt.Printf("backtrack took %s", elapsed)
}

func lowerLimit(g Graph, x []int, k int) int {
	lower := k
	coloredCount := 0
	for _, c := range x {
		if c != 0 {
			coloredCount++
		}
	}

	adjToAllColoredMark := make([]bool, g.V)
	adjToAllColored := list.New()
	inc := false
	for v := 0; v < g.V; v++ {
		if x[v] == 0 {
			extCount := 0
			inc = g.g.Visit(v, func(n int, w int64) (skip bool) {
				if x[n] != 0 {
					extCount++
					if extCount == coloredCount {
						adjToAllColoredMark[v] = true
						adjToAllColored.PushBack(v)
						return true
					}
				}
				return
			}) || inc
		}
	}

	if inc {
		lower++
	}

	for e := adjToAllColored.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		if g.g.Visit(v, func(n int, w int64) (skip bool) {
			if adjToAllColoredMark[n] == true {
				lower++
				return true
			}
			return
		}) {
			break
		}
	}
	return lower
}

func upperLimit(g Graph, x []int, k int) (int, []int) {
	nonColoredCount := 0
	for _, c := range x {
		if c == 0 {
			nonColoredCount++
		}
	}

	pq := make(PriorityQueue, nonColoredCount)
	i := 0
	for v, c := range x {
		if c == 0 {
			pq[i] = &Item{
				value:    v,
				priority: g.g.Degree(v),
				index:    v,
			}
			i++
		}
	}
	heap.Init(&pq)

	l := list.New()
	for pq.Len() > 0 {
		l.PushBack(heap.Pop(&pq).(*Item))
	}

	greedyX := deepcopy(x)
	greedyPlus(g, l, greedyX)
	greedyK := len(count(greedyX))
	return greedyK, greedyX
}

func bab(g Graph, v int, x []int, k int) {
	for c := 1; c <= k+1; c++ {
		// if finished == true {
		// 	return
		// }
		if canColor(g, v, c, x) {
			// fmt.Println(x)
			l := lowerLimit(g, x, k)
			if l >= kMin {
				branchedSolutions++
				continue
			}
			u, ux := upperLimit(g, x, k)
			// _, _ = upperLimit(g, x, k)
			// u, _ := upperLimit(g, x, k)
			limitsChecked++
			// fmt.Printf("lower=%d	upper=%d\n", l, u)
			if l == u {
				// fmt.Println(l, u, kMin, sol)
				// time.Sleep(time.Millisecond * 100)
				pkMin := kMin
				kMin = u
				sol = ux
				branchedSolutions++
				if pkMin != kMin || branchedSolutions%100000 == 0 {
					fmt.Printf("k=%d min=%d exploredSolutions=%d branchedSolutions=%d limitsChecked=%d\n", u, kMin, exploredSolutions, branchedSolutions, limitsChecked)
				}
				kMin = u
				sol = ux
				// finished = true
				// fmt.Printf("x=%d k=%d min=%d exploredSolutions=%d\n", x, k, kMin, exploredSolutions)
				// fmt.Println("ACABOOOOO")
				// fmt.Println(kMin, sol)
				return
			}
			// fmt.Printf("x=%d k=%d\n", x, k)
			x[v] = c
			k = len(count(x))
			if v+1 < g.V {
				if k > kMin {
					return
				}
				bab(g, v+1, deepcopy(x), k)
			} else {
				pkMin := kMin
				kMin = min(kMin, k)
				exploredSolutions++
				if pkMin != kMin || exploredSolutions%100000 == 0 {
					fmt.Printf("k=%d min=%d exploredSolutions=%d branchedSolutions=%d limitsChecked=%d\n", k, kMin, exploredSolutions, branchedSolutions, limitsChecked)
				}
				sol = x
				return
			}
		}
	}
}

func runBab(g Graph) {
	sol = make([]int, g.V)
	x := make([]int, g.V)

	start := time.Now()

	U = g.V
	L = 0
	bab(g, 0, deepcopy(x), 0)

	elapsed := time.Since(start)

	fmt.Printf("sol=%v min=%d\n", sol, kMin)
	fmt.Printf("exploredSolutions=%d\n", exploredSolutions)
	fmt.Printf("branchedSolutions=%d\n", branchedSolutions)
	fmt.Printf("branch and bound took %s", elapsed)
}

func greedy(g Graph, l *list.List, x []int) {
	front := l.Front()
	x[front.Value.(*Item).value] = 1
	for c := 1; c <= g.V; c++ {
		for e := front.Next(); e != nil; e = e.Next() {
			item := e.Value.(*Item)
			if x[item.value] == 0 && canColor(g, item.value, c, x) {
				x[item.value] = c
			}
		}
	}
}

func greedyPlus(g Graph, l *list.List, x []int) {
	for l.Len() > 0 {
		e := l.Front()
		item := l.Remove(e).(*Item)
		for c := 1; c <= g.V; c++ {
			if x[item.value] == 0 && canColor(g, item.value, c, x) {
				x[item.value] = c
				break
			}
		}
	}
}

func heuristic(g Graph) (x []int) {
	pq := make(PriorityQueue, g.V)
	for v := 0; v < g.g.Order(); v++ {
		pq[v] = &Item{
			value:    v,
			priority: g.g.Degree(v),
			index:    v,
		}
	}
	heap.Init(&pq)

	l := list.New()
	for pq.Len() > 0 {
		l.PushBack(heap.Pop(&pq).(*Item))
	}

	x = make([]int, g.V)
	greedy(g, l, x)

	return x
}

func heuristicPlus(g Graph) (x []int) {
	pq := make(PriorityQueue, g.V)
	for v := 0; v < g.g.Order(); v++ {
		pq[v] = &Item{
			value:    v,
			priority: g.g.Degree(v), /* + random(0, 10)*/
			index:    v,
		}
	}
	heap.Init(&pq)

	l := list.New()
	for pq.Len() > 0 {
		l.PushBack(heap.Pop(&pq).(*Item))
	}

	x = make([]int, g.V)
	greedyPlus(g, l, x)

	return x
}

func runHeuristic(g Graph) {
	start := time.Now()

	x := heuristic(g)
	k := len(count(x))

	elapsed := time.Since(start)

	fmt.Printf("sol=%v k=%d\n", x, k)
	fmt.Printf("greedy took %s", elapsed)
}

func runHeuristicPlus(g Graph) {
	start := time.Now()

	x := heuristicPlus(g)
	k := len(count(x))

	elapsed := time.Since(start)

	fmt.Printf("sol=%v k=%d\n", x, k)
	fmt.Printf("greedy+ took %s", elapsed)
}

func meta(g Graph, randomicCeil int) (x []int) {
	pq := make(PriorityQueue, g.V)
	for v := 0; v < g.g.Order(); v++ {
		pq[v] = &Item{
			value:    v,
			priority: g.g.Degree(v), /* + random(0, 10)*/
			index:    v,
		}
	}
	heap.Init(&pq)

	l := list.New()
	for pq.Len() > 0 {
		l.PushBack(heap.Pop(&pq).(*Item))
	}

	x = make([]int, g.V)
	for l.Len() > 0 {
		e := l.Front()
		for i := 0; i < random(0, min(randomicCeil, l.Len())); i++ {
			e = e.Next()
		}
		item := l.Remove(e).(*Item)
		for c := 1; c <= g.V; c++ {
			if x[item.value] == 0 && canColor(g, item.value, c, x) {
				x[item.value] = c
			}
		}
	}

	return x
}

func runMeta(g Graph) {
	start := time.Now()

	maxIt := int(1e6)
	concurrencyLevel := int(1e3)

	ch := make(chan *Result)

	var resultsWg sync.WaitGroup
	resultsWg.Add(1)

	go func(ch <-chan *Result, wg *sync.WaitGroup) {
		defer resultsWg.Done()
		it := 1
		for result := range ch {
			if result.k < kMin {
				kMin = result.k
				sol = result.x
				fmt.Printf("k=%d in %s (%d iterations)\n", result.k, result.elapsed, it)
			}
			if it%1000 == 0 {
				fmt.Printf("Executed %d iterations\n", it)
			}
			it++
		}
	}(ch, &resultsWg)

	for i := 0; i < maxIt; i += concurrencyLevel {
		var wg sync.WaitGroup
		wg.Add(concurrencyLevel)
		for t := 0; t < concurrencyLevel; t++ {
			go func(start time.Time, ch chan<- *Result, wg *sync.WaitGroup) {
				defer wg.Done()

				x := meta(g, 1)
				k := len(count(x))

				ch <- &Result{
					x:       x,
					k:       k,
					elapsed: time.Since(start),
				}
			}(start, ch, &wg)
		}
		wg.Wait()
	}

	close(ch)

	resultsWg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("sol=%v k=%d\n", sol, kMin)
	fmt.Printf("meta runned for %s", elapsed)
}

var (
	finished          = false
	kMin              = 100000000
	exploredSolutions = 0
	sol               []int
	U                 int
	L                 = 0
	limitsChecked     = 0
	branchedSolutions = 0
)

func main() {
	rand.Seed(time.Now().Unix())
	instance := os.Args[1]
	g := createGraphFromFile(fmt.Sprintf("instances/%s.col", instance))

	algo := os.Args[2]
	switch algo {
	case "backtrack":
		runBacktrack(g)
		break
	case "bab":
		runBab(g)
		break
	case "heuristic":
		runHeuristic(g)
		break
	case "heuristic+":
		runHeuristicPlus(g)
		break
	case "meta":
		runMeta(g)
		break
	}

	// test reference/value
	// x := 0
	// foo(x)
	// fmt.Println(x)

	// b := make([]int, 3)
	// // b := [g.V]int{0, 0, 0}
	// foo(b[0])
	// fmt.Println(b)
	// bCopy := deepcopy(b)
	// bar(bCopy)
	// fmt.Println(b)

	// set := make(map[int]bool)
	// i := 0
	// rec(set, i)
	// fmt.Println(set)
}

func rec(set map[int]bool, i int) {
	set[i] = true
	fmt.Println("1", set)
	if i == 3 {
		return
	}
	rec(set, i+1)
	fmt.Println("2", set)
}

// TESTS
func foo(x int) {
	x++
}

func bar(x []int) {
	for i := 0; i < len(x); i++ {
		x[i]++
	}
	fmt.Println(x)
}

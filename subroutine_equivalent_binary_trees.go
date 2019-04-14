package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch) // Needed if one wants to check if a channel is closed or not
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	
	ch <- t.Value
	walkHelper(t.Left, ch)
	walkHelper(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		v1, ok1 := <-ch1    // Check if a channel is closed (via close(ch))
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			break
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
	return true
}
	
func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))
}

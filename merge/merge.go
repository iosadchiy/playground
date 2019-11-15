package main

import "fmt"

// Merge two channels with values sorted in acsending order
// into a channel with unique values preseving the order. 
func merge(c1, c2 chan int) chan int {
	res := make(chan int)

	go func() {
		var x int
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		for {
			switch {
			case !ok1 && !ok2:
				close(res)
				return
			case !ok1:
				x = v2
			case !ok2:
				x = v1
			default:
				x = v1
				if v1 > v2 {
					x = v2
				}
			}
			res <- x
			if v1 == x {
				v1, ok1 = read(c1, x)
			}
			if v2 == x {
				v2, ok2 = read(c2, x)
			}
		}
	}()

	return res
}

func read(c chan int, skip int) (x int, ok bool) {
	for {
		x, ok = <-c
		if !ok || x != skip {
			return x, ok
		}
	}
}

func write(c chan int, vv []int) {
	for _, v := range vv {
		c <- v
	}
	close(c)
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c := merge(c1, c2)
	go write(c1, []int{1, 5, 9, 9, 10})
	go write(c2, []int{1, 3, 8, 9})

	for v := range c {
		fmt.Println(v)
	}
}

package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
}

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(^uint(0) >> 1) 
const MinInt = -MaxInt - 1 


func sameBstUtil(a[] int, b[] int, n int, i1 int, i2 int, min int, max int) int{
	var j int
	var k int
	
	for  j := i1; j < n; j++ {
		if a[j] > min && a[j] < max{
			break
		}
	}
	for k := i2; k < n; k++ {
		if b[k] > min && b[k] < max{
			break	
		}	
	}
	if j == n && k == n {
		return true
	}
	if (j == n) ^ (k == n) || a[j] != b[k]{
		return false
	}
	return sameBst(a, b, j+1, k+1, a[j], max) && sameBst(a, b, n, j+1, k+1, min, a[j])
}

func sameBst(a[] int, b[] int, n int) bool{
	return sameBstUtil(a, b, n, 0, 0, INT_MIN, INT_MAX)
}

func main(){
	a := make([]int, 7)
	a[0], a[1], a[2], a[3], a[4], a[5], a[6] = 8, 3, 6, 1, 4, 7, 10
	b := make([]int, 7)
	b[0], b[1], b[2], b[3], b[4], b[5], b[6] = 1, 6, 8, 4, 3, 10, 7
	n := len(a)
	fmt.Printf("same %s\n", sameBst(a, b, n))
}



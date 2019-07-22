package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func checkOrderedSubset(a []int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	} else if len(a) == 2 {
		if a[0] == a[1] {
			return -1
		}
		return max(a[0], a[1])
	}
	if a[0]+a[1] <= a[len(a)-1] {
		return -1
	}

}

func main() {

}

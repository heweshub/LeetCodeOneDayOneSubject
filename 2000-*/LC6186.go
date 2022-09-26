package main

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	type pair struct{ or, i int }
	ors := []pair{}
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ors = append(ors, pair{0, i})
		ors[0].or |= num
		k := 0
		for _, p := range ors[1:] {
			p.or |= num
			if ors[k].or == p.or {
				ors[k].i = p.i
			} else {
				k++
				ors[k] = p
			}
		}
		ors = ors[:k+1]
		// fmt.Println(ors)
		ans[i] = ors[0].i - i + 1
	}
	// fmt.Println(ans)
	return ans
}

func main() {
	smallestSubarrays([]int{1, 0, 2, 1, 3})
}

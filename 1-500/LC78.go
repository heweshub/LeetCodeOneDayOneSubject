package main

func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	// cur会逐渐增加,set会考虑选或者不选
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

// func main() {
// 	fmt.Println(subsets([]int{1, 2, 3}))
// }

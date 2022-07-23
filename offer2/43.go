package main

func countDigitOne(n int) (ans int) {
    for k, mulk := 0, 1; n >= mulk; k++ {
        ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk)
        mulk *= 10
    }
    return
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
func min(x,y int) int {
    if x < y {
        return x
    }
    return y
}

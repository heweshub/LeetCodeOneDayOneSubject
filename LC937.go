package main

import (
	"sort"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) (res []string) {
	resNum := []string{}
	resLett := [][]string{}
	for i := range logs {
		itemList := strings.Split(logs[i], " ")
		if _, err := strconv.Atoi(string(itemList[1][0])); err == nil {
			resNum = append(resNum, logs[i])
		} else {
			resLett = append(resLett, []string{itemList[0], logs[i][len(itemList[0])+1:]})
		}
	}
	sort.Slice(resLett, func(i, j int) bool {
		return resLett[i][0] < resLett[j][0]
	})
	sort.Slice(resLett, func(i, j int) bool {
		return resLett[i][1] < resLett[j][1]
	})
	for i := range resLett {
		res = append(res, resLett[i][0]+" "+resLett[i][1])
	}
	res = append(res, resNum...)
	// fmt.Println(res)
	return res
}

func main() {
	reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"})
}

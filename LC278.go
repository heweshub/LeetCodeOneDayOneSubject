package main

import "sort"

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}

package main

import (
	"fmt"
	"sort"
	"strings"
)

// for realisation "Interface" interface package sort
type sortString []rune

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortString) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortString) Len() int {
	return len(s)
}

func Sort(str sortString) string {
	sort.Sort(str)
	return string(str)
}

func checkUnique(arr []string) (res []string) {
	v := ""
	for _, s := range arr {
		if s == v {
			continue
		}
		v = s
		res = append(res, s)
	}
	return
}

func main() {
	m := make(map[string]*[]string)
	var arr = []string{"ПОЛКОВНИК", "ТЕРПЕЛИВОСТЬ", "клоповник", "ПРОСВЕТИТЕЛЬ"}

	fmt.Println("до сортировки:")
	for i, s := range arr {
		fmt.Print(s, " ")
		arr[i] = strings.ToLower(s)
	}

	sort.Strings(arr)
	arr = checkUnique(arr)

	fmt.Printf("\nпосле:\n")
	for _, s := range arr {
		fmt.Print(s, " ")
	}

	m = SetAnagrams(arr)

	fmt.Printf("\n\nАнаграммы:\n")
	for k, v := range m {
		fmt.Println(k, ":")
		for _, s := range *v {
			fmt.Print(s, " ")
		}
		fmt.Println()
	}
}

func SetAnagrams(arr []string) map[string]*[]string {
	a := make(map[string][]string)
	for _, s := range arr {
		ss := Sort(sortString(s))
		a[ss] = append(a[ss], s)
	}

	res := make(map[string]*[]string)
	for _, v := range a {
		val := v
		res[val[0]] = &val
	}

	return res
}
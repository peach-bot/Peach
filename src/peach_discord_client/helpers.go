package main

import "sort"

func sliceContains(s []string, searchterm string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func sliceRemove(s []string, str string) []string {
	if sliceContains(s, str) {
		i := sort.SearchStrings(s, str)
		s = append(s[:i], s[i+1:]...)
	}
	return s
}

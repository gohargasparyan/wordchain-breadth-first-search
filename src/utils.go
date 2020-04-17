package main

import "strings"

func Contains(element string, slice *[]string) bool {
	for _, current := range *slice {
		if strings.Compare(element, current) == 0 {
			return true
		}
	}
	return false
}

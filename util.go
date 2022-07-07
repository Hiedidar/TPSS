package main

import "sort"

func SortSplitInfoSlice(split []SplitInfo) {
	sort.SliceStable(split, func(i, j int) bool {
		if split[i].SplitType == "FLAT" && split[j].SplitType == "PERCENTAGE" {
			return true
		}
		if split[i].SplitType == "PERCENTAGE" && split[j].SplitType == "RATIO" {
			return true
		}

		if split[i].SplitType == "FLAT" && split[j].SplitType == "RATIO" {
			return true
		}

		if split[i].SplitType == "RATIO" && (split[j].SplitType != "FLAT" && split[j].SplitType != "PERCENTAGE" && split[j].SplitType != "RATIO") {
			return true
		}

		return false
	})
}

func GetNumRatio(info []SplitInfo) float32 {
	var count float32 = 0
	for _, i := range info {
		if i.SplitType == "RATIO" {
			count += i.SplitValue
		}
	}
	return count
}

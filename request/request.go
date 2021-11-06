package request

import (
	"sort"
	"strconv"
)

type Request interface {
	//String struct to json string
	String() string
}

//removeStringDuplicates 去除重复的item
func removeStringDuplicates(item []string) (ids []string) {
	if len(item) <= 0 {
		return ids
	}
	sort.Strings(item)
	for i, id := range item {
		if (i >= 1 && id == item[i-1]) || len(id) <= 0 {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

//removeIntDuplicatesToString 去除重复的item
func removeIntDuplicatesToString(item []int) (ids []string) {
	if len(item) <= 0 {
		return ids
	}
	sort.Ints(item)
	for i, id := range item {
		if (i >= 1 && id == item[i-1]) || id <= 0 {
			continue
		}
		ids = append(ids, strconv.Itoa(id))
	}
	return ids
}

//removeIntDuplicates 去除重复的item
func removeIntDuplicates(item []int) (ids []int) {
	if len(item) <= 0 {
		return ids
	}
	sort.Ints(item)
	for i, id := range item {
		if (i >= 1 && id == item[i-1]) || id <= 0 {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

import (
	"sort"
	"strconv"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

// removeStringDuplicates 去除重复的item
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

// removeIntDuplicatesToString 去除重复的item
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

// removeIntDuplicates 去除重复的item
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

// removeStatusDuplicates 去除重复的人员状态
func removeStatusDuplicates(status []employee.Status) (state []string) {
	temp := map[string]struct{}{}
	for _, s := range status {
		if _, ok := temp[string(s)]; !ok {
			temp[string(s)] = struct{}{}
			state = append(state, string(s))
		}
	}
	return state
}

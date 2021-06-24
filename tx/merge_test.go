package tx

func merge(list1 []int, list2 []int) []int {
	if len(list1) == 0 {
		return list2
	}

	if len(list2) == 0 {
		return list1
	}
	pos1 := 0
	pos2 := 0
	res := []int{}
	for pos1 < len(list1) && pos2 < len(list2) {
		if list1[pos1] <= list2[pos2] {
			res = append(res, list1[pos1])
			pos1++
		} else {
			res = append(res, list2[pos2])
			pos2++
		}
	}
	for pos1 < len(list1) {
		res = append(res, list1[pos1])
		pos1++
	}

	for pos2 < len(list2) {
		res = append(res, list2[pos2])
		pos2++
	}
	return res
}

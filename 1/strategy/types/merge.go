package types

type MergeSort struct{}

func (m *MergeSort) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return mergeSort(left, right)
}

func mergeSort(left, right []int) []int {
	result := make([]int, 0, len(left) + len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right){
		if left[i] < left[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, left[j:]...)
	return result
}

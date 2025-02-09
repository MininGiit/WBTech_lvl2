package types


type QickSort struct{}

func (q *QickSort) Sort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)-1]
    less, equal, greater := partition(arr, pivot)
	
   return append(append(q.Sort(less), equal...), q.Sort(greater)...)
}

func partition(arr []int, pivot int) (less, equal, greater []int){
	for _, elem := range arr {
		if elem < pivot {
			less = append(less, elem)
		} else if elem == pivot {
			less = append(equal, elem)
		} else {
			less = append(greater, elem)
		}
	}
	return
}


package types

type SortStrategy interface {
    Sort([]int) []int
}

type Sorter struct {
    strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
    return &Sorter{strategy: strategy}
}

func (s *Sorter) Sort(arr []int) []int {
    return s.strategy.Sort(arr)
}
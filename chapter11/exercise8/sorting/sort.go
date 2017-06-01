// Sorting defines the sorter interface and provides for sorting method
package sorting

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(s Sorter) {
	for pass := 1; pass < s.Len(); pass++ {
		for i := 0; i < s.Len()-pass; i++ {
			if s.Less(i+1, i) {
				s.Swap(i, i+1)
			}
		}
	}
}


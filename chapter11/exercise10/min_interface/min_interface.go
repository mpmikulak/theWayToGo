package min

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(x, y int) bool
}

func Min(m Miner) interface{} {
	min := m.ElemIx(0)
	for i := 1; i < m.Len(); i++ {
		if m.Less(i, i-1) {
			min = m.ElemIx(i)
		}
	}
	return min
}

// IntArray definition and methods
type IntArray []int

func (i IntArray) Len() int                  { return len(i) }
func (i IntArray) ElemIx(ix int) interface{} { return i[ix] }
func (i IntArray) Less(x, y int) bool        { return i[x] < i[y] }

// StringArray definition and methods
type StringArray []string

func (i StringArray) Len() int                  { return len(i) }
func (i StringArray) ElemIx(ix int) interface{} { return i[ix] }
func (i StringArray) Less(x, y int) bool {
	return i[x] < i[y]
}

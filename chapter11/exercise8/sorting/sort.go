package sorting

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(s Sorter) {
	for pass := 0; pass < s.Len(); pass++ {
		for i := 0; i < s.Len()-pass; i++ {
			if s.Less(i+1, i) {
				s.Swap(i+1, i)
			}
		}
	}
}

// Person definition
type Person struct {
	firstName string
	lastName  string
}

// Persons definition and methods
type Persons []Person

func (p Persons) Len() int      { return len(p) }
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Persons) Less(i, j int) bool {
	f := p[i].firstName + " " + p[i].lastName
	s := p[j].firstName + " " + p[j].lastName
	return f < s
}

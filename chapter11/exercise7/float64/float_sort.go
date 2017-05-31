package float64

import (
	"math/rand"
	"strconv"
	"time"
)

// Sorter interface definition to provide sorting capabilities and methods
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

// Float64Array type definition and methods to implement the Sorter interface
type Float64Array []float64

func (a Float64Array) Len() int           { return len(a) }
func (a Float64Array) Less(i, j int) bool { return a[i] < a[j] }
func (a Float64Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// New returns a pointer to an initialized Float64Array that is 25 long
func New() Float64Array {
	return make(Float64Array, 25)
}

// List pretty-prints the array
func (a Float64Array) List() string {
	str := ""
	for ix, v := range a {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.FormatFloat(v, 'f', -1, 64) + "]\n"
	}
	return str
}

// String wraps the List method for printing
func (a Float64Array) String() string {
	return a.List()
}

func (a Float64Array) Fill() {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < 10; i++ {
		f := r.Intn(len(a))
		a[f] = r.NormFloat64()
	}
}

package main

func main() {

}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, len(data)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

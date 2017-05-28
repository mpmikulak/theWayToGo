package main

import "fmt"

type Base struct{}

func (Base) Magic() { fmt.Print("base magic ") }

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() { fmt.Println("Voodoo magic") }

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic() // I got this wrong, when MoreMagic() is called, the type is Base
}

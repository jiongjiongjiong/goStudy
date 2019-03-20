package ch4

import "testing"

const  (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T)  {
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,3,2,4}
	//c:=[...]int{1,2,3,4,5}
	d:=[...]int{1,2,3,4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}


func TestBitClear(t *testing.T)  {
	a:=7//0111
	//a:=1
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable==Readable, a&Writeable==Writeable, a&Executable==Executable)
}
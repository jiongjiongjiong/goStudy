package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednessday
)

const  (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T)  {
	t.Log(Monday,Tuesday,Wednessday)
}

func TestContantTry1(t *testing.T)  {
	//a:=7//0111
	a:=1
	t.Log(a&Readable==Readable, a&Writeable==Writeable, a&Executable==Executable)
}
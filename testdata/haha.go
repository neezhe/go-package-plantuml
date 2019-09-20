package testdata

import "strings"

type Foo struct {
	field       int
	PublicField bool
	*strings.Builder
}
type MyInterface interface {
	GetValue() bool
}

func (f *Foo) GetValue() bool {
	return false
}

func (f *Foo) Lizhe(aa bool) bool {
	return true
}

func (f *Foo) Caiying(bb bool) bool {
	return true
}
type MMM struct {
	a int
	bb MyInterface
}

var zzz MMM
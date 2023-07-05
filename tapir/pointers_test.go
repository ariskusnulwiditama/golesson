package tapir

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	p0 := new(int) //p0 points toa zero int value
	fmt.Println(p0) //(a hex address string)
	fmt.Println(*p0) //0

	//x is a copy of the value at the address store in p0
	x := *p0
	// both tahe the address of x
	// x, *p1 and p2 represent the same value
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) //true
	fmt.Println(p0 == p1) //false
	p3 := &*p0 // <=> p3 := &(*p0) <=> p3 := p0
	// now p3 and p0 store the same address
	fmt.Println(p0 == p3) //true
	*p0, *p1 = 123, 789
	fmt.Println(*p0, *p1)
}

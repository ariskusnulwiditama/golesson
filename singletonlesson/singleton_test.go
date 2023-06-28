package singletonlesson

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myStruct struct {
	data string
}

var instance *myStruct
var once sync.Once

func GetInstance() *myStruct {
	once.Do(func() {
		instance = &myStruct{
			data: "this is singleton example",
		}
	})
	return instance
}

func (m *myStruct) GetData() string {
	return m.data
}

func TestSingleton(t *testing.T) {
	obj := GetInstance()
	obj2 := GetInstance()

	if obj != obj2 {
		t.Errorf("expected same instance, but got different instances")
	}	
}

func TestIsSame(t *testing.T) {
	obj1 := GetInstance()
	obj2 := GetInstance()
	assert.Equal(t, obj1, obj2, "is equals")
}

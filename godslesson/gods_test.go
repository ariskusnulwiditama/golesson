package godslesson

import (
	"testing"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/sirupsen/logrus"
)

func TestGods(t *testing.T) {
	list := arraylist.New()
	list.Add("nanas")
	list.Add("apel")
	list.Add("anggur")
	list.Add("kiwi")
	fd,_:= list.Get(0)
	logrus.Info(fd)
	list.Remove(2)
}

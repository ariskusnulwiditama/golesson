package jsonlesson

import (
	"fmt"
	"testing"
)

func TestLogJson(t *testing.T) {
	logJson("testing")
	logJson(12345)
	logJson(true)
}

func TestEncode(t *testing.T) {
	res, err := encodeJson()
	if err != nil {
		t.Error("error encode")
	}

	fmt.Println("result json: ", string(res))
}

func TestDecode(t *testing.T) {
	decodeJson()
}
package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("我爱你", "爱")
	want := []string{"我", "你"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%v got:%v", want, got)
	}
	t.Run("", func(t *testing.T) {

	})
	a := map[int]string{}
	for i, j := range a {
		fmt.Println(i, j)
	}
}

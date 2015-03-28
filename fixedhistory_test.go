package fixedhistory

import "testing"

func TestArray(t *testing.T) {
	x := NewHistory(10)
	x.Push(1)
	x.Push(2)
	x.Push(3)
	x.Push(4)
	if !x.Contains(4) {
		t.Errorf("x [%+v] doesn't contain 4", x)
	}
	if x.Contains(5) {
		t.Errorf("x [%+v] shouldn't contain 5", x)
	}
	x.Remove(4)
	if x.Contains(4) {
		t.Errorf("x [%+v] shouldn't contain 4", x)
	}
	// cleanup even numbers
	x.Cleanup(func(v interface{}) bool {
		switch t := v.(type) {
		case int:
			return t%2 == 0
		}
		return false
	})
	if x.Contains(2) {
		t.Errorf("x [%+v] shouldn't contain 2", x)
	}
	// hacking eval'ed value
	x.ValueMap = func(i interface{}) interface{} {
		return 8
	}
	if !x.Contains(8) {
		t.Errorf("x [%+v] should contain 8", x)
	}
}

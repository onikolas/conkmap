package conkmap

import "testing"

//import "fmt"

func TestSerialAccess(t *testing.T) {

	m := New()

	var pos = map[interface{}]interface{}{

		1:     2,
		123:   "asdf",
		true:  false,
		nil:   true,
		false: nil,
	}

	for k, v := range pos {
		m.Set(k, v)
		if m.Get(k) != v {
			t.Error("Failed ", t, v)
		}
	}
}

//checks for write violations
func TestConcurrentAccess(t *testing.T) {

	m := New()

	// writes
	for i := 0; i < 1111; i++ {
		go func() {
			m.Set(i, i*10)
		}()
	}

	//reads
	for i := 0; i < 1111; i++ {
		go func() {
			m.Get(i)
		}()
	}

	// writes and reads
	for i := 0; i < 1024; i++ {
		go func() {
			m.Set(i, m.Get(i))
		}()
	}

}

func TestNotThere(t *testing.T) {
	m := New()
	a := m.Get(99)
	if a != nil {
		t.Error("has val")
	}
}


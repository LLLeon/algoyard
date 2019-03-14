package bitmap

import "testing"

func TestBitMap(t *testing.T) {
	bm := New(80)

	for i := 0; i <= 100; i += 10 {
		bm.Set(uint(i))
	}

	for i := 0; i <= 100; i += 10 {
		if exists := bm.Get(uint(i)); !exists {
			t.Logf("there is no value [%d] in the bitmap\n", i)
		}
	}

	t.Log("success")
}

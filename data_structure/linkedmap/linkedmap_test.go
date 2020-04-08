/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-08
+************************************************************************/

package linkedmap

import "testing"

func TestLinkedMap_Walk(t *testing.T) {
	lm := NewLinkedMap()

	lm.Set("a", 1)
	lm.Set("b", 2)
	lm.Set("c", 3)

	lm.Walk(lm.WalkCallback)
}

/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-03-22
+************************************************************************/

package linkedmap

import (
	"container/list"
	"fmt"
)

// LinkedMap implements an ordered map.
type LinkedMap struct {
	dataMap map[string]interface{}

	// store key for map
	linkedList *list.List
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		dataMap:    make(map[string]interface{}),
		linkedList: list.New(),
	}
}

// Exists checks for the presence of k in the map.
func (lm *LinkedMap) Exists(k string) bool {
	if lm == nil {
		return false
	}

	if _, ok := lm.dataMap[k]; !ok {
		return false
	}

	return true
}

// Set stores a key-value pair.
func (lm *LinkedMap) Set(k string, v interface{}) bool {
	if lm == nil {
		return false
	}

	if lm.Exists(k) {
		lm.dataMap[k] = v
		return true
	}

	lm.linkedList.PushBack(k)
	lm.dataMap[k] = v

	return true
}

// Get gets a key-value pair.
func (lm *LinkedMap) Get(k string) (interface{}, bool) {
	if lm == nil {
		return nil, false
	}

	v, ok := lm.dataMap[k]

	return v, ok
}

// Delete deletes a key-value pair.
func (lm *LinkedMap) Delete(k string) {
	if lm == nil {
		return
	}

	if _, ok := lm.dataMap[k]; !ok {
		return
	}

	delete(lm.dataMap, k)
	lm.linkedList.Remove(&list.Element{Value: k})
}

// Walk handles all key-value pairs in the map based on the function passed in.
func (lm *LinkedMap) Walk(cb func(k string)) {
	if lm == nil {
		return
	}

	for k := lm.linkedList.Front(); k != nil; k = k.Next() {
		cb(k.Value.(string))
	}
}

func (lm *LinkedMap) WalkCallback(k string) {
	v, _ := lm.Get(k)
	fmt.Printf("key: %s, value: %v\n", k, v)
}

package iterators

import "reflect"

type Iterator struct {
	data interface{}
}

func NewIterator(v interface{}) Iterator {
	return Iterator{v}
}

func (i *Iterator) MapInt(closure func(int) int) []int {
	var slc []int
	v := i.data
	refVal := reflect.ValueOf(v)
	switch refVal.Kind() {
	case reflect.Slice:
		for _, val := range v.([]int) {
			slc = append(slc, closure(val))
		}
	case reflect.Int:
		slc = append(slc, closure(v.(int)))
	}
	return slc
}

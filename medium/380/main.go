package main

import (
	"log"
	"math/rand/v2"
	"reflect"
)

type valuesDetails struct {
	count int
	idx   int
}

type RandomizedSet struct {
	data   []int
	size   int
	values map[int]*valuesDetails
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:   make([]int, 0),
		values: make(map[int]*valuesDetails),
		size:   0,
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.values[val]
	if !ok {
		rs.values[val] = &valuesDetails{
			count: 1,
			idx:   rs.size,
		}
		rs.size++
		rs.data = append(rs.data, val)
	} else {
		rs.values[val].count++
	}
	return !ok
}

func (rs *RandomizedSet) Remove(val int) bool {
	v, ok := rs.values[val]
	if ok {
		if v.count-1 == 0 {
			last := rs.data[rs.size-1]
			rs.data[v.idx] = last
			rs.values[last] = &valuesDetails{
				count: v.count,
				idx:   v.idx,
			}
			rs.data = rs.data[:rs.size-1]
			delete(rs.values, val)
		} else {
			rs.values[val].count--
		}
		rs.size--
	}
	return ok
}

func (rs *RandomizedSet) GetRandom() int {
	if rs.size == 0 {
		return -1
	}
	return rs.data[rand.IntN(rs.size)]
}

func main() {
	tests := []struct {
		operations []string
		values     []int
		result     []any
	}{
		//{
		//	operations: []string{"Insert", "Insert", "GetRandom", "GetRandom", "Insert", "Remove", "GetRandom", "GetRandom", "Insert", "Remove"},
		//	values:     []int{3, 3, 0, 0, 1, 3, 0, 0, 0, 0},
		//	result:     []any{true, false, 3, 3, true, true, 1, 1, true, true},
		//},
		//{
		//	operations: []string{"Insert", "Insert", "Insert", "Insert", "Insert", "Remove", "Insert", "GetRandom", "Insert", "Remove"},
		//	values:     []int{0, 2, 1, 1, 1, 0, 0, -1, 1, 2},
		//	result:     []any{true, true, true, false, false, true, true, 2, false, true},
		//},
		//{
		//	operations: []string{"Insert", "Insert", "Remove", "Insert", "Remove", "GetRandom"},
		//	values:     []int{0, 1, 0, 2, 1, -1},
		//	result:     []any{true, true, true, true, true, 2},
		//},
		{
			operations: []string{"Insert", "Insert", "Insert", "GetRandom", "Remove", "GetRandom"},
			values:     []int{1, 1, 2, -1, 1, -1},
			result:     []any{true, false, true, 2, true, 1},
		},
	}
	for i, t := range tests {
		randomizedSet := Constructor()
		value := reflect.ValueOf(&randomizedSet)
		for j, op := range t.operations {
			var v []reflect.Value
			if op == "GetRandom" {
				v = value.MethodByName(op).Call([]reflect.Value{})
			} else {
				v = value.MethodByName(op).Call([]reflect.Value{reflect.ValueOf(t.values[j])})
			}
			res := v[0].Interface()
			if !reflect.DeepEqual(res, t.result[j]) {
				log.Printf("[Test %v] Operation '%v'; Expected %v, got %v", i+1, op, t.result[j], res)
			} else {
				log.Printf("[Test %v] Operation '%v', test passed!", i+1, op)
			}

		}
	}
}

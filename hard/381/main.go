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

type RandomizedCollection struct {
	data   []int
	size   int
	values map[int]*valuesDetails
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		data:   make([]int, 0),
		values: make(map[int]*valuesDetails),
		size:   0,
	}
}

func (rs *RandomizedCollection) Insert(val int) bool {
	_, ok := rs.values[val]
	//log.Printf("Insert=%d, ok=%v", val, ok)
	//rs.printValues()
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

func (rs *RandomizedCollection) Remove(val int) bool {
	v, ok := rs.values[val]
	//log.Printf("Remove=%d, v=%+v, ok=%v", val, v, ok)
	//rs.printValues()
	if ok {
		if v.count-1 == 0 {
			rs.data[v.idx] = rs.data[rs.size-1]
			rs.values[rs.data[v.idx]].idx = v.idx
			rs.data = rs.data[:rs.size-1]
			delete(rs.values, val)
			rs.size--
		} else {
			rs.values[val].count--
		}
	}
	return ok
}

func (rs *RandomizedCollection) GetRandom() int {
	if rs.size == 0 {
		return -1
	}
	return rs.data[rand.IntN(rs.size)]
}

func (rs *RandomizedCollection) printValues() {
	for k, v := range rs.values {
		log.Printf("\tKey=%d, Value=%+v", k, v)
	}
}

func main() {
	tests := []struct {
		operations []string
		values     []int
		result     []any
	}{
		//{
		//	operations: []string{"Insert", "Insert", "Insert", "GetRandom", "Remove", "GetRandom"},
		//	values:     []int{1, 1, 2, -1, 1, -1},
		//	result:     []any{true, false, true, 2, true, 1},
		//},
		{
			operations: []string{"Insert", "Insert", "Remove", "Insert", "Remove", "GetRandom"},
			values:     []int{0, 1, 0, 2, 1, -1},
			result:     []any{true, true, true, true, true, 2},
		},
		//{
		//	operations: []string{"Insert", "Insert", "Insert", "Insert", "Insert", "Remove", "Remove", "Remove", "Insert", "Remove", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom", "GetRandom"},
		//	values:     []int{1, 1, 2, 2, 2, 1, 1, 2, 1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		//	result:     []any{true, false, true, false, false, true, true, true, true, true, 1, 2, 2, 2, 1, 1, 2, 1, 2, 1},
		//},
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
				if op != "GetRandom" {
					log.Printf("Data: %v", randomizedSet.data)
					log.Printf("Values: %v", randomizedSet.values)
					return
				}
			} else {
				//log.Printf("[Test %v] Operation '%v', test passed!", i+1, op)
			}
		}
	}
}

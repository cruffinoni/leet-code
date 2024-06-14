package main

import (
	"log"
	"reflect"
)

type MajorityChecker struct {
	arr []int
}

func Constructor(arr []int) MajorityChecker {
	return MajorityChecker{
		arr: arr,
	}
}

func (m *MajorityChecker) Query(left int, right int, threshold int) int {
	subArr := m.arr[left : right+1]
	count := make(map[int]int, right+1-left)
	for i := 0; i < right+1-left; i++ {
		count[subArr[i]]++
		if count[subArr[i]] >= threshold {
			return subArr[i]
		}
	}
	return -1
}

func main() {
	tests := []struct {
		operations []string
		values     []any
		result     []any
	}{
		{
			operations: []string{"Constructor", "Query", "Query", "Query"},
			values:     []any{[]int{1, 1, 2, 2, 1, 1}, []int{0, 5, 4}, []int{0, 3, 3}, []int{2, 3, 2}},
			result:     []any{nil, 1, -1, 1},
		},
	}
	for i, t := range tests {
		constr := Constructor(t.values[0].([]int))
		original := reflect.ValueOf(&constr)
		for j, op := range t.operations {
			switch op {
			case "Query":
				s := t.values[j].([]int)
				res := original.MethodByName(op).Call([]reflect.Value{reflect.ValueOf(s[0]), reflect.ValueOf(s[1]), reflect.ValueOf(s[2])})
				if !reflect.DeepEqual(res, t.result[j]) {
					log.Printf("[Test %v] Operation '%v'; Expected %v, got %v", i+1, op, t.result[j], res[0].Int())
					return
				}
			}
		}
	}
}

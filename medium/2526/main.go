package main

import "log"

type DataStream struct {
	k        int
	value    int
	queue    []int
	queueLen int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		k:        k,
		value:    value,
		queue:    make([]int, 0),
		queueLen: 0,
	}
}

func (d *DataStream) Consec(num int) bool {
	if num == d.value {
		d.queue = append(d.queue, num)
		d.queueLen++
		return d.queueLen == d.k
	}
	d.queue = make([]int, 0)
	d.queueLen = 0
	return false
}

func main() {
	constr := Constructor(4, 3)
	log.Printf("Result: %v", constr.Consec(4))
	log.Printf("Result: %v", constr.Consec(4))
	log.Printf("Result: %v", constr.Consec(4))
	log.Printf("Result: %v", constr.Consec(3))
}

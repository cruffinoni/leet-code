package main

type data struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]data
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]data),
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.m[key] = append(t.m[key], data{value: value, timestamp: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	d, ok := t.m[key]
	if !ok {
		return ""
	}
	if d[0].timestamp > timestamp {
		return ""
	}
	l := len(d)
	lo, hi := 0, l-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if d[mid].timestamp == timestamp {
			return d[mid].value
		} else if d[mid].timestamp < timestamp {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if hi < 0 {
		return ""
	}
	return d[hi%l].value
}

func main() {
	constr := Constructor()
	constr.Set("foo", "bar", 10)
	constr.Set("foo", "bar2", 20)
	constr.Set("foo", "bar3<<", 15)
	//res := constr.Get("foo", 5)
	//if !reflect.DeepEqual(res, "bar") {
	//	log.Fatalf("Test failed for %v, expected %s, got %s", "foo", "bar", res)
	//}
	//log.Printf("Test passed for %v", "foo")
	//res = constr.Get("foo", 3)
	//if !reflect.DeepEqual(res, "bar") {
	//	log.Fatalf(">Test failed for %v, expected %s, got %s", "foo", "bar", res)
	//}
	//constr.Set("foo", "bar2", 4)
	//res = constr.Get("foo", 4)
	//if !reflect.DeepEqual(res, "bar2") {
	//	log.Fatalf("Test failed for %v, expected %s, got %s", "foo", "bar2", res)
	//}
	//res = constr.Get("foo", 5)
	//if !reflect.DeepEqual(res, "bar2") {
	//	log.Fatalf(">Test failed for %v, expected %s, got %s", "foo", "bar2", res)
	//}
}

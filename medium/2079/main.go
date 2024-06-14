package main

func wateringPlants(plants []int, capacity int) int {
	l := len(plants)
	if capacity == 0 || l == 0 {
		return 0
	}
	steps := 0
	c := capacity
	for i := 0; i < l; i++ {
		if c-plants[i] < 0 {
			steps += i * 2
			//log.Printf("+%d", i)
			c = capacity
		}
		c -= plants[i]
		steps++
		//log.Printf("+1")
		//log.Printf("c=%d & p=%d > %d", c, plants[i], steps)
	}
	return steps
}

type test struct {
	plants   []int
	capacity int
	expected int
}

func main() {
	tests := []test{
		//{plants: []int{2, 2, 3, 3}, capacity: 5, expected: 14},
		//{plants: []int{1, 1, 1, 4, 2, 3}, capacity: 4, expected: 30},
		//{plants: []int{7, 7, 7, 7, 7, 7, 7}, capacity: 8, expected: 49},
		{plants: []int{3, 2, 4, 2, 1}, capacity: 6, expected: 17},
	}
	for _, t := range tests {
		result := wateringPlants(t.plants, t.capacity)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}

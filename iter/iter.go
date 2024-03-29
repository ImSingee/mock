package iter

// List mimic the python range function.
// List(10) -> generate a sequence from 0 to 9.
// List(3, 10) -> generate a sequence from 3 to 9.
// List(3, 10, 2) -> generate a sequence from 3 to 9 by incrementing value by 2.
func List(in ...int) []int {
	var start, end int
	step := 1
	switch len(in) {
	case 0:
		return []int{}
	case 1:
		end = in[0]
	case 2:
		start = in[0]
		end = in[1]
	case 3:
		start = in[0]
		end = in[1]
		step = in[2]
	}

	capacity := (end-start)/step + 1
	result := make([]int, 0, capacity)

	for i := start; i < end; i += step {
		result = append(result, i)
	}

	return result
}

// N mimic the python range function.
// N(10) -> generate a sequence from 0 to 9.
// N(3, 10) -> generate a sequence from 3 to 9.
// N(3, 10, 2) -> generate a sequence from 3 to 9 by incrementing value by 2.
func N(in ...int) <-chan int {
	ch := make(chan int)
	var start, end int
	step := 1
	switch len(in) {
	case 0:
		close(ch)
		return ch
	case 1:
		if in[0] > 0 {
			end = in[0]
		}
	case 2:
		if in[0] > 0 {
			start = in[0]
		}
		if in[0] > 0 {
			end = in[1]
		}
	case 3:
		if in[0] > 0 {
			start = in[0]
		}
		if in[1] > 0 {
			end = in[1]
		}
		if in[2] > 0 {
			step = in[2]
		}

	}
	go func(ch1 chan<- int) {
		for i := start; i < end; i += step {
			ch1 <- i
		}
		close(ch1)
	}(ch)
	return ch
}

// L generate a range of letters, such as: a-z or A-Z.
func L(start, end byte) <-chan byte {
	ch := make(chan byte)
	go func(ch1 chan<- byte) {
		for i := start; i <= end; i++ {
			ch1 <- i
		}
		close(ch1)
	}(ch)
	return ch
}

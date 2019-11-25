package util

func PermuteInt(options []int) <-chan []int {
	c := make(chan []int)

	go func() {
		defer close(c)
		permuteHeapInt(options, len(options), c)
	}()

	return c
}

type intArray []int

func (arr intArray) swap(a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func permuteHeapInt(options []int, size int, c chan []int) {
	if size == 1 {
		c <- options
		return
	}
	for i := 0; i < size-1; i++ {
		permuteHeapInt(options, size-1, c)
		if i%2 == 0 {
			intArray(options).swap(size-1, i)
		} else {
			intArray(options).swap(size-1, 0)
		}
	}
	permuteHeapInt(options, size-1, c)
}

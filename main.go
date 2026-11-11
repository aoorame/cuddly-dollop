package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements возвращает слайс случайных чисел в количестве size в диапазоне от 1 до 1000 включительно. При size <= 0 возвращается пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1000) + 1
	}
	return result
}

// maximum возвращает максимальное число в слайсе. При пустом слайсе возвращается 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks возвращает максимальное значение переданного слайса. Ищет параллельно, разбив данные на CHUNKS частей, по горутине на часть. Пустой слайс дает 0.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	for i := range maxValues {
		maxValues[i] = data[0]
	}

	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		subSlice := data[start:end]

		wg.Add(1)
		go func(chunkIdx int, s []int) {
			defer wg.Done()

			if len(s) == 0 {
				return
			}

			maxValues[chunkIdx] = maximum(s)
		}(i, subSlice)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	nums := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(nums)
	duration := time.Since(start)
	elapsed := duration.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(nums)
	duration = time.Since(start)
	elapsed = duration.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n\n", max, elapsed)
}

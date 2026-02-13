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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	m := make([]int, 0)
	if size <= 0 {
		return m
	}
	for i := 0; i < size; i++ {
		m = append(m, rand.Intn(SIZE))
	}
	return m
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	max := data[0]
	for _, d := range data {
		if max >= d {
			continue
		}
		max = d
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data) < 8 {
		return maximum(data)
	}

	var wg sync.WaitGroup
	sp := splitSlice(data)
	results := make([]int, CHUNKS)
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		go func() {
			defer wg.Done()
			results[i] = maximum(sp[i])
		}()
	}
	wg.Wait()

	return maximum(results)
}

func splitSlice(s []int) [][]int {
	partSize := len(s) / CHUNKS
	remaining := len(s) % CHUNKS

	result := make([][]int, 0, CHUNKS)
	start := 0
	for i := 0; i < CHUNKS; i++ {
		end := start + partSize
		if i < remaining {
			end++
		}
		result = append(result, s[start:end])
		start = end
	}
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел/n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	t1 := time.Now()
	max := maximum(data)
	elapsed := time.Since(t1)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	t1 = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(t1)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

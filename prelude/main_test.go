package main

import (
    "testing"
	"math/rand"
	"runtime"
	"gonum.org/v1/gonum/mat"
)

var Input = make([]int, 1000000)
var Vec = mat.NewVecDense(1000000, nil)

func init() {
	for i := range Input {
		Input[i] = rand.Intn(100)
	}
    data := make([]float64, len(Input))
    for i, v := range Input {
        data[i] = float64(v)
    }
    
    Vec = mat.NewVecDense(len(data), data)
}

func BenchmarkSequentialSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        c := 0
        for j := 0; j < len(Input); j++ {
            c += Input[j]
        }
        _ = c
    }
}

func BenchmarkUnrolledSum2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        c := 0
        d := 0
        for j := 0; j < len(Input); j += 2 {
            c += Input[j]
            d += Input[j+1]
        }
        result := c + d
        _ = result 
    }
}

func QuadScalarSlice(input []int) int {
    sumA, sumB, sumC, sumD := 0, 0, 0, 0
    
    // Process 4 elements at a time
    for i := 0; i < len(input); i += 4 {
        sumA += input[i]
        sumB += input[i+1] 
        sumC += input[i+2]
        sumD += input[i+3]
    }
    
    return sumA + sumB + sumC + sumD
}

//go.noinline
func BenchmarkUnrolledSum4(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		result := QuadScalarSlice(Input)
		_ = result
	}
}

func simdHintSum(data []int) int {
    sum1 := 0
    sum2 := 0
    sum3 := 0
    sum4 := 0
    sum5 := 0
    sum6 := 0
    sum7 := 0
    sum8 := 0

    // Process in chunks that hint at vectorization
    for i := 0; i < len(data); i += 8 {
        sum1 += data[i]
        sum2 += data[i+1]
        sum3 += data[i+2]
        sum4 += data[i+3]
        sum5 += data[i+4]
        sum6 += data[i+5]
        sum7 += data[i+6]
        sum8 += data[i+7]
    }

    return sum1 + sum2 + sum3 + sum4 + sum5 + sum6 + sum7 + sum8
}

func BenchmarkSIMDHint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := simdHintSum(Input)
		_ = result
	}
}

func BenchmarkGonumSum(b *testing.B) {

    for i := 0; i < b.N; i++ {
        sum := mat.Sum(Vec)
        _ = sum
    }
}

func parallelSum(input []int, numWorkers int) int {
	chunkSize := len(input) / numWorkers
	results := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(input)
		}

		go func(slice []int) {
			sum := 0
			for _, val := range slice {
				sum += val
			}
			results <- sum
		}(input[start:end])
	}

	total := 0
	for i := 0; i < numWorkers; i++ {
		total += <-results
	}
	return total
}

func BenchmarkParallelSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := parallelSum(Input, runtime.NumCPU())
		_ = result
	}
}


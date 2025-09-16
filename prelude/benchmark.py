import random
import timeit

import cython_sum
import numpy as np

# Initialize data
COUNT = 500000
INPUT = [random.randint(0, 99) for _ in range(COUNT)]
INPUT_ARRAY = np.array(INPUT, dtype=np.uint32)  # Convert to numpy array with correct type

def sequential_sum():
    c = 0
    for item in INPUT:
        c += item
    return c

def unrolled_sum_2():
    c = 0
    d = 0
    for j in range(0, len(INPUT), 2):
        c += INPUT[j]
        d += INPUT[j+1]
    return c + d

def cython_sum4():
    return cython_sum.CythonSum(COUNT, INPUT_ARRAY)

def benchmark_function(func, name, iterations=1000):
    # Time the function
    ghz = 2.3
    time_taken = timeit.timeit(func, number=iterations)
    time_per_op = (time_taken / iterations) * 1_000_000_000  # Convert to nanoseconds
    adds_per_cycle = COUNT / (time_per_op * ghz)
    
    print(f"{name:20} {iterations:8} iterations, {time_per_op:10.0f} ns/op")
    print(f"\t\t {adds_per_cycle:0.5f} adds/cycle")

if __name__ == "__main__":
    print("Function             Iterations    Time per op")
    print(COUNT)
    print("-" * 50)
    
    benchmark_function(sequential_sum, "Sequential", 1000)
    benchmark_function(unrolled_sum_2, "Unrolled x2", 1000)
    benchmark_function(cython_sum4, "Cython x4", 1000)
    
    # Also test with NumPy for comparison
    np_array = np.array(INPUT)
    
    def numpy_sum():
        return np.sum(np_array)

    def numpy_sum_split4():
        quarter = len(np_array) // 4
        return (np.sum(np_array[:quarter]) +
                np.sum(np_array[quarter:2*quarter]) +
                np.sum(np_array[2*quarter:3*quarter]) +
                np.sum(np_array[3*quarter:]))

    benchmark_function(numpy_sum, "NumPy", 1000)
    benchmark_function(numpy_sum_split4, "NumPy Split x4", 1000)

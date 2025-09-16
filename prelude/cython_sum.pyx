
import array
from libc.stdint cimport uint32_t
cdef extern from "immintrin.h":
    ctypedef struct __m256i:
        pass
    __m256i _mm256_setzero_si256()
    __m256i _mm256_add_epi32(__m256i a, __m256i b)
    __m256i _mm256_loadu_si256(__m256i *mem_addr)
    __m256i _mm256_hadd_epi32(__m256i a, __m256i b)
    __m256i _mm256_permute2x128_si256(__m256i a, __m256i imm, int control)
    int _mm256_cvtsi256_si32(__m256i a)

def CythonSum(unsigned int TotalCount, unsigned int[::1] InputArray):
    cdef unsigned int Count
    cdef unsigned int *Input
    cdef __m256i SumA, SumB, SumC, SumD
    cdef __m256i SumAB, SumCD
    cdef __m256i Sum, SumS

    Input = &InputArray[0]

    SumA = _mm256_setzero_si256()
    SumB = _mm256_setzero_si256()
    SumC = _mm256_setzero_si256()
    SumD = _mm256_setzero_si256()

    Count = TotalCount >> 5
    while Count != 0:
        SumA = _mm256_add_epi32(SumA, _mm256_loadu_si256(<__m256i *>&Input[0]))
        SumB = _mm256_add_epi32(SumB, _mm256_loadu_si256(<__m256i *>&Input[8]))
        SumC = _mm256_add_epi32(SumC, _mm256_loadu_si256(<__m256i *>&Input[16]))
        SumD = _mm256_add_epi32(SumD, _mm256_loadu_si256(<__m256i *>&Input[24]))

        Input += 32
        Count -= 1
        
    SumAB = _mm256_add_epi32(SumA, SumB)
    SumCD = _mm256_add_epi32(SumC, SumD)
    Sum = _mm256_add_epi32(SumAB, SumCD)

    Sum = _mm256_hadd_epi32(Sum, Sum)
    Sum = _mm256_hadd_epi32(Sum, Sum)
    SumS = _mm256_permute2x128_si256(Sum, Sum, 1 | (1 << 4))
    Sum = _mm256_add_epi32(Sum, SumS)

    return _mm256_cvtsi256_si32(Sum)

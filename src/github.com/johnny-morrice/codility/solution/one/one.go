package solution

import (
    "encoding/binary"
)

func Solution(N int) int {
    bits := uint(binary.Size(N))
    longest := 0
    count := 0
    for i := uint(0); i < bits; i++ {
        b := int(1) << i
        if N & b == 0 {
            if count > longest {
                longest = count
            }
            count++
        } else {
            count = 0
        }
    }
    return longest
}
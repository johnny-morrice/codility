package solution

func sum(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}

func Solution(A []int) int {
    tail := A[1:]
    fore := A[0]
    aft := sum(tail)
    min := abs(fore - aft)
    for _, n := range tail[:len(tail)-1] {
        fore = fore + n
        aft = aft - n
        diff := abs(fore - aft)
        if diff < min {
            min = diff
        }
    }
    return min
}
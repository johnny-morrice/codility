package solution

func Solution(A []int) int {
    lookup := make(map[int]*int)
    count := make([]*int, len(A))
    for i, a := range A {
        cp, ok := lookup[a]
        if ok {
            *cp++
        } else {
            cp = new(int)
            *cp = 1
            lookup[a] = cp
        }
        count[i] = cp
    }
    for i, cp := range count {
        if *cp % 2 == 1 {
            return A[i]
        }
    }
    panic("Contract broken")
    return 0
}
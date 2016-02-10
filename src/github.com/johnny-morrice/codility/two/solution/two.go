package solution

func Solution(A []int, K int) []int {
    cnt := len(A)
    shift := make([]int, cnt)
    for i, a := range A {
        j := i + K
        j = j % cnt
        shift[j] = a
    }
    return shift
}
package solution

func Solution(A []int, K int) []int {
    cnt := len(A)
    shift := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        j := i + K
        j = j % cnt
        shift[j] = A[i]
    }
    return shift
}
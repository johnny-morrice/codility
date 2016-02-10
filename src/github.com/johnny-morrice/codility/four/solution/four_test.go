package solution

import (
    "testing"
)

func TestSolution(t *testing.T) {
    test([]int{3,1,2,4,3}, 1, t)
    test([]int{-1000, 1000}, 2000, t)
}

func test(input []int, expect int, t *testing.T) {
    actual := Solution(input)
    if expect != actual {
        t.Error("Expected ", expect, "but received", actual)
    }
}
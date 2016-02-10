package solution

import (
    "testing"
)

func TestSolution(t *testing.T) {
    input := []int{9, 3, 9, 3, 9, 7, 9}
    expect := 7
    actual := Solution(input)
    if expect != actual {
        t.Error("Solution(input) != ", expect)
    }
}
package solution

import (
    "testing"
)

func SolutionTest(t *testing.T) {
    orig := []int{3, 8, 9, 7, 6}
    expect := []int{9, 7, 6, 3, 8}
    actual := Solution(orig, 3)
    for i, e := range expect {
        if actual[i] != e {
            t.Fatal("Expected ", actual, "but received ", expect)
        }
    }
}
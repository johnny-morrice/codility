package solution

import (
    "testing"
)

func TestSolution(t *testing.T) {
    if x := Solution(1041); x != 5 {
        t.Error("Solution(1041) != 5 (was ", x , ")")
    }
}
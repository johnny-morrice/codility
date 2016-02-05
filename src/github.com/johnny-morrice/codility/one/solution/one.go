package solution

func Solution(N int) int {
    const maxu = ^uint(0)
    const maxu64 = ^uint64(0)
    bits := uint(32)
    if uint64(maxu) == maxu64 {
        bits = 64
    }
    longest := 0
    count := 0
    lefthit := false
    for i := uint(0); i < bits; i++ {
        b := int(1) << i
        if N & b != 0 {
            if lefthit {
                if count > longest {
                    longest = count
                    count = 0
                }
            } else {
                lefthit = true
            }
        } else {
            count++
        }
    }
    return longest
}
package solution

func Solution(N int) int {
    // Derive int length
    const maxu = ^uint(0)
    const maxu64 = ^uint64(0)
    bits := uint(32)
    if uint64(maxu) == maxu64 {
        bits = 64
    }

    // Count leading/trailing bits
    lead := uint(0)
    trail := uint(0)
    lcnt, tcnt := true, true
    for i := uint(0); i < bits; i++ {
        l := 1 << i
        t := 1 << (bits - 1 - i)
        if N & l == 0 && lcnt {
            lead++
        } else {
            lcnt = false
        }
        if N & t == 0 && tcnt {
            trail++
        } else {
            tcnt = false
        }
    }

    // Remove lead
    N = N >> lead
    max := trail + lead
    // Set trailing high
    for i := uint(0); i < max; i++ {
        b := 1 << (bits - 1 - i)
        N = N | b
    }

    // Hacker's delight
    N = ^N
    k := 0
    for ; N > 0; k++ {
        N = N & (2*N)
    }

    return k
}
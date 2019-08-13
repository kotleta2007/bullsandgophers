package main

import (
	"fmt"
    "math/rand"
    "time"
)

func genRand() (r[4] int) {
    // Seeding the RNG from Unix time
    source := rand.NewSource(time.Now().UnixNano())
    seed := rand.New(source)
    rand.Seed(seed.Int63())
    
    digits := [...] int {0,1,2,3,4,5,6,7,8,9}
    
    rand.Shuffle(len(digits), func(i, j int) {
        digits[i], digits[j] = digits[j], digits[i]
    })
    
    r = [...] int {digits[0], digits[1], digits[2], digits[3]}
    
    return
}

func main() {
	// TODO: Generate random 4-digit number (no repetitions allowed)
    
    fmt.Println(genRand())
}

package main

import (
    "fmt"
    "math/rand"
    "time"
    "errors"
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

func strToArray(str string) (arr[] int, err error) {
    if len(str) != 4 {
        return nil, errors.New("Invalid string size") 
    }
    
    for i := 0; i < len(str); i++ {
        arr = append(arr, int(str[i]) - 48)
    }
    
    return
}

func countBulls(arr1[4] int, arr2[4] int) (int) {
    res := 0
    
    for i := 0; i < 4; i++ {
        if arr1[i] == arr2[i] {
            res++
        }
    }
    
    return res
}

func countCows(arr1[4] int, arr2[4] int) (int) {
    res := -countBulls(arr1, arr2)
    
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if arr1[i] == arr2[j] {
                res++
            }
        }
    }
    
    return res
}

func main() {    
    a := [...] int {1, 2, 3, 4}
    b := [...] int {1, 2, 4, 7}
    
    fmt.Println("Cows:", countCows(a, b))
    fmt.Println("Bulls:", countBulls(a, b))
    // fmt.Println(genRand())
}

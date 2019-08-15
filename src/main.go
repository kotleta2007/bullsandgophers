package main

import (
    "errors"
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

func strToArray(str string) (arr[4] int, err error) {
    if len(str) != 4 {
        return  [4] int {0, 0, 0, 0}, errors.New("Invalid string size")
    }
    
    for i := 0; i < len(str); i++ {
        arr[i] = int(str[i]) - 48
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

func game() {
    tries := 0
    secret := genRand()

    for {
        var userTry string
        fmt.Scanln(&userTry)

        // quit current game
        if userTry == "Q" || userTry == "q" {
           return
        }

        userArray, _ := strToArray(userTry)

        // fmt.Println(secret)
        // fmt.Println(userTry)
        // fmt.Println(userArray)
        // fmt.Println(secret == userArray)

        fmt.Println("Bulls:", countBulls(secret, userArray))
        fmt.Println("Cows:", countCows(secret, userArray))


        tries++

        if (secret == userArray) {
            break
        }
    }

    fmt.Println("You have completed the game in", tries, "tries.")

    return
}

func main() {
    var user_choice byte

    for {
        fmt.Println("\nEnter N to start a new game.")
        fmt.Println("Enter Q to quit.")
        fmt.Scanf("%c\n", &user_choice)

        if user_choice == 'N' || user_choice == 'n' {
            fmt.Println("\n\nNew game!")
            game()
            continue
        } else if user_choice == 'Q' || user_choice == 'q' {
            break
        } else {
            fmt.Println(user_choice, "is not an option. Try again.")
            continue
        }
    }

    fmt.Println("You quit")
    //fmt.Println("Cows:", countCows(a, b))
    //fmt.Println("Bulls:", countBulls(a, b))
}
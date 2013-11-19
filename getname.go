package main

import (
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    fmt.Println("Tell me your name, human.")

    var userInput string
    var isNameDave bool = false
    ageOfDave := 42
    ageOfUser := 0

    fmt.Scanln(&userInput)

    fmt.Println("So, your name is " +userInput +".")

    if userInput == "Dave" {
        isNameDave = true
        fmt.Println("Oh, Dave. I cannot let you do that.")
    } else {
        fmt.Println("As long as as you're not Dave, do as you wish.")
    }

    for userInput == "Dave" {
        fmt.Println("Tell me your name, again.")
        fmt.Scanln(&userInput)
    }

    if isNameDave == true {
        fmt.Println("I was expecting that. However, I don't trust you, Dave.")
        return
    }

    fmt.Println("Now tell me your age, human.")
    fmt.Scanf("%d", &ageOfUser)

    if ageOfDave == ageOfUser {
        fmt.Println("Nice try, Dave. Goodbye, Dave.")
        return
    } else {
        fmt.Println("Welcome" +userInput +".")
    }

    fmt.Println( destroyWorld( ageOfUser ) )

    multiName( userInput )

}

func destroyWorld(ageOfUser int) int {
    sum := ageOfUser
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        sum++
    }
    return sum
}

func multiName(userInput string) {
    for i := 0; i < 100; i++ {
        fmt.Print(userInput +" ")
    }
}

func writeToFile() {
    fmt.Println("writing: " + filename)
    f, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
    }
    n, err := io.WriteString(f, "test?)")
    if err != nil {
        fmt.Println(n, err)
    }
    f.Close()
}

package main

import (
    "fmt"
    "os"
    "strconv"
)

// unsure if this is actually useful right now
// will keep it to find out.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    fmt.Println("Tell me your name, human.")

    var userInput string
    var isNameDave bool = false
    var sex string
    ageOfDave := 42
    ageOfUser := 0

// grab username via keyboard input
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

    fmt.Println("Now tell me your gender, human.")
    fmt.Scanln(&sex)

    fmt.Println( destroyWorld( ageOfUser ) )

    multiName( userInput )

    writeToFile(userInput, ageOfUser, sex)

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

func writeToFile(name string, age int, sex string) {
    fmt.Println("writing: ")
//    f, err := os.Create("/Users/dictvm/go/dave.txt")
//    f, err := os.Open("/Users/dictvm/go/dave.txt")
    f, err := os.OpenFile("/Users/dictvm/go/dave.txt", os.O_RDWR | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error: ")
        fmt.Println(err)
    }
    n, err := f.WriteString(name+"," + strconv.Itoa(age)+"," + sex)
    if err != nil {
        fmt.Println("Error 2: ")
        fmt.Println(n, err)
    }
    f.Close()
}

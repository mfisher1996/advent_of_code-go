package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main(){
    fmt.Println("Day 1")
    //part 1 & 2 
    day_1_file, _ := os.Open("./prod_1")
    scanner := bufio.NewScanner(day_1_file)
    list := []int{}
    b := 0
    for scanner.Scan(){
        line := scanner.Text()
        var a int
        if _,err := fmt.Sscanf(line, "%d", &a); err != nil {
            list = append(list, b)
            b = 0
        } else {
            b = b + a
        }
    }
    slices.Sort(list)
    slices.Reverse[[]int](list)
    fmt.Print("part 1: ")
    fmt.Println(list[0])
    fmt.Print("part 2: ")
    fmt.Println(list[0] + list[1] + list[2])

    fmt.Println("Day 2")
    //part 1 
    day_2_file, _ := os.Open("./prod_2")
    scanner = bufio.NewScanner(day_2_file)
    sum := 0
    for scanner.Scan(){
        line := scanner.Text()
        var opponent, me rune 
        fmt.Sscanf(line, "%c %c", &opponent, &me)
        sum = sum + get_score(opponent, me)
    }
    fmt.Print("part 1: ")
    fmt.Println(sum)
    //part 2
    day_2_file, _ = os.Open("./prod_2")
    scanner = bufio.NewScanner(day_2_file)
    sum = 0
    for scanner.Scan(){
        line := scanner.Text()
        var opponent, me rune 
        fmt.Sscanf(line, "%c %c", &opponent, &me)
        sum = sum + get_outcome(opponent, me)
    }
    fmt.Print("part 2: ")
    fmt.Println(sum)
}

func get_score(o, m rune) int {
    o = o - 65
    m = m - 88
    score := 0
    if (o + 1) % 3 == m {
        score = 6
    } else if o == m {
        score = 3
    }
    return int(m) + score + 1
}

func get_outcome(o, m rune) int {
    switch m {
    case 'X': return ((int(o - 65) + 2) % 3) + 1
    case 'Y': return int(o - 65)  + 4
    case 'Z': return ((int(o - 65) + 1) % 3)  + 7
    default: return 0
    }
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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
        strings.Trim(line, "\n ")
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
        strings.Trim(line, "\n ")
        var opponent, me rune 
        fmt.Sscanf(line, "%c %c", &opponent, &me)
        sum = sum + get_score(opponent, me)
    }
    fmt.Print("part 1: ")
    fmt.Println(sum)

}

func get_score(o, m rune) int {
    //var o, m = int(o), int(m)
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

//func get_outcome(o, m rune) int {
    //opp := int(o)
    //me := switch o
//}

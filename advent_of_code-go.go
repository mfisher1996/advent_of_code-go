package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
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
    
    fmt.Print("part 1: ")
    fmt.Println(max_of_arr(list))
    //fmt.Print("part 2: ")
    //fmt.Println(max_of_arr(list))
}

func max_of_arr(a []int) int {
    max := a[0]
    for _, v := range a {
        if v > max {
            max = v
        }
    }
    return max
}


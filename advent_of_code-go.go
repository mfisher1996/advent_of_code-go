package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main(){

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

    //day_2_file, _ := os.Open("./prod_2")

}

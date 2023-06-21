package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        fmt.Println( "Error while reading file.")
    }
}

func main() {
    file, err := os.Open("day01.txt")
    check(err)

    defer file.Close()

    scanner := bufio.NewScanner(file)

    calories := []int{}
    calorie := 0


    for scanner.Scan() {
        line := scanner.Text()

        line = strings.TrimSpace(line)


        // fmt.Println(line)

        if len(line) != 0 {
            num, err := strconv.Atoi(line)
            check(err)

            calorie += num
        } else {
            calories = append(calories, calorie)
            calorie = 0
        }
    }

    sort.Ints(calories)

    // for _, nums := range calories {
    //     fmt.Println(nums)
    // }

    n := len(calories)

    fmt.Println("Part one: " + fmt.Sprint(calories[n-1]))

    top3 := 0
    for i := n-1; i > n-4; i-- {
        top3 += calories[i]
    }

    fmt.Println("Part two: " + fmt.Sprint(top3))

}

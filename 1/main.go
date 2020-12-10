package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    f,_ := os.Open("input")
    scanner := bufio.NewScanner(f)
    n := []int{}
    for scanner.Scan() {
        l := scanner.Text()
        num, _ := strconv.Atoi(l)
        n = append(n, num)
    }

    for i, v := range n {
        for j, v2 := range n[i + 1:] {
            for _, v3 := range n[j + 1:] {
                if v + v2 + v3 == 2020 {
                    fmt.Printf("%d\n", v * v2 * v3)
                }
            }
        }
    }
}

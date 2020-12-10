package main

import (
	"bufio"
	"fmt"
	"os"
)

// To solve both
type Slope struct {
    x int
    y int
}


func main() {
    f, _ := os.Open("input")
    s := bufio.NewScanner(f)

    matrix := [][]int{}

    index := 0
    for s.Scan() {
        line := s.Text()
        arr := []int{}
        for i := range line {
            if string(line[i]) == "." {
                arr = append(arr, 0)
            }
            if string(line[i]) == "#" {
                arr = append(arr, 1)
            }
        }
        matrix = append(matrix, arr)
        index++
    }

    slopes := []Slope{
        {1, 1},
        {3, 1},
        {5, 1},
        {7, 1},
        {1, 2},
    }
    prot := 1
    for _,s := range slopes {
        prot *= Interation(matrix, s)
    }
    fmt.Println(prot)
}


// A single iter
func Interation(matrix [][]int, slope Slope) int {
    sum := 0
    j := 0
    m := len(matrix[0])
    for i := 0; i < len(matrix); i += slope.y {
        sum += matrix[i][j]
        j = (j + slope.x) % m
    }
    return sum
}

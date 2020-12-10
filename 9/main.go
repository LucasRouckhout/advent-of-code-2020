package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
    file, _ = os.Open("input")
    scanner = bufio.NewScanner(file)
    preambleLength = 25
)

func main() {
    numbers := make([]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
        num, _ := strconv.Atoi(line)
        numbers = append(numbers, num)
    }


    for i, num := range numbers {
        lowerIndex := i - preambleLength
        if lowerIndex < 0 {
            lowerIndex = 0
        }

        arr := numbers[lowerIndex:i]
        _, _, err := FindCouple(arr, num)
        if err == nil {
            continue
        } else {
            fmt.Printf("Did not found a couple for %d\n", num)
        }

    }


}

func FindCouple(arr []int, value int) (int, int, error) {
    for i := range arr {
        for _, v := range arr[i:] {
            if v + arr[i] == value {
                return v, arr[i], nil
            }
        }
    }
    return 0, 0, errors.New("Oops")
}

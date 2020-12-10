package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Represents a Condition.
// A condition has a min and max amount
// of times a letter can occur
type Condition struct {
    min int
    max int
    letter string
}

// Converts a string with structure: %d-%d %s
// to a Condition struct and returns it
func ReadCondition(text string) Condition {
    c := Condition{}
    f := strings.Split(text, " ")
    c.letter = f[1]
    minmax := strings.Split(f[0], "-")
    min, _ := strconv.Atoi(minmax[0])
    max, _ := strconv.Atoi(minmax[1])
    c.min = min
    c.max = max
    return c
}

func main() {

    // Create a scanner from the input file
    f, _ := os.Open("input")
    s := bufio.NewScanner(f)
    
    // Initialize the count and start looping over the lines 
    // of the input file
    count := 0
    for s.Scan() {
        line := s.Text()
        split := strings.Split(line, ":")
        c := ReadCondition(split[0])
        p := strings.Trim(split[1], " ")

        if CheckCondition2(c, p) { count++ }
    }

    fmt.Printf("%d\n", count)
}

// Checks if the given condition holds for the given password in part 1
func CheckCondition(cond Condition, pass string) bool {
    fmt.Printf("Checking pass: %s with %+v\n", pass, cond)
    count := strings.Count(pass, cond.letter)
    return (count <= cond.max) && (count >= cond.min)
}

// Checks if the given condition holds for the given password in part 2
func CheckCondition2(cond Condition, pass string) bool {
    X := string(pass[cond.min - 1]) == cond.letter
    Y := string(pass[cond.max - 1]) == cond.letter
    r := (X || Y) && !(X && Y)
    fmt.Printf("Checking pass: %s with %+v, Result: %t\n", pass, cond, r)
    return r
}


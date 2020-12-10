package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
    byr int
    iyr int
    eyr int
    hgt int
    hcl string
    ecl string
    pid int
    cid int
}

// Checks wheter the Password is valid by checking if any of the values
// was not set to somethine other then the default value. The assumption 
// here is that the none of these default values are actually given as
// a value to use in the problem
func (p Passport) IsValid() bool {
    return p.byr == 0 || p.iyr == 0 || p.eyr == 0 || p.hgt == 0 || p. hcl == "" || p.ecl == "" || p.pid == 0 || p.cid == 0
}

// Assigns the given field value to the field with the given name.
// Just a wrapper arround a switch statement. You can do this with
// Reflection but I was to lazy to look up the API
func (p Passport) Assign(fieldName string, value string) {
    switch fieldName {
    case "byr":
        v, _ := strconv.Atoi(value)
        p.byr = v
    case "iyr":
        v, _ := strconv.Atoi(value)
        p.iyr = v
    case "eyr":
        v, _ := strconv.Atoi(value)
        p.eyr = v
    case "hgt":
        v, _ := strconv.Atoi(value)
        p.hgt = v
    case "hcl":
        p.hcl = value
    case "elc":
        p.ecl = value
    case "pid":
        v, _ := strconv.Atoi(value)
        p.pid = v
    case "cid":
        v, _ := strconv.Atoi(value)
        p.cid = v
    }
}

var (
    file,_      = os.Open("input")
    scanner     = bufio.NewScanner(file)
)

func main() {
    p := Passport{}
    count := 0
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" && p.IsValid() {
            fmt.Printf("%+v\n", p)
            count++
            p = Passport{}
            continue
        }
        fields := strings.Split(line, " ")
        for _, field := range fields {
            o := strings.Split(field, ":")
            if len(o) == 1 {
                p.Assign(o[0], "")
            } else {
                p.Assign(o[0], o[1])
            }

        }
    }
    fmt.Println(count)
}

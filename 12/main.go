package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
)

// Represents an action the ship can take
type Action struct {
    Instruction string
    Value int
}

// Creates an Action from a string line of the form
// SX Where S is a letter and X is an integer value
// It checks that the whole string is valid UTF-8 runes
// because it uses the DecodeRuneInString function to get the 
// first "character" 
func NewAction(line string) (Action, error) {
    action := Action{}
    isValid := utf8.ValidString(line)
    if !isValid {
        return action, errors.New("Action string contains invalid encoded UTF-8")
    }

    instructionRune, length := utf8.DecodeRuneInString(line)
    v, err := strconv.Atoi(line[length:])

    if err != nil {
        return action, err
    }

    action.Instruction = string(instructionRune)
    action.Value = v
    return action, nil
}


// Log out the action and its values
func (a Action) Print() {
    log.Printf("%+v\n", a)
}

func main() {
    f, _ := os.Open("input")
    s := bufio.NewScanner(f)

    actions := []Action{}
    for s.Scan() {
        line := s.Text()
        action, err := NewAction(line)
        if err != nil { log.Fatal(err) }
        actions = append(actions, action)
    }

    angle := 0
    northDistance := 0
    eastDistance := 0
    for _, a := range actions {
        fmt.Printf("NorthD: %d, EastD: %d, Angle: %d\n", northDistance, eastDistance, angle)
        fmt.Printf("%+v\n", a)
        switch a.Instruction {
        case "N":
            northDistance += a.Value
        case "S":
            northDistance -= a.Value
        case "W":
            eastDistance -= a.Value
        case "E":
            eastDistance += a.Value
        case "R":
            angle = (angle - a.Value) % 360
        case "L":
            angle = (angle + a.Value) % 360
        case "F":

            absAngle := int(math.Abs(float64(angle)))
            switch absAngle {
            case 0:
                eastDistance += a.Value
            case 90:
                northDistance += a.Value
            case 180:
                eastDistance -= a.Value
            case 270:
                northDistance -= a.Value
            }
        }

    }
    northDistance = int(math.Abs(float64(northDistance)))
    eastDistance = int(math.Abs(float64(eastDistance)))
    println(northDistance + eastDistance)

}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
    file, _             = os.Open("input")
    scanner             = bufio.NewScanner(file)
)

func main() {

    // Load in the instructions into memory.
    instructions := make([]string, 0)
    for scanner.Scan() {
        instructions = append(instructions, scanner.Text())
    }

    for i := 0; i < len(instructions); i++ {
        newInstructions := make([]string, len(instructions))
        copy(newInstructions, instructions)

        insu, _ := UnpackInstruction(newInstructions[i])

        if insu == "jmp" {
            newInstructions[i] = strings.ReplaceAll(newInstructions[i], "jmp", "nop")
        } else if insu == "nop" {
            newInstructions[i] = strings.ReplaceAll(newInstructions[i], "nop", "jmp")
        } else {
            continue
        }

        instructionPointer := 0
        accumulator := 0
        count := 0
        for instructionPointer < len(newInstructions) && count < 10000 {
            ins, value := UnpackInstruction(newInstructions[instructionPointer])
            switch ins {
            case "nop":
                instructionPointer++
            case "jmp":
                instructionPointer += value
            case "acc":
                accumulator += value
                instructionPointer++
            }
            count++
        }

        if count != 10000 {
            println(accumulator)
        }

    }

}


func UnpackInstruction(ins string) (string, int) {
    splits := strings.Split(ins, " ")
    value, err := strconv.Atoi(splits[1])
    if err != nil {
        log.Fatal(err)
    }
    return splits[0], value
}

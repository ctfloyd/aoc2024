package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type CPU struct {
	a, b, c, pc int
	output      []int
}

type OpCode int

const (
	OpcodeDivisionA = iota
	OpcodeBitwiseXorLiteral
	OpcodeModulo
	OpcodeJumpNotZero
	OpcodeBitwiseXorBAndC
	OpcodeOut
	OpcodeDivisionB
	OpcodeDivisionC
)

type Instruction struct {
	opcode  OpCode
	operand int
}

func (cpu *CPU) ClearOutput() {
	cpu.output = []int{}
}

func (cpu *CPU) Execute(instructions []Instruction) []int {
	cpu.ClearOutput()
	cpu.pc = 0
	for cpu.pc < len(instructions) {
		instr := instructions[cpu.pc]
		if instr.opcode == OpcodeDivisionA {
			cpu.divisionA(instr)
		} else if instr.opcode == OpcodeBitwiseXorLiteral {
			cpu.bitwiseXorLiteral(instr)
		} else if instr.opcode == OpcodeModulo {
			cpu.modulo(instr)
		} else if instr.opcode == OpcodeJumpNotZero {
			cpu.jumpNotZero(instr)
		} else if instr.opcode == OpcodeBitwiseXorBAndC {
			cpu.bitwiseXorBAndC(instr)
		} else if instr.opcode == OpcodeOut {
			cpu.out(instr)
		} else if instr.opcode == OpcodeDivisionB {
			cpu.divisionB(instr)
		} else if instr.opcode == OpcodeDivisionC {
			cpu.divisionC(instr)
		} else {
			fmt.Printf("invalid instruction %v\n", instr)
		}
	}
	return cpu.output
}

func (cpu *CPU) divisionA(instr Instruction) {
	numerator := cpu.a
	denominator := int(math.Pow(2, float64(cpu.getValueForComboOperand(instr.operand))))
	cpu.a = numerator / denominator
	cpu.pc++
}

func (cpu *CPU) bitwiseXorLiteral(instr Instruction) {
	cpu.b = cpu.b ^ instr.operand
	cpu.pc++
}

func (cpu *CPU) modulo(instr Instruction) {
	cpu.b = cpu.getValueForComboOperand(instr.operand) % 8
	cpu.pc++
}

func (cpu *CPU) jumpNotZero(instr Instruction) {
	if cpu.a == 0 {
		cpu.pc++
		return
	}
	cpu.pc = instr.operand
}

func (cpu *CPU) bitwiseXorBAndC(instr Instruction) {
	cpu.b = cpu.b ^ cpu.c
	cpu.pc++
}

func (cpu *CPU) out(instr Instruction) {
	v := cpu.getValueForComboOperand(instr.operand) % 8
	cpu.output = append(cpu.output, v)
	cpu.pc++
}

func (cpu *CPU) divisionB(instr Instruction) {
	numerator := cpu.a
	denominator := int(math.Pow(2, float64(cpu.getValueForComboOperand(instr.operand))))
	cpu.b = numerator / denominator
	cpu.pc++
}

func (cpu *CPU) divisionC(instr Instruction) {
	numerator := cpu.a
	denominator := int(math.Pow(2, float64(cpu.getValueForComboOperand(instr.operand))))
	cpu.c = numerator / denominator
	cpu.pc++
}

func (cpu *CPU) getValueForComboOperand(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}

	if operand == 4 {
		return cpu.a
	}

	if operand == 5 {
		return cpu.b
	}

	if operand == 6 {
		return cpu.c
	}

	fmt.Printf("invalid value for combo operand %v\n", operand)
	return -1
}

func readInput(input string) (CPU, []Instruction) {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	cpu := CPU{}
	instructions := []Instruction{}

	r, _ := regexp.Compile(`\d+`)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		line := scanner.Text()
		m := r.FindString(line)

		v, _ := strconv.Atoi(m)
		if i == 0 {
			cpu.a = v
		} else if i == 1 {
			cpu.b = v
		} else if i == 2 {
			cpu.c = v
		}
	}

	scanner.Scan()
	scanner.Scan()
	line := scanner.Text()
	m := r.FindAllString(line, -1)

	for i := 0; i < len(m); i += 2 {
		c, _ := strconv.Atoi(m[i])
		o, _ := strconv.Atoi(m[i+1])
		instructions = append(instructions, Instruction{opcode: OpCode(c), operand: o})
	}
	return cpu, instructions
}

func partOne(input string) {
	cpu, instrs := readInput(input)
	cpu.Execute(instrs)
	fmt.Printf("result: ")
	for _, v := range cpu.output {
		fmt.Printf("%v,", v)
	}
	fmt.Println()
}

func partTwo(input string) {
	cpu, instrs := readInput(input)
	seed := 0

	target := []int{}
	for _, i := range instrs {
		target = append(target, int(i.opcode))
		target = append(target, i.operand)
	}

	for itr := len(target) - 1; itr >= 0; itr-- {
		seed <<= 3
		for !slices.Equal(exec(cpu, instrs, seed), target[itr:]) {
			seed++
		}
	}

	fmt.Printf("result: %v\n", seed)
}

func exec(cpu CPU, instrs []Instruction, a int) []int {
	cpu.a = a
	return cpu.Execute(instrs)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}

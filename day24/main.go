package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Wire struct {
	name  string
	value int
	iv    int
}

func (wire *Wire) reset() {
	wire.value = wire.iv
}

type GateKind int

const (
	GateKindAnd = iota
	GateKindOr
	GateKindXor
)

type Gate struct {
	kind       GateKind
	aWire      *Wire
	bWire      *Wire
	oWire      *Wire
	evalulated bool
}

func (gate *Gate) canEvaluate() bool {
	return !gate.evalulated && gate.aWire.value >= 0 && gate.bWire.value >= 0
}

func (gate *Gate) evaluate() {
	if gate.kind == GateKindAnd {
		gate.oWire.value = gate.aWire.value & gate.bWire.value
	} else if gate.kind == GateKindOr {
		gate.oWire.value = gate.aWire.value | gate.bWire.value
	} else if gate.kind == GateKindXor {
		gate.oWire.value = gate.aWire.value ^ gate.bWire.value
	}
	gate.evalulated = true
}

func readInput(input string) (map[string]*Wire, []Gate) {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	m := make(map[string]*Wire)
	g := []Gate{}

	wires := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			wires = false
		} else if wires {
			parts := strings.Split(line, ": ")
			value, _ := strconv.Atoi(parts[1])
			wire := Wire{
				name:  parts[0],
				value: value,
				iv:    value,
			}
			m[parts[0]] = &wire
		} else {
			parts := strings.Split(line, " ")
			aWire, ok := m[parts[0]]
			if !ok {
				aWire = &Wire{
					name:  parts[0],
					value: -1,
					iv:    -1,
				}
				m[parts[0]] = aWire
			}
			bWire, ok := m[parts[2]]
			if !ok {
				bWire = &Wire{
					name:  parts[2],
					value: -1,
					iv:    -1,
				}
				m[parts[2]] = bWire
			}
			oWire, ok := m[parts[4]]
			if !ok {
				oWire = &Wire{
					name:  parts[4],
					value: -1,
					iv:    -1,
				}
				m[parts[4]] = oWire
			}

			var kind GateKind = GateKindAnd
			if parts[1] == "OR" {
				kind = GateKindOr
			} else if parts[1] == "XOR" {
				kind = GateKindXor
			}
			gate := Gate{
				kind:       kind,
				aWire:      aWire,
				bWire:      bWire,
				oWire:      oWire,
				evalulated: false,
			}
			g = append(g, gate)
		}
	}
	return m, g
}

func wireAsNumber(wires map[string]*Wire, wire byte) int {
	keys := slices.Collect(maps.Keys(wires))
	slices.SortFunc(keys, func(a, b string) int {
		return strings.Compare(b, a)
	})

	result := 0
	for _, k := range keys {
		if k[0] == wire {
			result = (result << 1) | wires[k].value
		}
	}
	return result
}

func run(gates []Gate) {
	moreWork := true
	for moreWork {
		moreWork = false

		for i := 0; i < len(gates); i++ {
			if gates[i].canEvaluate() {
				moreWork = true
				gates[i].evaluate()
			}
		}
	}
}

func partOne(input string) {
	wires, gates := readInput(input)
	run(gates)
	fmt.Printf("result %v\n", wireAsNumber(wires, 'z'))
}

func verifyZ(formulas map[string]*Gate, wire string, num int) bool {
	gate, ok := formulas[wire]
	if !ok {
		return false
	}

	if gate.kind != GateKindXor {
		return false
	}

	x := gate.aWire.name
	y := gate.bWire.name
	if num == 0 {
		return sort([2]string{x, y}) == [2]string{"x00", "y00"}
	}

	return verifyIntermediateXor(formulas, x, num) && verifyCarry(formulas, y, num) || verifyIntermediateXor(formulas, y, num) && verifyCarry(formulas, x, num)
}

func verifyIntermediateXor(formulas map[string]*Gate, wire string, num int) bool {
	gate, ok := formulas[wire]
	if !ok {
		return false
	}

	if gate.kind != GateKindXor {
		return false
	}

	x := gate.aWire.name
	y := gate.bWire.name
	return sort([2]string{x, y}) == [2]string{name('x', num), name('y', num)}
}

func verifyCarry(formulas map[string]*Gate, wire string, num int) bool {
	gate, ok := formulas[wire]
	if !ok {
		return false
	}

	x := gate.aWire.name
	y := gate.bWire.name
	if num == 1 {
		if gate.kind != GateKindAnd {
			return false
		}
		return sort([2]string{x, y}) == [2]string{"x00", "y00"}
	}
	if gate.kind != GateKindOr {
		return false
	}

	return verifyDirectCarry(formulas, x, num-1) && verifyRecarry(formulas, y, num-1) || verifyDirectCarry(formulas, y, num-1) && verifyRecarry(formulas, x, num-1)
}

func verifyDirectCarry(formulas map[string]*Gate, wire string, num int) bool {
	gate, ok := formulas[wire]
	if !ok {
		return false
	}

	if gate.kind != GateKindAnd {
		return false
	}

	x := gate.aWire.name
	y := gate.bWire.name
	return sort([2]string{x, y}) == [2]string{name('x', num), name('y', num)}
}

func verifyRecarry(formulas map[string]*Gate, wire string, num int) bool {
	gate, ok := formulas[wire]
	if !ok {
		return false
	}

	if gate.kind != GateKindAnd {
		return false
	}

	x := gate.aWire.name
	y := gate.bWire.name
	return verifyIntermediateXor(formulas, x, num) && verifyCarry(formulas, y, num) || verifyIntermediateXor(formulas, y, num) && verifyCarry(formulas, x, num)

}

func sort(slice [2]string) [2]string {
	if slice[0] < slice[1] {
		return slice
	} else {
		return [2]string{slice[1], slice[0]}
	}
}

func name(char byte, num int) string {
	if num < 10 {
		return fmt.Sprintf("%s0%d", string(char), num)
	} else {
		return fmt.Sprintf("%s%d", string(char), num)
	}
}

func verify(formulas map[string]*Gate, i int) bool {
	return verifyZ(formulas, name('z', i), i)
}

func progress(formulas map[string]*Gate) int {
	i := 0
	for {
		if !verify(formulas, i) {
			break
		}
		i += 1
	}
	return i
}

func makeFormulas(gates []Gate) map[string]*Gate {
	m := make(map[string]*Gate)
	for i := 0; i < len(gates); i++ {
		g := gates[i]
		m[g.oWire.name] = &g
	}
	return m
}

func partTwo(input string) {
	_, gates := readInput(input)
	formulas := makeFormulas(gates)

	swaps := []string{}
	for range 4 {
		baseline := progress(formulas)
		for x := range formulas {
			swapped := false
			for y := range formulas {
				if x == y {
					continue
				}

				formulas[x], formulas[y] = formulas[y], formulas[x]
				if progress(formulas) > baseline {
					swaps = append(swaps, x)
					swaps = append(swaps, y)
					swapped = true
					break
				}
				formulas[x], formulas[y] = formulas[y], formulas[x]
			}
			if swapped {
				break
			}
		}
	}
	slices.Sort(swaps)
	fmt.Printf("swaps %v\n", strings.Join(swaps, ","))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}

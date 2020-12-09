package day8

import (
  "log"
  "strings"
  "strconv"
)

type Instructions struct {
  List []*Instruction
  Accumulator int
  Traverser int
}

//Traverse returns the value of accumulator right before executing an already visited instruction.
func (i Instructions) Traverse() int {
  if i.List[i.Traverser].Visited == true { return i.Accumulator}
  i.Accumulator, i.Traverser = i.List[i.Traverser].Step(i.Accumulator, i.Traverser)
  return i.Traverse()

}

//Return the value of accumulator for the permutation of i.List where changing exactly one Nop for a Jump or vice versa results in the i.Traverser reaching the end of the List instead of reaching an already visited instruction.
func (i Instructions) TraverseTerminal() int {
  for _, instruction := range i.List {
    i.cleanTraversals()
    originalOp := instruction.Op
    if instruction.Op == Jump {
      instruction.Op = Nop
    } else if instruction.Op == Nop {
      instruction.Op = Jump
    } else { continue }
    list := []Operation{}
    for _,inst := range i.List {
      list = append(list, inst.Op)
    }
    result, ok := i.checkTerminalTraversal()
    if ok { return result }
    instruction.Op = originalOp
  }
  log.Fatalf("No working instruction permutations found")
  return -99
}

func (i Instructions) checkTerminalTraversal() (int, bool) {
  if i.Traverser == len(i.List) {
     return i.Accumulator, true
     }
  if i.List[i.Traverser].Visited == true {
     return i.Accumulator, false
     }
  
  i.Accumulator, i.Traverser = i.List[i.Traverser].Step(i.Accumulator, i.Traverser)
  return i.checkTerminalTraversal()
}

func (i Instructions) cleanTraversals() {
  for _, inst := range i.List {
    inst.Visited = false
  }
}

type Instruction struct {
  Op Operation
  Arg int
  Visited bool
}

//Given the current accumulator and traverser, it will return the new value of accumulator and traverser for this instruction.
func (i *Instruction) Step(currAcc int, currTrav int) (int, int) {
  i.Visited = true
  switch i.Op {
    case Accumulate:return currAcc + i.Arg, currTrav + 1
    case Jump:return currAcc, currTrav + i.Arg
    case Nop: return currAcc, currTrav + 1
    default: return currAcc, currTrav + 1
  }
}

type Operation string

const (
  Accumulate Operation = "acc"
  Jump Operation = "jmp"
  Nop Operation = "nop"
)

func NewInstructions(data []string) Instructions {
  instructions := Instructions{
    List: make([]*Instruction, len(data)),
    Accumulator: 0,
    Traverser: 0,
  }
  for i, inst := range data {
    split := strings.Split(inst, " ")
    action := split[0]
    amount, err := strconv.ParseInt(split[1], 10, 64)
    if err != nil {
      log.Fatalf("Error converting %s to int: %s", inst, err)
    }
    instructions.List[i] = &Instruction{
      Op: ParseOperation(action),
      Arg: int(amount),
      Visited: false,
    }

  }
  return instructions
}

func ParseOperation(action string) Operation {
  switch action {
    case "acc":
    return Accumulate
    case "jmp":
    return Jump
    case "nop":
    return Nop
    default:
    log.Fatalf("Action %s couldn't be converted to an operation", action)
  }
  return Nop
}
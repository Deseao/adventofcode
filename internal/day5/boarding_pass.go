package day5

type BoardingPass struct {
  Row int
  Column int
}

func NewBoardingPass(input string) BoardingPass {
  row := computeRow(input[:7])
  column := computeColumn(input[7:])
  return BoardingPass{
    Row: row,
    Column: column,
  }
}

func computeRow(row string) int {
rowSequence := ConvertRowToSequence(row)
return ComputePartition(rowSequence, 0, 127)
}

func computeColumn(column string) int {
  columnSequence := ConvertColumnToSequence(column)
  return ComputePartition(columnSequence, 0, 7)
}

func ConvertRowToSequence(row string) Sequence {
  directions := make([]Direction, len(row))
  for k, v := range row {
    directions[k] = RowToDirection(v)
  }
  return Sequence{
    List: directions,
  }
}

func ConvertColumnToSequence(column string) Sequence {
  directions := make([]Direction, len(column))
  for k, v := range column {
    directions[k] = ColumnToDirection(v)
  }
  return Sequence{
    List: directions,
  }
}

func RowToDirection(char rune) Direction {
  if char == 'F' { return Lower }
  return Upper
}

func ColumnToDirection(char rune) Direction {
  if char == 'L' { return Lower }
  return Upper
}

func (b BoardingPass) ID() int {
  return b.Row * 8 + b.Column
}

func ComputePartition(seq Sequence, lower int, upper int) int {
  directions := seq.List
  if len(directions) == 1 {
    if directions[0] == Lower { return lower }
    return upper
  }
 
 midPoint := (((upper + 1) - lower) / 2) + lower
  if directions[0] == Lower {
    return ComputePartition(Sequence{List: directions[1:]}, lower, midPoint - 1)
  }
  if directions[0] == Upper {
    
    return ComputePartition(Sequence{List: directions[1:]}, midPoint, upper)
  }
  return 0
}
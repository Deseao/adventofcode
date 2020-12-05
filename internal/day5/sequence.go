package day5

type Sequence struct {
  List []Direction
}

type Direction int

const (
  Lower Direction = iota
  Upper
)

func (d Direction) String() string {
  return [...]string{"Lower", "Upper"}[d]
}
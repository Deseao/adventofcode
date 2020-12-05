package day5_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "advent/internal/day5"
)

var _ = Describe("BoardingPass", func() {
  Context("NewBoardingPass", func() {
    It("Should split the input into the correct row and column information", func() {
      pass := NewBoardingPass("BFFFBBFRRR")

      Expect(pass.Row).To(Equal(70))
      Expect(pass.Column).To(Equal(7))

      pass = NewBoardingPass("FBFBBFFRLR")
      Expect(pass.Row).To(Equal(44))
      Expect(pass.Column).To(Equal(5))
    })
  })
  Context("ConvertRowToSequence", func() {
    It("Should build the correct sequence", func() {
      result := ConvertRowToSequence("BFFFBBF")
      Expect(result).To(Equal(Sequence{List: []Direction{Upper, Lower, Lower, Lower, Upper, Upper, Lower}}))
    })
  })
  Context("RowToDirection", func() {
    It("Should convert F to Lower", func() {
      Expect(RowToDirection('F')).To(Equal(Lower))
    })
    It("Should convert B to Upper", func() {
      Expect(RowToDirection('B')).To(Equal(Upper))
    })
  })
  Context("ComputePartition", func() {
    It("Solves the base case", func() {
      baseSequence := Sequence{List: []Direction{Lower}}
      result := ComputePartition(baseSequence, 44, 45)
      Expect(result).To(Equal(44))
      upperSequence := Sequence{List: []Direction{Upper}}
      result = ComputePartition(upperSequence, 44, 45)
      Expect(result).To(Equal(45))
    })
    It("Recurses on halves properly", func() {
      input := Sequence{List: []Direction{Lower, Upper, Lower, Upper,Upper, Lower, Lower}}
      result := ComputePartition(input, 0, 127)
      Expect(result).To(Equal(44))
    })
  })
  Context("ID", func() {
    It("Should return the correct ID", func() {
      pass := BoardingPass{Row: 44, Column: 5}
      Expect(pass.ID()).To(Equal(357))
    })
  })
})
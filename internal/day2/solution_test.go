package day2_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "advent/internal/day2"
)

var _ = Describe("ParseInput", func() {
  var testInput []string
  BeforeEach(func() {
    testInput = append(testInput, "1-1 a: MOCK_INPUT")
    testInput = append(testInput, "1-2 b: MOCK_INPUT_2")
    testInput = append(testInput, "10-40 c: MOCK_INPUT_3")
  })
  It("Should convert the input into the correct passwords and policies", func() {
    entries, err := ParseInput(testInput)
    Expect(err).NotTo(HaveOccurred())
    Expect(entries).To(HaveLen(len(testInput)))
    Expect(entries[0].Policy.Minimum).To(Equal(1))
    Expect(entries[0].Policy.Maximum).To(Equal(1))
    Expect(entries[0].Policy.Letter).To(Equal('a'))
    Expect(entries[0].Password).To(Equal("MOCK_INPUT"))
    Expect(entries[1].Policy.Minimum).To(Equal(1))
    Expect(entries[1].Policy.Maximum).To(Equal(2))
    Expect(entries[1].Policy.Letter).To(Equal('b'))
    Expect(entries[1].Password).To(Equal("MOCK_INPUT_2"))
    Expect(entries[2].Policy.Minimum).To(Equal(10))
    Expect(entries[2].Policy.Maximum).To(Equal(40))
    Expect(entries[2].Policy.Letter).To(Equal('c'))
    Expect(entries[2].Password).To(Equal("MOCK_INPUT_3"))
  })
})
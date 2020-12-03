package day2_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "advent/internal/day2"
)

var _ = Describe("IsValid", func() {
  var entry Entry
  BeforeEach(func() {
    entry = Entry{
      Policy: Policy{
        Minimum: 1,
        Maximum: 2,
        Letter: 'z',
      },
    }
  })
  When("the password matches the policy", func() {
    BeforeEach(func() {
      entry.Password = "abcdefghijklmnopqrstuvwxyz"
    })
    It("should return true", func() {
        Expect(entry.IsValid()).To(BeTrue())
    })
  })
  When("the password doesn't match the policy", func() {
    BeforeEach(func() {
      entry.Password = "abcdefghijklmnopqrstuvwxy"
    })
    It("should return false", func() {
        Expect(entry.IsValid()).To(BeFalse())
    })
  })
})

var _ Describe("IsPositionallyValid", func() {
  var entry Entry
   BeforeEach(func() {
    entry = Entry{
      Policy: Policy{
        Minimum: 1,
        Maximum: 2,
        Letter: 'z',
      },
    }
  })
  When("the password matches the policy", func() {
    BeforeEach(func() {
      entry.Password = "zbcdefghijklmnopqrstuvwxyz"
    })
    It("should return true", func() {
        Expect(entry.IsPositionallyValid()).To(BeTrue())
    })
  })
  When("the password doesn't match the policy", func() {
    BeforeEach(func() {
      entry.Password = "abcdefghijklmnopqrstuvwxy"
    })
    It("should return false", func() {
        Expect(entry.IsValid()).To(BeFalse())
    })
  })
})
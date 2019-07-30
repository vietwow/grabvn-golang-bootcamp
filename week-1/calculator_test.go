package calculator

import (
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "testing"
)

type randomMock struct {
  mock.Mock
}

func TestAdd(t *testing.T) {
  calc := newCalculator()
  assert.Equal(t, 9, calc.Add(5, 4))
}

func TestSubtract(t *testing.T) {
  calc := newCalculator()
  assert.Equal(t, 1, calc.Subtract(5, 4))
}

func TestMultiply(t *testing.T) {
  calc := newCalculator()
  assert.Equal(t, 20, calc.Multiply(5, 4))
}

func TestDivide(t *testing.T) {
  calc := newCalculator()
  assert.Equal(t, 5, calc.Divide(20, 4))
}


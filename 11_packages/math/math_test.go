package math

import "testing"

type testpair struct {
  values []float64
  average float64
}

var tests = []testpair{
  { []float64{1, 2}, 1.5 },
  { []float64{1, 1, 1, 1, 1, 1}, 1 },
  { []float64{-1, 1}, 0 },
}

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}

func TestMax(t *testing.T)  {
  max := Max([]float64{1.0, 2.0, 3.0, 4.0})
  if max != 4.0 {
    t.Error("Expected 4, got", max)
  }
}

func TestMin(t *testing.T)  {
  min:= Min([]float64{1.0, 2.0, 3.0, 4.0})
  if min != 1.0 {
    t.Error("Expected 4, got", min)
  }
}

package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  if len(xs) == 0 {
    return 0
  }
  return total / float64(len(xs))
}

// Finds the max value of a slice of float64s
func Max(x []float64) float64 {
  highest := x[0]
  for _, num := range x {
    if num > highest {
      highest = num
    }
  }
  return highest
}

// Finds the min value of a slice of float64s
func Min(x []float64) float64 {
  lowest := x[0]
  for _, value := range x {
    if value < lowest {
      lowest = value
    }
  }
  return lowest
}

package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(s []string, i int) int {
  if i == 0 {
    i = 2
  }
  return len(s) * i
}
// TODO: define the 'Quantities()' function
func Quantities(s []string) (int, float64) {
  a := 0
  b := 0
  for i := range s {
    if s[i] == "sauce" {
      a += 1
    }
    if s[i] == "noodles" {
      b += 1
    }
  }
  c := float64(a) * 0.2
  b = b * 50 
  return b, c
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(s1, s2 []string) {
  i := len(s1) - 1
  t := s1[i]
  s2[len(s2)-1] = t
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(f []float64, s int) []float64 {
var c []float64
  for i := range f {
    c = append(c, f[i] * float64(s)/2)
  }
  return c
}
>>>>>>> lasagna-master

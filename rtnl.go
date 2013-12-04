package rtnl

import (
   "fmt"
)

// struct to represent a rational number
type Nbr struct {
   Numer int
   Denom int
}

// returns string representation of a rational
func (r Nbr) String() string {
   var s string
   if abs(r.Denom) != 1 && r.Numer != 0 {
      s = fmt.Sprintf(s + "%v/%v", r.Numer, r.Denom)
   } else {
      s = fmt.Sprintf(s + "%v", r.Numer)
   }
   return s
}

// simplify rational
func Simplify(r Nbr) Nbr {
   cf := gcf(uint(abs(r.Numer)), uint(abs(r.Denom)))
   result := Nbr{r.Numer / int(cf), r.Denom / int(cf)}

   if result.Denom < 0 {
      result.Numer = -result.Numer
      result.Denom = -result.Denom
   }
   return result
}

// compares rationals and returns:
//   -1 if lh < rh
//    0 if lh = rh
//    1 if lh > rh
func Compare(lh, rh Nbr) int {
   cp1 := lh.Numer * rh.Denom
   cp2 := rh.Numer * lh.Denom
   if cp1 < cp2 {
      return -1
   } else if cp1 > cp2 {
      return 1
   }
   return 0
}

// multiply rationals
func Mult(lh, rh Nbr) Nbr {
   result := Nbr{lh.Numer * rh.Numer, lh.Denom * rh.Denom}
   result = Simplify(result)
   return result
}

// divide rationals
func Div(lh, rh Nbr) Nbr {
   result := Nbr{lh.Numer * rh.Denom, lh.Denom * rh.Numer}
   result = Simplify(result)
   return result
}

// add rationals
func Add(lh, rh Nbr) Nbr {
   var r3 Nbr
   if lh.Denom != rh.Denom {
      cm := lcm(uint(lh.Denom), uint(rh.Denom))
      r3.Denom = int(cm)
      r3.Numer = lh.Numer * (r3.Denom / lh.Denom) + rh.Numer * (r3.Denom / rh.Denom)
   } else {
      r3.Numer = lh.Numer + rh.Numer
      r3.Denom = lh.Denom
   }
   r3 = Simplify(r3)
   return r3
}

// subtract rationals
func Sub(lh, rh Nbr) Nbr {
   rh.Numer = -rh.Numer
   return Add(lh, rh)
}

// invert rational
func Inv(r Nbr) {
   result := Nbr{r.Denom, r.Numer}
   if result.Denom < 0 {
      result.Numer = -result.Numer
      result.Denom = -result.Denom
   }
}

// find the least common multiple - taken from Wikipedia
func lcm(x, y uint) uint {
   return (x * y)/gcf(x, y)
}

// count trailing zero bits - taken from Wikipedia
func ctz(x uint) uint {
   if x == 0 {
      return 32
   }
   var n uint = 0
   if (x & 0x0000FFFF) == 0 {
      n += 16
      x = x >> 16
   }
   if (x & 0x000000FF) == 0 {
      n += 8
      x = x >> 8
   }
   if (x & 0x0000000F) == 0 {
      n += 4
      x = x >> 4
   }
   if (x & 0x00000003) == 0 {
      n += 2
      x = x >> 2
   }
   if (x & 0x00000001) == 0 {
      n += 1
   }
   return n
}

// find the greatest common factor - taken from Wikipedia
func gcf(x, y uint) uint {
   if x == 0 {
      return y
   } else if y == 0 {
      return x
   }

   cf2 := ctz(x | y)
   x = x >> ctz(x)

   for {
      y = y >> ctz(y)
      if x == y {
         break;
      } else if x > y {
         tmp := x
         x = y
         y = tmp
      } else if x == 1 {
         break
      }
      y -= x
   }
   return x << cf2
}

// returns absolute value of an int
func abs(i int) int {
   if i < 0 {
      return -i
   }
   return i
}

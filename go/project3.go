package main
import ("fmt"
        "math")

/* The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ? */

func main() {
  num := 600851475143
  // factor all 2's
  for ; num % 2 == 0; num = num / 2 {}
  // check all odd factors up to ceil sqrt of num
  for i := 3; i <= int(math.Sqrt(float64(num))+1); i=i+2 {
    // if factor, recheck all factors, unless factor equals index
    for ; num % i == 0 && i != num; num = num / i {}
  }
  fmt.Println(num)
}

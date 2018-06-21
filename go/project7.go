package main
import ("fmt"
        "math"
        "time"
        "log")

/* By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10,001st prime number? */

func check_prime(i int) bool {
  for j := 3; j < int(math.Sqrt(float64(i))+1); j+=2 {
    if i % j == 0 {
      return false
    }
  }
  return true
}

func main() {
  defer TimeTaken(time.Now(), "main()")

  count := 6 // according to scope of the problem, we know 2, 3, 5, 7, 11, and 13 are prime
  i := 15 // 13 is prime, so increment two and start looping

  for ; count < 10001; i+=2 {
    if check_prime(i) == true {
      count++
    }
  }
  fmt.Println(i)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

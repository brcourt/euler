package main
import ("fmt"
        "time"
        "log")

/* The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million */

func check_prime(i int) bool {
  // 6k ± 1 optimization
  if i <= 1 || i % 2 == 0 || i % 3 == 0 {
    return false
  }
  for j := 5; j*j <= i; j += 6 {
    if i % j == 0 || i % (j + 2) == 0 {
      return false
    }
  }
  return true
}

func main() {
  defer TimeTaken(time.Now(), "main()")

  num := 2000000 // two million
  sum := 17 // Problem already states sum of all primes below 10 is 17

  // start at 11 because we alread know the sum of the first 10 primes
  for i := 11; i < num; i=i+2 {
    if check_prime(i) {
      sum += i
    }
  }
  fmt.Println(sum)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

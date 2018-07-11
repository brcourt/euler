package main
import ("fmt"
        "time"
        "log")

/* By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10,001st prime number? */

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

  count := 6 // according to scope of the problem, we know 2, 3, 5, 7, 11, and 13 are prime
  i := 15 // 13 is prime, so increment two and start looping

  for {
    if check_prime(i) == true {
      count++
      if count == 10001 {
        fmt.Println(i)
        break
      }
    }
    i+=2
  }
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

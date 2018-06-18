package main
import ("fmt"
        "time"
        "log")

/* 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20? */


/* this solution is garbage, but it was fun nonetheless. We increment by 20, since we know it must end with 0 in order to be divisible by 20 and the 10s place is even. This takes care of 2, 5, and 10 as well. After that, we check all primes. Assuming all primes pass, we check the fewest number of composite numbers. The comments in each if statement show what numbers are being proven to be divisble by. */

func main() {
  defer TimeTaken(time.Now(), "main()")

  for i :=20; ; i+=20 {
    // 20, 10, 2, 5
    if i % 19 != 0 {
      // 19
      continue
    } else if i % 17 != 0 {
      // 17
      continue
    } else if i % 13 != 0 {
      // 13
      continue
    } else if i % 11 != 0 {
      // 11
      continue
    } else if i % 7 != 0 {
      // 7, 14
      continue
    } else if i % 16 != 0 {
      // 16
      continue
    } else if i % 18 != 0 {
      // 18
      continue
    } else if i % 9 != 0 {
      // 3, 9, 15, 18, 6
      continue
    } else if i % 4 != 0 {
      // 4, 12
      continue
    } else if i % 8 != 0 {
      // 8
      continue
    } else {
      fmt.Println(i)
      break
    }
  }
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

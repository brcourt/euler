package main
import ("fmt"
        "time"
        "log")

/* The sum of the squares of the first ten natural numbers is 385
The square of the sum of the first ten natural numbers is 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum. */

func main() {
  defer TimeTaken(time.Now(), "main()")

  a,b := 0,0
  for i :=0; i < 101; i++ {
    a += (i*i)
    b += i
  }
  fmt.Println((b*b)-a)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

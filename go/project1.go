package main
import ("fmt"
        "time"
        "log")

/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000. */

func main(){
  defer TimeTaken(time.Now(), "main()")

  value := 0
  for i := 0; i < 1000; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      value += i
    }
  }
  fmt.Println(value)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

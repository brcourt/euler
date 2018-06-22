package main
import ("fmt"
        "math/big"
        "time"
        "log")


func main() {
  defer TimeTaken(time.Now(), "main()")

  i := big.NewInt(2)
  isPrime := i.ProbablyPrime(1)
  fmt.Println(isPrime)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

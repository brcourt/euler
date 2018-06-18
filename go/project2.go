package main
import ("fmt"
        "time"
        "log")

/* Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms. */


func main(){
  defer TimeTaken(time.Now(), "main()")

  fibo := []int{1, 2}
  value := 0
  // Fn = Fn-1 + Fn-2

  for latest := fibo[len(fibo)-1]; latest < 3999999; latest = fibo[len(fibo)-1]{
    fibo = append(fibo, (fibo)[len(fibo)-1] + (fibo)[len(fibo)-2])
    if latest % 2 == 0 {
      value = value + latest
    }
  }
  fmt.Println(value)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

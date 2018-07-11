package main
import ("fmt"
        "strconv"
        "time"
        "log")

/* A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers. */

func is_palindrome(i int) bool{
  string_num := strconv.Itoa(i)
  for i, v := 0, len(string_num)-1; i < v ; i, v = i+1, v-1 {
    if string_num[v] != string_num[i] {
      return false
    }
  }
  return true
}

func main() {
  defer TimeTaken(time.Now(), "main()")

  a,b := 999, 999
  max := 0

  for ; b > 100 && a > 100; b--{
    test := a*b
    if test < max || b == 101  {
      b = 999
      a--
    } else if is_palindrome(test) == true {
      if test > max {
        max = test
      }
    }
    // stop at 500, since we we want to avoid duplicated test cases (a*b == b*a)
    if a < 500 {
      break
    }
  }
  fmt.Println(max)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

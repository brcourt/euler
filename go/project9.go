package main
import ("fmt"
        "time"
        "log")

/* A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc. */

func main() {
  defer TimeTaken(time.Now(), "main()")

  sum := 1000

  calculate:
    // a will never be more than a third of all numbers combined
    for a := 1; a <= sum/3; a++ {
      // b will mever be more than a half of all numbers combined
      for b := a + 1; b <= sum/2; b++ {
        c := sum - b - a

        if a*a + b*b == c*c {
          fmt.Println(a*b*c)
          break calculate
        }
      }
    }

}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

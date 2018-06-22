package main

import ("fmt"
        "time"
        "log")

func main() {
  defer TimeTaken(time.Now(), "main()")

	var a, b, sum, temp int = 1, 2, 0, 0
	for {
		if a%2 == 0 {
			sum += a
		}
		temp = b
		b = a + b
		a = temp
		if a >= 4000000 {
			break
		}
	}
	fmt.Println(sum)
}

func TimeTaken(t time.Time, name string) {
    elapsed := time.Since(t)
    log.Printf("TIME: %s took %s\n", name, elapsed)
}

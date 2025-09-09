package main

const size = 10 // Change to 1024 to observe stack growth differences

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{}) // Initial call
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 { // Recursion depth limit
		return
	}

	stackCopy(s, c, a)
}

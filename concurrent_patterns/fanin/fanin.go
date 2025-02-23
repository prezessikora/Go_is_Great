package fanin

// Multiplexer or fanin function that joins two channels into returned one
func FanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

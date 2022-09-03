package jokes

var (
	Jokes        = make([]string, 0)
	JokesChan    = make(chan string, 0)
	JokesStarted = false
)

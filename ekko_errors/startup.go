package ekko_errors

type Startup struct {
	message string
}

func (s *Startup) Error() string {
	return s.message
}

var startup = &Startup{}

func WithMsg(message string) *Startup {
	startup.message = message
	return startup
}

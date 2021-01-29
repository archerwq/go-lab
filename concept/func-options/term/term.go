package term

// Term Terminal
type Term struct {
	dev   string
	raw   bool
	speed int
}

// RawMode places the terminal into raw mode.
func RawMode(t *Term) error {
	return t.setRawMode()
}

// Speed set the baud rate option for the terminal.
func Speed(speed int) func(*Term) error {
	return func(t *Term) error {
		return t.setSpeed(speed)
	}
}

// Open opens an asynchronous communication port.
func Open(dev string, options ...func(*Term) error) (*Term, error) {
	term, err := openTerm(dev)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		if err := opt(term); err != nil {
			return nil, err
		}
	}

	return term, nil
}

func openTerm(dev string) (*Term, error) {
	// TODO
	return &Term{dev: dev}, nil
}

func (t *Term) setRawMode() error {
	t.raw = true
	// TODO
	return nil
}

func (t *Term) setSpeed(speed int) error {
	t.speed = speed
	// TODO
	return nil
}

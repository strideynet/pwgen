package pwgen

import "log"

type settings struct {
	length int
	uppercase int
	lowercase int
	number int
	special int
}

func (s settings) processOptions(opts ...Option) settings  {
	for _, opt := range opts {
		s = opt(s)
	}

	return s
}

var defaultSettings = settings{
	length: 8,
	uppercase: 2,
	lowercase: 2,
	number: 2,
	special: 2,
}

type Option func(settings) settings

func Generate(opts ...Option) (string, error) {
	s := defaultSettings.processOptions(opts...)

	log.Printf("%d", s.lowercase)

	return "", nil
}
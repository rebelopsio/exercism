package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct{}

type Portuguese struct{}

func SayHello(s string, g Greeter) string {
	ln := g.LanguageName()
	gg := g.Greet(s)
	return fmt.Sprintf("I can speak %s: %s", ln, gg)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

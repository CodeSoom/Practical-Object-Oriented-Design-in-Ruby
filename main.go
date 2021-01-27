package main

type Toy interface {
	Play() string
}

type Duck interface {
	Quack() string
}

type RuberDuck struct {
}

func (d RuberDuck) Quack() string {
	return "Bbick!"
}

func (d RuberDuck) Play() string {
	return "It's so funny"
}

func main() {
	duck := RuberDuck{}
	putInCage(duck);
	play(duck);
}

func putInCage(duck Duck) {
	println(duck.Quack());
}

func play(toy Toy) {
	println(toy.Play())
}

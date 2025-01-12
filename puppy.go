package puppy

import (
	"github.com/hazembelkhiria/dogs"
)

func Bark() string {
	return "woof!!"
}

func Barks() string {
	return "woof!! woof!! woof!!"
}

func BigBark() string {
	return dogs.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dogs.WhenGrownUp(Barks())
}

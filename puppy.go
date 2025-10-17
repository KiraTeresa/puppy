package puppy

import (
	"fmt"
	"github.com/KiraTeresa/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func versionTest() {
	fmt.Println("Change of V1.0.0")
}

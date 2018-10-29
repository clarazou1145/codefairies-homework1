package song

import (
	"fmt"
	"testing"
)

func TestGroupBy(t *testing.T) {
	input := []animal{
		animal{"fly", ""},
		animal{"spider", "That wriggled and wiggled and tickled inside her."},
		animal{"bird", "How absurd to swallow a bird."},
		animal{"cat", "Fancy that to swallow a cat!"},
		animal{"dog", "What a hog, to swallow a dog!"},
		animal{"cow", "I don't know how she swallowed a cow!"},
		animal{"horse", "...She's dead, of course!"},
	}
	song := tell(input)
	fmt.Println(song)
}

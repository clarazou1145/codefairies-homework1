package song

import "fmt"

var i int

type animal struct {
	Name     string
	Describe string
}

var animals []animal

func tell(animals []animal) string {
	var result string
	var s string
	for k, v := range animals {
		if k != 0 && k < len(animals)-1 {
			s = fmt.Sprintf("There was an old lady who swallowed a %s;\n%s\n", v.Name, v.Describe)
			result += s
			for i := k; i > 0; i-- {
				if i > 1 {
					s = fmt.Sprintf("She swallowed the %s to catch the %s;\n", animals[i].Name, animals[i-1].Name)
					result += s
				} else {
					s = fmt.Sprintf("She swallowed the %s to catch the %s;\n", animals[i].Name, animals[0].Name)
					result += s
				}
			}
			s = "I don't know why she swallowed a fly - perhaps she'll die!\n\n"
			result += s
		}
	}
	var end string
	end = fmt.Sprintf("There was an old lady who swallowed a %s ...\n%s\n", animals[len(animals)-1].Name, animals[len(animals)-1].Describe)
	result += end
	return result
}

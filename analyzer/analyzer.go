package analyzer

import "strings"

type Count struct {
	Word  string
	Count int64
}

func Find(cs []*Count, word string) (int, bool) {
	for i, c := range cs {
		if c.Word == word {
			return i, true
		}
	}
	return -1, false
}

func Analyze(text string) []*Count {
	text = strip(text)
	words := strings.Split(text, " ")

	var cs []*Count
	for _, w := range words {
		if i, ok := Find(cs, w); ok {
			cs[i].Count++
		} else {
			c := &Count{Word: w, Count: 1}
			cs = append(cs, c)
		}
	}
	return cs
}

func strip(text string) string {
	r := strings.NewReplacer(";", "", ",", "", "`", "", "'", "", "(", "", ")", "", ".", "", "\n", " ", "\r", " ")
	return r.Replace(text)
}

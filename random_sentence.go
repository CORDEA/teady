package main

import (
	"math/rand"
)

type RandomSentence struct {
	BaseSentences []string
	Count         int
}

func NewRandomSentence(sentences []string, count int) *RandomSentence {
	return &RandomSentence{BaseSentences: sentences, Count: count}
}

func (m *RandomSentence) Generate() string {
	if len(m.BaseSentences) > 0 {
		return m.genFromBase()
	}
	return m.gen()
}

func (m *RandomSentence) genFromBase() string {
	return m.BaseSentences[rand.Intn(len(m.BaseSentences)-1)]
}

func (m *RandomSentence) gen() string {
	result := ""
	for i := 0; i < m.Count; i++ {
		r := rand.Intn(93) + 33
		result = result + string(rune(r))
	}
	return result
}

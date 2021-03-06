package main

import (
	"fmt"
)

// 68. Text Justification Hard
// Add to List

// Share
// Given an array of words and a width maxWidth, format the text such that
// each line has exactly maxWidth characters and is fully (left and right) justified.

// You should pack your words in a greedy approach; that is,
// pack as many words as you can in each line. Pad extra spaces ' ' when necessary
// so that each line has exactly maxWidth characters.

// Extra spaces between words should be distributed as evenly as possible.
// If the number of spaces on a line do not divide evenly between words,
// the empty slots on the left will be assigned more spaces than the slots on the right.

// For the last line of text, it should be left justified and no extra space is inserted between words.

// Note:
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.

type wordGroup struct {
	groups   []string
	totalLen int
	last     bool
}

func covertEachArray(wg *wordGroup, maxWidth int) string {
	var result string

	if wg.last == false {
		var spaces = make([]string, len(wg.groups)-1)
		for i := 0; i < (maxWidth - wg.totalLen); i++ {
			spaces[i%len(spaces)] = spaces[i%len(spaces)] + " "
		}
		for i := 0; i < len(wg.groups)-1; i++ {
			result = result + wg.groups[i] + spaces[i]
		}
		result = result + wg.groups[len(wg.groups)-1]
	}
	if wg.last == true {
		space := ""
		for i := 0; i < len(wg.groups)-1; i++ {
			result = result + wg.groups[i] + " "
		}
		result = result + wg.groups[len(wg.groups)-1]
		for i := 0; i < (maxWidth - len(result)); i++ {
			space = space + " "
		}
		result = result + space
	}

	fmt.Println(len(result))
	return result
}

func fullJustify(words []string, maxWidth int) []string {
	wordGroups := []*wordGroup{}
	wg := &wordGroup{last: false}
	result := []string{}
	for i, word := range words {
		if len(wg.groups) == 0 {
			wg.groups = append(wg.groups, word)
			wg.totalLen += len(word)
		} else if len(wg.groups) == 1 && len(word) <= maxWidth-wg.totalLen-1 {
			wg.groups = append(wg.groups, word)
			wg.totalLen += len(word)
		} else if len(wg.groups) == 1 && len(word) > maxWidth-wg.totalLen-1 {
			wg.last = true
			wordGroups = append(wordGroups, wg)
			wg = &wordGroup{groups: []string{word}, totalLen: len(word), last: false}
		} else if len(word) < maxWidth-wg.totalLen-len(wg.groups)+1 {
			wg.groups = append(wg.groups, word)
			wg.totalLen += len(word)
		} else {
			wordGroups = append(wordGroups, wg)
			wg = &wordGroup{groups: []string{word}, totalLen: len(word), last: false}
		}
		if i == len(words)-1 {
			wg.last = true
			fmt.Println(wg)
			wordGroups = append(wordGroups, wg)
		}
	}
	for _, wg := range wordGroups {
		eachResult := covertEachArray(wg, maxWidth)
		result = append(result, eachResult)
	}
	return result
}

func main() {
	words := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth := 16
	fmt.Println(fullJustify(words, maxWidth))
}

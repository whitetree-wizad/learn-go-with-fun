package main

import (
	"fmt"
	"strings"
)

func sepereat(text string) []string {
	sepereatText := strings.Fields(strings.ToLower(text))
	for i, q := range sepereatText {
		sepereatText[i] = strings.Trim(q, ",.;")
	}
	return sepereatText
}

func main() {
	words := `As far as eye could reach he saw nothing but the stems of the great plants about him
receding in the violet shade, and far overhead the multiple transparency of huge leaves
filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
he felt able he ran again; the ground continued soft and springy, covered with the same
resilient weed which was the first thing his hands had touched in Malacandra. Once or
twice a small red creature scuttled across his path, but otherwise there seemed to be no
life stirring in the wood; nothing to fear-except the fact of wandering unprovisioned
and alone in a forest of unknown vegetation thousands or millions of miles beyond the
reach or knowledge of man.`

	wordsSlice := sepereat(words)
	countMap := make(map[string]int)

	for _, q := range wordsSlice {
		countMap[q]++
	}

	fmt.Println(countMap)

}

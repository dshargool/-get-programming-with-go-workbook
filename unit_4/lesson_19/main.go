package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

func word_freq(p_string string) map[string]int {
	p_string = strings.ToLower(p_string)
	var all_string []string = strings.Fields(p_string)
	string_map := make(map[string]int)
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	for _, word := range all_string {
		word = reg.ReplaceAllString(word, "")
		string_map[word]++
	}
	return string_map
}

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temperature)

	planets := map[string]string{
		"Earth": "Earthling",
		"Mars":  "Martian",
	}

	planets_2 := planets           // These are both the same data
	planets_2["Mars"] = "Marstian" // Changes both planets_2 and planets
	fmt.Println(planets)
	fmt.Println(planets_2)

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}

	groups := make(map[float64][]float64) // A map with float64 keys and []float64 values

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	fmt.Println(groups)
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	unique := make([]float64, 0, len(set))

	for t := range set {
		unique = append(unique, t)
	}

	sort.Float64s(unique)
	fmt.Println(unique)

	passage := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	words := word_freq(passage)
	for word, count := range words {
		if count > 1 {
			fmt.Println(word, count)
		}
	}
}

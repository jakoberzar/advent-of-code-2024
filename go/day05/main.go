package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jakoberzar/advent-of-code-2024/internal/sets"
	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-05/simple.txt"
const FULL_INPUT = "./../../inputs/day-05/full.txt"

type Rule struct {
	before, after int
}
type Update = []int
type BeforeToAfterMap = map[int]sets.IntSet

func main() {
	rules, updates := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(rules, updates))
	fmt.Printf("Result of star 2 is %d\n", star2(rules, updates))
}

func star1(rules []Rule, updates []Update) (sum int) {
	baMap := makeBeforeAfterMap(rules)
	for _, update := range updates {
		if isUpdateCorrect(update, baMap) {
			sum += middleElement(update)
		}
	}
	return
}

func star2(rules []Rule, updates []Update) (sum int) {
	baMap := makeBeforeAfterMap(rules)
	for _, update := range updates {
		if !isUpdateCorrect(update, baMap) {
			update = fixUpdate(update, baMap)
			sum += middleElement(update)
		}
	}
	return
}

func makeBeforeAfterMap(rules []Rule) BeforeToAfterMap {
	baMap := make(BeforeToAfterMap)
	for _, rule := range rules {
		baMap[rule.before] = sets.Insert(rule.after, baMap[rule.before])
	}
	return baMap
}

func isUpdateCorrect(update Update, baMap BeforeToAfterMap) bool {
	notAllowed := make(map[int]bool)
	for i := len(update) - 1; i >= 0; i-- {
		el := update[i]
		if notAllowed[el] {
			return false
		}
		for i := range baMap[el] {
			notAllowed[i] = true
		}
	}
	return true
}

func middleElement(update Update) int {
	idx := len(update) / 2
	return update[idx]
}

func makeLocalBeforeAfterMap(update Update, baMap BeforeToAfterMap) BeforeToAfterMap {
	depMap := make(BeforeToAfterMap)
	for i := 0; i < len(update); i++ {
		el := update[i]
		depMap[el] = make(sets.IntSet)
		for j := 0; j < len(update); j++ {
			if _, in := baMap[el][update[j]]; in {
				depMap[el][update[j]] = true
			}
		}
	}
	return depMap
}

func findMapElementWithEmptySet(baMap BeforeToAfterMap) (int, bool) {
	for el, set := range baMap {
		if len(set) == 0 {
			return el, true
		}
	}
	return 0, false
}

// Removes element from all the sets in the map.
// Also returns a hint about the next empty set.
func removeElementFromBaMap(el int, baMap BeforeToAfterMap) (m BeforeToAfterMap, hint int) {
	delete(baMap, el)
	for key, set := range baMap {
		_, in := set[el]
		if in {
			delete(set, el)
			baMap[key] = set
			if len(set) == 0 {
				hint = key
			}
		}
	}
	return baMap, hint
}

func fixUpdate(update Update, baMap BeforeToAfterMap) Update {
	newUpdate := make([]int, 0, len(update))
	localMap := makeLocalBeforeAfterMap(update, baMap)
	el, _ := findMapElementWithEmptySet(localMap)
	for len(localMap) > 0 {
		newUpdate = append(newUpdate, el)
		localMap, el = removeElementFromBaMap(el, localMap)
	}
	slices.Reverse(newUpdate)
	return newUpdate
}

func readInput(path string) (rules []Rule, updates []Update) {
	lines := utils.ReadLinesOrDie(path)
	readingRules := true
	for _, line := range lines {
		if line == "" {
			readingRules = false
		}
		if readingRules {
			strNums := strings.Split(line, "|")
			before, _ := strconv.Atoi(strNums[0])
			after, _ := strconv.Atoi(strNums[1])
			rules = append(rules, Rule{before, after})
		} else {
			update := make([]int, 0)
			for _, s := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(s)
				update = append(update, n)
			}
			updates = append(updates, update)
		}
	}
	return
}

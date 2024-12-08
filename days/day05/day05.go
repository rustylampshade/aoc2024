package day05

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type PageOrderingRule []string
type UpdatePages []string

func Solve() {
	dat, err := os.ReadFile("inputs/day05.txt")
	if err != nil {
		panic(err)
	}

	rules := []PageOrderingRule{}
	updates := []UpdatePages{}

	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "|") {
			rules = append(rules, strings.Split(line, "|"))
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	sumMiddlePages := 0
	sumCorrectedPages := 0
	for _, update := range updates {
		if allRulesValid(&rules, &update) {
			sumMiddlePages += middlePageValue(&update)
		} else {
			correctedUpdate := UpdatePages{processRules(rules, update)}
			sumCorrectedPages += middlePageValue(&correctedUpdate)
		}
	}
	fmt.Println(sumMiddlePages)
	fmt.Println(sumCorrectedPages)
}

func ruleIsApplicable(rule PageOrderingRule, update *UpdatePages) bool {
	foundLower := false
	foundHigher := false
	for _, page := range *update {
		if page == rule[0] {
			foundLower = true
		}
		if page == rule[1] {
			foundHigher = true
		}
	}
	return foundLower && foundHigher
}

func ruleIsValid(rule PageOrderingRule, update *UpdatePages) bool {
	foundHigher := false
	for _, page := range *update {
		if page == rule[0] {
			return !foundHigher
		}
		if page == rule[1] {
			foundHigher = true
		}
	}
	return true
}

func allRulesValid(rules *[]PageOrderingRule, update *UpdatePages) bool {
	for _, rule := range *rules {
		if !ruleIsValid(rule, update) {
			return false
		}
	}
	return true
}

func middlePageValue(update *UpdatePages) int {
	pagelist := *update
	if len(pagelist) == 0 {
		return 0
	}
	middlePage := pagelist[len(pagelist)/2]
	middlePageInt, err := strconv.Atoi(middlePage)
	if err != nil {
		panic(err)
	}
	return middlePageInt
}

func processRules(rules []PageOrderingRule, update UpdatePages) string {
	mutableRules := make([]PageOrderingRule, len(rules))
	mutableUpdate := make(UpdatePages, len(update))
	copy(mutableRules, rules)
	copy(mutableUpdate, update)
	for i := 0; i < len(update)/2; i++ {

		// Find all the rules that define a strict ordering between two elements, both of which exist in this update
		tightlyRelevantRules := []PageOrderingRule{}
		for _, rule := range mutableRules {
			if ruleIsApplicable(rule, &mutableUpdate) {
				tightlyRelevantRules = append(tightlyRelevantRules, rule)
			}
		}
		mutableRules = make([]PageOrderingRule, len(tightlyRelevantRules))
		copy(mutableRules, tightlyRelevantRules)

		// Make a map out of the lower and higher values of the tight rules for easier checking
		lowerRules := make(map[string]int)
		higherRules := make(map[string]int)
		for _, rule := range tightlyRelevantRules {
			if _, exists := lowerRules[rule[0]]; !exists {
				lowerRules[rule[0]] = 1
			}
			if _, exists := higherRules[rule[1]]; !exists {
				higherRules[rule[1]] = 1
			}
		}

		// Identify the single value that is only mentioned in the lower map. That is the element that MUST come first
		// in the rest of the order. Similarly find the only one mentioned in the higher map.
		var lowest, highest string
		for lower := range lowerRules {
			if _, exists := higherRules[lower]; !exists {
				lowest = lower
			}
		}
		for higher := range higherRules {
			if _, exists := lowerRules[higher]; !exists {
				highest = higher
			}
		}

		// Turns out we don't even care about what the values ARE. We're just working into the middle. Delete all the rules
		// about the highest and lowest outliers and repeat.
		var ruleIdxToRemove = []int{}
		var updateIdxToRemove = []int{}
		for i, rule := range mutableRules {
			if rule[0] == lowest {
				ruleIdxToRemove = append(ruleIdxToRemove, i)
				continue
			}
			if rule[1] == highest {
				ruleIdxToRemove = append(ruleIdxToRemove, i)
			}
		}
		for i, page := range mutableUpdate {
			if page == lowest {
				updateIdxToRemove = append(updateIdxToRemove, i)
				continue
			}
			if page == highest {
				updateIdxToRemove = append(updateIdxToRemove, i)
			}
		}
		sort.Ints(ruleIdxToRemove)
		sort.Ints(updateIdxToRemove)
		for i := len(ruleIdxToRemove) - 1; i >= 0; i-- {
			idx := ruleIdxToRemove[i]
			mutableRules = slices.Delete(mutableRules, idx, idx+1)
		}

		mutableUpdate = removetwo(mutableUpdate, updateIdxToRemove[0], updateIdxToRemove[1])
	}
	return mutableUpdate[0]
}

func removetwo(s []string, l, h int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:l]...)
	ret = append(ret, s[l+1:h]...)
	return append(ret, s[h+1:]...)
}

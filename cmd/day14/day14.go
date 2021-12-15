package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

var log = logger.NewScoped("day14")

func main() {
	common.Init()

	inputLines := common.ReadInputLines()
	if len(inputLines) < 3 {
		log.Error().WithInt("lines", len(inputLines)).
			Message("Expected at least 3 lines of input.")
		os.Exit(1)
	}

	template := inputLines[0]
	rules := map[RuleKey]rune{}
	for _, line := range inputLines[2:] {
		rule, err := ParseRule(line)
		if err != nil {
			log.Error().WithError(err).
				WithString("line", line).
				Message("Failed to parse rule.")
			os.Exit(1)
		}
		rules[rule.RuleKey] = rule.Insertion
	}

	log.Info().WithString("template", template).
		WithInt("rules", len(rules)).
		Message("Scanning complete.")

	steps := 10
	if common.Part2 {
		steps = 40
	}
	for i := 1; i <= steps; i++ {
		template = ApplyRules(template, rules)
		if len(template) > 90 {
			log.Debug().
				WithInt("step", i).
				WithInt("outputLen", len(template)).
				Message("Applied rules.")
		} else {
			log.Debug().
				WithInt("step", i).
				WithString("output", template).
				Message("Applied rules.")
		}
	}

	leastCommon, mostCommon := CountMinMaxRunes(template)
	log.Debug().
		WithStringf("leastCommon", "%c:%d", leastCommon.rune, leastCommon.count).
		WithStringf("mostCommon", "%c:%d", mostCommon.rune, mostCommon.count).
		Message("Found min/max runes.")

	result := mostCommon.count - leastCommon.count
	log.Info().WithInt("result", result).
		Message("Calculated (most - least).")
}

type RuneCount struct {
	rune
	count int
}

func CountMinMaxRunes(s string) (RuneCount, RuneCount) {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	var leastCommon, mostCommon RuneCount
	for r, count := range m {
		if leastCommon.count == 0 || count < leastCommon.count {
			leastCommon = RuneCount{r, count}
		}
		if mostCommon.count == 0 || count > mostCommon.count {
			mostCommon = RuneCount{r, count}
		}
	}
	return leastCommon, mostCommon
}

func ApplyRules(template string, rules map[RuleKey]rune) string {
	var sb strings.Builder
	sb.Grow(len(template) * 2)
	prev, i := utf8.DecodeRuneInString(template)
	sb.WriteRune(prev)
	for _, r := range template[i:] {
		if insert, ok := rules[RuleKey{prev, r}]; ok {
			sb.WriteRune(insert)
		}
		sb.WriteRune(r)
		prev = r
	}
	return sb.String()
}

type RuleKey struct {
	Left  rune
	Right rune
}

type Rule struct {
	RuleKey
	Insertion rune
}

func ParseRule(s string) (Rule, error) {
	var rule Rule
	_, err := fmt.Sscanf(s, "%c%c -> %c", &rule.Left, &rule.Right, &rule.Insertion)
	return rule, err
}

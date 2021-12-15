package main

import (
	"fmt"
	"os"

	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/jilleJr/adventofcode-2021-go/internal/common"
)

// There's an off-by-1 error somewhere here. I haven't found it.

var log = logger.NewScoped("day14")

func main() {
	common.Init()

	inputLines := common.ReadInputLines()
	if len(inputLines) < 3 {
		log.Error().WithInt("lines", len(inputLines)).
			Message("Expected at least 3 lines of input.")
		os.Exit(1)
	}

	templateString := inputLines[0]
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

	log.Info().WithString("template", templateString).
		WithInt("rules", len(rules)).
		Message("Scanning complete.")

	steps := 10
	if common.Part2 {
		steps = 40
	}
	templatePairs := countTemplatePairs(templateString)
	for i := 1; i <= steps; i++ {
		templatePairs = ApplyRules(templatePairs, rules)
	}

	leastCommon, mostCommon := CountMinMaxRunes(templatePairs)
	log.Debug().
		WithStringf("leastCommon", "%c:%d", leastCommon.rune, leastCommon.count).
		WithStringf("mostCommon", "%c:%d", mostCommon.rune, mostCommon.count).
		WithInt("steps", steps).
		Message("Found min/max runes.")

	result := mostCommon.count - leastCommon.count
	log.Info().WithInt("result", result).
		Message("Calculated (most - least).")
}

type RuneCount struct {
	rune
	count int
}

func CountMinMaxRunes(pairs map[RuleKey]int) (RuneCount, RuneCount) {
	m := map[rune]int{}
	for p, count := range pairs {
		log.Debug().
			WithStringf("left", "%c", p.Left).
			WithStringf("right", "%c", p.Right).
			WithInt("count", count).
			Message("")
		m[p.Left] += count
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

func countTemplatePairs(template string) map[RuleKey]int {
	result := map[RuleKey]int{}
	var prev rune
	for i, r := range template {
		if i > 0 {
			result[RuleKey{prev, r}]++
		}
		prev = r
	}
	return result
}

func ApplyRules(template map[RuleKey]int, rules map[RuleKey]rune) map[RuleKey]int {
	result := map[RuleKey]int{}
	for pair, count := range template {
		if insert, ok := rules[pair]; ok {
			result[RuleKey{pair.Left, insert}] += count
			result[RuleKey{insert, pair.Right}] += count
		} else {
			result[pair]++
		}
	}
	return result
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

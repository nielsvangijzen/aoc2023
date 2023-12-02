package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2023/util"
	"log"
	"strconv"
	"unicode"
)

func partOne() error {
	input := util.MustInputLines("day-01/input.txt")

	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		var digits []string
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
		}

		// Dirty hack for strings with only one digit
		if len(digits) == 1 {
			digits = append(digits, digits[0])
		}

		converted, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		sum += converted
	}

	log.Printf("D01P01: %d", sum)
	return nil
}

type Parser struct {
	EOL    bool
	line   string
	pos    int
	digits []string
}

func NewParser(line string) *Parser {
	return &Parser{
		EOL:    false,
		line:   line,
		pos:    -1,
		digits: []string{},
	}
}

func (parser *Parser) Next() bool {
	if parser.pos+1 >= len(parser.line) {
		return false
	}

	parser.pos += 1

	return true
}

func (parser *Parser) Current() rune {
	return rune(parser.line[parser.pos])
}

func (parser *Parser) Parse() ([]string, error) {
	if len(parser.line) == 0 {
		return nil, fmt.Errorf("empty string")
	}

	for parser.Next() {
		current := parser.Current()
		if unicode.IsDigit(current) {
			parser.digits = append(parser.digits, string(current))
		} else {
			parser.TryParseWord()
		}
	}

	return parser.digits, nil
}

var wordIntMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func (parser *Parser) TryParseWord() {
	for word, digit := range wordIntMap {
		take := len(word)
		if parser.pos+take > len(parser.line) {
			continue
		}

		wordToCheck := parser.line[parser.pos : parser.pos+take]
		if wordToCheck == word {
			parser.digits = append(parser.digits, digit)
		}
	}
}

func partTwo() error {
	input := util.MustInputLines("day-01/input.txt")

	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		p := NewParser(line)
		digits, err := p.Parse()
		if err != nil {
			return err
		}

		// Dirty hack for strings with only one digit
		if len(digits) == 1 {
			digits = append(digits, digits[0])
		}

		converted, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		sum += converted
	}

	log.Printf("D01P02: %d", sum)
	return nil
}

func main() {
	err := partOne()
	if err != nil {
		log.Printf("D01P01: error, %s", err)
	}

	err = partTwo()
	if err != nil {
		log.Printf("D01P02: error, %s", err)
	}
}

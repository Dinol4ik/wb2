package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := "a4bc2d5e"
	fmt.Println(unpack(str))
}
func unpack(input string) (string, error) {
	runes := []rune(input)
	var b strings.Builder

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) && runes[i] < unicode.MaxASCII {
			if i == 0 {
				return "", nil
			}

			var num strings.Builder
			num.WriteRune(runes[i])
			letter := runes[i-1]

			for j := i + 1; j < len(runes)-1 && unicode.IsDigit(runes[j]) && runes[j] < unicode.MaxASCII; j++ {
				num.WriteRune(runes[j])
				i++
			}

			res, err := strconv.Atoi(num.String())
			if errors.Is(err, strconv.ErrRange) {
				// if number out of range, just print it
				b.WriteString(num.String())
				continue
			}
			if err != nil {
				return "", err
			}

			for j := 0; j < res-1; j++ {
				b.WriteRune(letter)
			}

			continue
		}
		_, err := b.WriteRune(runes[i])
		if err != nil {
			return "", err
		}
	}

	return b.String(), nil
}

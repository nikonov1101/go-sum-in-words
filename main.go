package main

import (
	"fmt"
	"math"
	"strings"
)

type numberString struct {
	name string
	rubl string
	kop  string
}

var number = map[int64]numberString{
	1:  {name: "один", rubl: "рубль", kop: "копейка"},
	2:  {name: "два", rubl: "рубля", kop: "копейки"},
	3:  {name: "три", rubl: "рубля", kop: "копейки"},
	4:  {name: "четыре", rubl: "рубля", kop: "копейки"},
	5:  {name: "пять", rubl: "рублей", kop: "копеек"},
	6:  {name: "шесть", rubl: "рублей", kop: "копеек"},
	7:  {name: "семь", rubl: "рублей", kop: "копеек"},
	8:  {name: "восемь", rubl: "рублей", kop: "копеек"},
	9:  {name: "девять", rubl: "рублей", kop: "копеек"},
	0:  {name: "", rubl: "рублей", kop: "копеек"},
	10: {name: "десять", rubl: "рублей", kop: "копеек"},
	11: {name: "одиндцать", rubl: "рублей", kop: "копеек"},
	12: {name: "двенадцать", rubl: "рублей", kop: "копеек"},
	13: {name: "тринадцать", rubl: "рублей", kop: "копеек"},
	14: {name: "четырнадцать", rubl: "рублей", kop: "копеек"},
	15: {name: "пятнадцать", rubl: "рублей", kop: "копеек"},
	16: {name: "шестнадцать", rubl: "рублей", kop: "копеек"},
	17: {name: "семнадцать", rubl: "рублей", kop: "копеек"},
	18: {name: "восемнадцать", rubl: "рублей", kop: "копеек"},
	19: {name: "девятнадцать", rubl: "рублей", kop: "копеек"},
}
var ten = map[int64]string{
	1: "десять",
	2: "двадцать",
	3: "тридцать",
	4: "сорок",
	5: "пятьдесят",
	6: "шестьдесят",
	7: "семьдесят",
	8: "восемьдесят",
	9: "девяносто",
}
var hundred = map[int64]string{
	1: "сто",
	2: "двести",
	3: "триста",
	4: "четыреста",
	5: "пятьсот",
	6: "шестьсот",
	7: "семьсот",
	8: "восемьсот",
	9: "девятьсот",
}

func formatInteger(v int64) string {
	sum := ""
	pos := 1

	for v > 0 {
		digit := v % 10
		switch pos {
		//обрабатываем каждый первый разряд в тройке
		case 1, 4, 7:
			if v%100/10 == 1 {
				digit = v % 100
			}

			//добавляем "тысячи", если необходимо
			if pos == 4 && v%1000 > 0 {
				switch digit {
				case 1:
					sum = " тысяча" + sum
				case 2, 3, 4:
					sum = " тысячи" + sum
				default:
					sum = " тысяч" + sum
				}
			}
			//добавляем "миллионы", если необходимо
			if pos == 7 {
				switch digit {
				case 1:
					sum = " миллион" + sum
				case 2, 3, 4:
					sum = " миллиона" + sum
				default:
					sum = " миллионов" + sum
				}
			}

			// добавляем "рублей" в требуемой форме
			if sum == "" {
				sum = " " + number[digit].rubl
			}

			//корректировка для тысяч
			if pos == 4 && digit == 1 {
				sum = "одна" + sum
			} else {
				if pos == 4 && digit == 2 {
					sum = "две" + sum
				} else {
					if digit != 0 {
						sum = " " + number[digit].name + sum
					}
				}
			}

			pos++
			if v%100/10 == 1 {
				pos++
				v = v / 10
			}
		case 2, 5, 8:
			if digit != 0 {
				sum = " " + ten[digit] + sum
			}
			pos++
		case 3, 6, 9:
			if digit != 0 {
				sum = " " + hundred[digit] + sum
			}
			pos++
		}
		v = v / 10

	}

	return capFirst(strings.TrimSpace(sum))
}

func formatFraction(v int64) string {
	var sum string
	switch {
	case v >= 3 && v <= 19:
		sum = sum + " " + number[v].name + " " + number[v].kop
	case v%10 == 0:
		sum = sum + " " + ten[v/10] + " " + number[0].kop
	case v%10 == 1:
		sum = sum + ten[v/10] + " одна " + number[1].kop
	case v%10 == 2:
		sum = sum + ten[v/10] + " две " + number[2].kop
	default:
		sum = sum + " " + ten[v/10] + " " + number[v%10].name + " " + number[v%10].kop
	}

	return capFirst(strings.TrimSpace(sum))
}

// FormatSum convert float value presenring cost of smth into string value
// example:  "Восемьсот пятьдесят рублей", "двадцать пять копеек"
func FormatSum(v float64, withFraction bool) (string, string) {
	var integerValue, fracValue string

	// separate Rouble
	intRouble, frac := math.Modf(v)
	iRubl := int64(intRouble)
	integerValue = formatInteger(iRubl)

	if withFraction {
		// separate Fraction
		fFrac := frac * 100
		_, frac = math.Modf(fFrac)
		if frac >= 0.5 {
			fFrac = math.Ceil(fFrac)
		} else {
			frac = math.Floor(fFrac)
		}
		iFrac := int64(fFrac)
		fracValue = formatFraction(iFrac)
	}

	return integerValue, fracValue
}

// capitalize First letter for unicode string
func capFirst(s string) string {
	if len(s) == 1 {
		return strings.ToUpper(s)
	}

	ltr := strings.ToUpper(string([]rune(s)[0]))
	rst := []rune(s)[1:]

	return ltr + strings.ToLower(string(rst))
}

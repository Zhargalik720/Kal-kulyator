package main

import "fmt"

func main() {
	for {
		var a, b, c, d string
		fmt.Scanln(&a, &c, &b, &d)
		if d != "" {
			fmt.Println("Вывод ошибки, так как Математическая операция не удовлетворяет заданию - два операнда и один оператор(+,-,/,*).")
			break
		} else if b == "" || c == "" {
			fmt.Println("Вывод ошибкиб так как строка не является математической операцией.")
			break
		} else if c != "+" && c != "-" && c != "*" && c != "/" {
			fmt.Println("Вывод ошибкиб так как строка не является математической операцией.")
			break
		} else if (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9" || a == "10") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX" || b == "X") {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			break
		} else if (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX" || a == "X") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" || b == "10") {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			break
		} else if (a == "1" || a == "2" || a == "3" || a == "4" || a == "5" || a == "6" || a == "7" || a == "8" || a == "9" || a == "10") && (b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" || b == "10") {
			a1 := odin(a)
			b1 := odin(b)
			chetyre(a1, b1, c)
		} else if (a == "I" || a == "II" || a == "III" || a == "IV" || a == "V" || a == "VI" || a == "VII" || a == "VIII" || a == "IX" || a == "X") && (b == "I" || b == "II" || b == "III" || b == "IV" || b == "V" || b == "VI" || b == "VII" || b == "VIII" || b == "IX" || b == "X") {
			a1 := dva(a)
			b1 := dva(b)
			d1 := pyat(a1, b1, c)
			if d1 < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе счисления нет отрицательных чисел.")
				break
			} else {
				tri(d1)
			}
		} else {
			fmt.Println("Ошибка")
			break
		}

	}
}
func odin(a string) int { //пеоевод строки в число
	var d int
	switch a {
	case "1":
		d = 1
	case "2":
		d = 2
	case "3":
		d = 3
	case "4":
		d = 4
	case "5":
		d = 5
	case "6":
		d = 6
	case "7":
		d = 7
	case "8":
		d = 8
	case "9":
		d = 9
	case "10":
		d = 10
	}
	return d
}
func dva(a string) int { //перевод римскив в арабские
	var d int
	switch a {
	case "I":
		d = 1
	case "II":
		d = 2
	case "III":
		d = 3
	case "IV":
		d = 4
	case "V":
		d = 5
	case "VI":
		d = 6
	case "VII":
		d = 7
	case "VIII":
		d = 8
	case "IX":
		d = 9
	case "X":
		d = 10
	}
	return d
}
func tri(a int) { //перевод арабских в римские в ответе
	roman := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C"}
	fmt.Println(roman[a])
}
func chetyre(a, b int, c string) { //выполнение операций с ответом для арабских
	if c == "+" {
		d := a + b
		fmt.Println(d)
	} else if c == "-" {
		d := a - b
		fmt.Println(d)
	} else if c == "*" {
		d := a * b
		fmt.Println(d)
	} else if c == "/" {
		d := a / b
		fmt.Println(d)
	}
}
func pyat(a, b int, c string) int { //выполнение операций для римских
	var d int
	if c == "+" {
		d = a + b
	} else if c == "-" {
		d = a - b
	} else if c == "*" {
		d = a * b
	} else if c == "/" {
		d = a / b
	}
	return d
}

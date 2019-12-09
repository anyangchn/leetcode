package main

func intToRoman(num int) string {
	// clear code
	return solutionA(num)

	// quick code
	//return solutionB(num)
}

func solutionA(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string
	for idx, val := range values {
		if num >= val {
			s := symbols[idx]
			for quo := num / val; quo > 0; quo-- {
				result += s
				num -= val
			}
		}
	}
	return result
}

func solutionB(num int) string {
	var tmp, result string
	tmp, num = baseRoman(num, 1000, 100, "M", "C")
	result += tmp
	tmp, num = baseRoman(num, 500, 100, "D", "C")
	result += tmp
	tmp, num = baseRoman(num, 100, 10, "C", "X")
	result += tmp
	tmp, num = baseRoman(num, 50, 10, "L", "X")
	result += tmp
	tmp, num = baseRoman(num, 10, 1, "X", "I")
	result += tmp
	tmp, num = baseRoman(num, 5, 1, "V", "I")
	result += tmp
	for i := 0; i < num; i++ {
		result += "I"
	}
	return result
}

func baseRoman(num, div, base int, roam, roamL string) (string, int) {
	result := ""
	quo, rem := num/div, num%div
	for i := 0; i < quo; i++ {
		result += roam
	}
	num = rem
	if num+base >= div {
		result += roamL + roam
		num = num - div + base
	}
	return result, num
}

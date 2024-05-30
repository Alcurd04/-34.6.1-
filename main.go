package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: программа <входной_файл> <выходной_файл>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка чтения входного файла:", err)
		return
	}

	expressions := strings.Split(string(inputData), "\n")
	resultMap := make(map[string]string)

	for _, exp := range expressions {
		re := regexp.MustCompile("(/d+)([+-])(/d+)=/?")
		match := re.FindStringSubmatch(exp)
		if len(match) == 4 {
			num1 := match[1]
			operator := match[2]
			num2 := match[3]

			res := calculateResult(num1, operator, num2)
			resultMap[exp] = res
		}
	}

	outputData := ""
	for key, value := range resultMap {
		outputData += fmt.Sprintf("%s%s\n", key, value)
	}

	err = ioutil.WriteFile(outputFile, []byte(outputData), 0644)
	if err != nil {
		fmt.Println("Ошибка записи в выходной файл:", err)
		return
	}

	fmt.Println("Результаты успешно записаны в файл", outputFile)
}

func calculateResult(num1, operator, num2 string) string {
	n1 := atoi(num1)
	n2 := atoi(num2)

	switch operator {
	case "+":
		return fmt.Sprintf("=%d\n", n1+n2)
	case "-":
		return fmt.Sprintf("=%d\n", n1-n2)
	default:
		return ""
	}
}

func atoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

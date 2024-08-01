package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("### Starting program")

	fileName := ParseInput()

	if len(fileName) == 0 {
		fmt.Println("@@@ Exiting program")
		return
	}

	if len(fileName) > 3 && fileName[len(fileName)-4:] != ".csv" {
		fileName = fileName + ".csv"
	}

	fmt.Println("### Loaded file name is", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("@@@ Failed to open file", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("@@@ Error reading CSV file:", err)
		return
	}

	var numbers []int

	for _, record := range records {
		for _, value := range record {
			number, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			numbers = append(numbers, number)
		}
	}

	fmt.Println("### Total numbers in the file:", len(numbers))

	results := newMarkSequences(numbers)

	err = saveToExcel(numbers, results)
	if err != nil {
		fmt.Println("@@@ Error saving to excel file:", err)
		return
	}

	fmt.Println("### Exiting program")
}

func saveToExcel(input, result []int) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	// Set headers
	f.SetCellValue(sheet, "A1", "Index")
	f.SetCellValue(sheet, "B1", "Input")
	f.SetCellValue(sheet, "C1", "Result")

	// Populate data
	for i := 0; i < len(input); i++ {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), i)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), input[i])
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), result[i])
	}

	// Save file
	if err := f.SaveAs("output.xlsx"); err != nil {
		return err
	}

	return nil
}

func newMarkSequences(input []int) []int {
	length := len(input)

	result := make([]int, 0, length)

	slices.Reverse(input)

	var counter int

	// 針對倒過來過來的輸入：
	// 如果是 0，標記其 0
	// 如果為大於五的數字，標記其 1，並且將 counter 設為該數字減一。counter 代表的意義為後面還有幾個數字要標記為 1
	// 如果 counter 大於 0，標記其 1，並且將 counter 減一
	// 其他情況標記其 0
	for i, number := range input {
		if number == 0 {
			result[i] = 0
			continue
		}

		if counter > 0 {
			result[i] = 1
			counter--
			continue
		}

		if number >= 5 {
			result[i] = 1
			counter = number - 1
			continue
		}

		result[i] = 0
	}

	slices.Reverse(result)

	return result
}

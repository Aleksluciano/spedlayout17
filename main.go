package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	//explain the code below
	//walk through the spedlayout16 folder
	err := filepath.Walk("spedlayout16", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

			//open the file
		file, err2 := OpenFile(path, err)
		if err2 != nil {
			return err2
		}
		defer file.Close()

		//read the file
		lines, err6 := ChangeContent(file)
		if err6 != nil {
			print(path)
			print(err6)
			return filepath.SkipDir
		}


		//create a new file
		newFile, err3 := CreateNewPathFile(path, err)
		if err3 != nil {
			return err3
		}
		defer newFile.Close()
		//write the new file
		err4 := SaveNewFile(lines, newFile)
		if err4 != nil {
			return err4
		}

		file.Close()
		err5 := ChangeOriginalFileName(path)
		if err5 != nil {
			return err5
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func ChangeOriginalFileName(path string) error {
	err5 := os.Rename(path, "processado/" + filepath.Base(path))
	if err5 != nil {
		return err5
	}
	return nil
}

func SaveNewFile(lines []string, newFile *os.File) error {
	for _, line := range lines {
		_, err := newFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateNewPathFile(path string, err error) (*os.File, error) {
	newFilePath := strings.Replace(path, "spedlayout16", "spedlayout17", 1)
	print("Processado: ", newFilePath)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func ChangeContent(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := []string{}
	c100_total := 0
	d100_total := 0
	//read the file line by line

	//make the code below get the index for the for statement
	indexFile := 0
	for scanner.Scan() {
		line := scanner.Text()
		indexFile++

		//rules to change the file

		isVersion16 := CheckVersion16(indexFile, line)
		if isVersion16 == false {
			return nil, errors.New("arquivo não é versão 16")
		}

		line = changeLayoutTo17(indexFile, line)

		skip, c100, d100 := checkValuesToSkip(line)
		if skip == true {
			c100_total += c100
			d100_total += d100
			continue
		}
        line = changeRegC990(c100_total, line)
		line = changeRegD990(d100_total, line)
		line = changeReg9900_C100(c100_total, line)
		line = changeReg9900_D100(d100_total, line)
		line = changeRegK990(line)
		line = changeReg9900_9900(line)
		line = changeReg9990(line)
		line = changeReg9999(c100_total + d100_total,line)
		lines = append(lines, line)
		lines = addRegK010_1(lines, line)
		lines = addReg9900_K010_1(lines, line)

	}
	return lines, nil
}

func changeReg9900_D100(total int, line string) string {
	if strings.Contains(line, "|9900|D100|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "9900" && parts[2] == "D100" {
			value, err := strconv.Atoi(parts[3])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value -= total
			parts[3] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeReg9900_C100(total int, line string) string {
	if strings.Contains(line, "|9900|C100|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "9900" && parts[2] == "C100" {
			value, err := strconv.Atoi(parts[3])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value -= total
			parts[3] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeRegD990(total int, line string) string {
	if strings.Contains(line, "|D990|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "D990"	{
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value -= total
			parts[2] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeRegC990(total int, line string) string {
	if strings.Contains(line, "|C990|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "C990"	{
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value -= total
			parts[2] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func checkValuesToSkip(line string) (bool, int, int) {
	if strings.Contains(line, "|C100|0|0||55|05|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "C100" &&
			parts[2] == "0" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "55" &&
			parts[6] == "05" {
			return true, 1,0
		}
	}
	if strings.Contains(line, "|C100|1|0||55|05|") {
       		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "C100" &&
			parts[2] == "1" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "55" &&
			parts[6] == "05" {
			return true, 1,0
		}
	}
	if strings.Contains(line, "|C100|0|0||55|04|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "C100" &&
			parts[2] == "0" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "55" &&
			parts[6] == "04" {
			return 	true, 1,0
		}

	}
	if strings.Contains(line, "|C100|1|0||55|04|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "C100" &&
			parts[2] == "1" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "55" &&
			parts[6] == "04" {
			return true, 1,0
		}
	}

	if strings.Contains(line, "|D100|0|0||57|05|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "D100" &&
			parts[2] == "0" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "57" &&
			parts[6] == "05" {
			return true, 0,1
		}
	}
if strings.Contains(line, "|D100|1|0||57|05|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "D100" &&
			parts[2] == "1" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "57" &&
			parts[6] == "05" {
			return true, 0,1
		}
}

	if strings.Contains(line, "|D100|0|0||57|04|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "D100" &&
			parts[2] == "0" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "57" &&
			parts[6] == "04" {
			return true, 0,1
		}
	}

	if strings.Contains(line, "|D100|1|0||57|04|") {
		parts := strings.Split(line, "|")
		if len(parts) > 6 &&
			parts[1] == "D100" &&
			parts[2] == "1" &&
			parts[3] == "0" &&
			parts[4] == "" &&
			parts[5] == "57" &&
			parts[6] == "04" {
			return true, 0,1
		}
	}



	return false, 0,0
}

func CheckVersion16(indexFile int, line string) bool {
	if indexFile == 1 {
		parts := strings.Split(line, "|")

		if len(parts) > 1 {
			value := parts[2]
			if value == "017" {
				return false
			}
		}
	}
	return true
}

func changeLayoutTo17(indexFile int, line string) string {
	if indexFile == 1 {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "0000" {
			value := parts[2]
			if value == "016" {
				parts[2] = "017"
				newline := strings.Join(parts, "|")
				return newline
			}
		}
	}
	return line
}
func addRegK010_1(lines []string, line string) []string {
	if strings.Contains(line, "|K001|0|") {
		parts := strings.Split(line, "|")
		if len(parts) > 2  && parts[1] == "K001" && parts[2] == "0" {
			lines = append(lines, "|K010|1|")
		}
	}
	return lines
}

func addReg9900_K010_1(lines []string, line string) []string {
	if strings.Contains(line, "|9900|K001|1|") {
		parts := strings.Split(line, "|")
		if len(parts) > 2  && parts[1] == "9900" && parts[2] == "K001" && parts[3] == "1" {
			lines = append(lines, "|9900|K010|1|")
		}
	}
	return lines
}

func changeReg9900_9900(line string) string {
	if strings.Contains(line, "|9900|9900|") {
		parts := strings.Split(line, "|")
		if len(parts) > 2  && parts[1] == "9900" && parts[2] == "9900" {
			value, err := strconv.Atoi(parts[3])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value++
			parts[3] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeReg9990(line string) string {
	if strings.Contains(line, "|9990|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "9990"	{
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value++
			parts[2] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeRegK990(line string) string {
	if strings.Contains(line, "|K990|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "K990" {
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
				return line
			}
			value++
			parts[2] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func changeReg9999(total int ,line string) string {
	if strings.Contains(line, "|9999|") {
		parts := strings.Split(line, "|")
		if len(parts) > 1 && parts[1] == "9999"{
			value, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
			}
			value += 2
			value -= total
			parts[2] = strconv.Itoa(value)
			newLine := strings.Join(parts, "|")
			return newLine
		}
	}
	return line
}

func OpenFile(path string, err error) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

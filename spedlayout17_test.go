package main

//crate a test for package main

import (
	"strings"
	"testing"
)


func TestCheckVersion16(t *testing.T) {
	var line = "|0000|016|"
	isVersion16 := CheckVersion16(1, line)
	if !isVersion16 {
		t.Errorf("Version 16 not found")
	}
}

func TestChangeLayoutTo17(t *testing.T) {
	var line = "|0000|016|"
	newline := changeLayoutTo17(1, line)
	if !strings.Contains(newline, "|0000|017|") {
		t.Errorf("Version %s not changed to |0000|017|, value = %s", line,newline)
	}
}

func TestChangeRegK990 (t *testing.T) {
	var line = "|K990|1535|"
	newline := changeRegK990(line)
	if !strings.Contains(newline, "|K990|1536|") {
		t.Errorf("Reg %s not changed to %s + 1, value = %s", line,line,newline)
	}
}


func TestChangeReg9900_9900 (t *testing.T) {
	var line = "|9900|9900|53|"
	newline := changeReg9900_9900(line)
	if !strings.Contains(newline, "|9900|9900|54|") {
		t.Errorf("Reg %s not changed to %s + 1, value = %s", line,line,newline)
	}
}

func TestChangeReg9990 (t *testing.T) {
	var line = "|9990|2|"
	newline := changeReg9990(line)
	if !strings.Contains(newline, "|9990|3|") {
		t.Errorf("Reg %s not changed to %s + 1, value = %s", line,line,newline)
	}
}

func TestChangeReg9999 (t *testing.T) {
	var line = "|9999|9373|"
	newline := changeReg9999(5 + 0,line)
	if !strings.Contains(newline, "|9999|9370|") {
		t.Errorf("Reg %s not changed to %s + 2 - 5, value = %s", line,line,newline)
	}
}

func TestAddRegK010_1 (t *testing.T) {
	var line = "|K001|0|"
	lines := addRegK010_1([]string{line}, line)
	if len(lines) < 2 || !strings.Contains(lines[1], "|K010|1|") {
		t.Errorf("Reg |K010|1| not added to array")
	}
}

func TestAddReg9900_K010_1 (t *testing.T) {
	var line = "|9900|K001|1|"
	lines := addReg9900_K010_1([]string{line}, line)
	if len(lines) < 2 || !strings.Contains(lines[1], "|9900|K010|1|") {
		t.Errorf("Reg |9900|K010|1| not added to array")
	}
}

func TestCheckValuesToSkip (t *testing.T) {
	var lines = []string{"|C100|0|0||55|05|" , "|C100|1|0||55|05|" , "|C100|0|0||55|04|" , "|C100|1|0||55|04|" , "|D100|0|0||57|05|" , "|D100|1|0||57|05|" , "|D100|0|0||57|04|" , "|D100|1|0||57|04|"}
	for _, line := range lines {
		if skip, _,_:= checkValuesToSkip(line) ; !skip {
			t.Errorf("Found a value to not skip: %s", line)
		}
	}

}

func TestCountValuesToSkip (t *testing.T) {
	var lines = []string{"|C100|0|0||55|05|" , "|C100|1|0||55|05|" , "|C100|0|0||55|04|" , "|C100|1|0||55|04|" , "|D100|0|0||57|05|" , "|D100|1|0||57|05|" , "|D100|0|0||57|04|" , "|D100|1|0||57|04|"}
	var c100_total, d100_total  = 0,0
	for _, line := range lines {
		_, c100, d100 := checkValuesToSkip(line)
		c100_total += c100
		d100_total += d100

	}
	if c100_total != 4 || d100_total != 4 {
		t.Errorf("Total to skip wrong c100_total should be %d and d100_total should be %d", c100_total, d100_total)
	}

}

func TestChange9900_D100 (t *testing.T) {
	var line = "|9900|D100|4|"
	newline := changeReg9900_D100(3,line)
	if !strings.Contains(newline, "|9900|D100|1|") {
		t.Errorf("Reg %s not changed to %s - 3, value = %s", line,line,newline)
	}
}

func TestChange9900_C100 (t *testing.T) {
	var line = "|9900|C100|4|"
	newline := changeReg9900_C100(3,line)
	if !strings.Contains(newline, "|9900|C100|1|") {
		t.Errorf("Reg %s not changed to %s - 3, value = %s", line,line,newline)
	}
}

func TestChangeC990 (t *testing.T) {
 	var line = "|C990|1535|"
	newline := changeRegC990(2,line)
	if !strings.Contains(newline, "|C990|1533|") {
		t.Errorf("Reg %s not changed to %s - 2, value = %s", line,line,newline)
	}
}

func TestChangeD990 (t *testing.T) {
	var line = "|D990|1535|"
	newline := changeRegD990(2,line)
	if !strings.Contains(newline, "|D990|1533|") {
		t.Errorf("Reg %s not changed to %s - 2, value = %s", line,line,newline)
	}
}
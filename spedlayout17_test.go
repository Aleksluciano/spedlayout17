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
	var line = "|9999|2|"
	newline := changeReg9999(line)
	if !strings.Contains(newline, "|9999|4|") {
		t.Errorf("Reg %s not changed to %s + 2, value = %s", line,line,newline)
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
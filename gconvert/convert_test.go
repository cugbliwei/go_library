package gconvert

import "testing"

func TestStringToInt(t *testing.T) {
	num := StringToInt("100")
	if num != 100 {
		t.Error(num)
	}
	t.Log(num)
}

func TestIntToString(t *testing.T) {
	num := IntToString(100)
	if num != "100" {
		t.Error(num)
	}
	t.Log(num)
}

func TestFloat64ToString(t *testing.T) {
	num := Float64ToString(0.1010, 3)
	if num != "0.101" {
		t.Error(num)
	}
	t.Log(num)
}

func TestStringToFloat64(t *testing.T) {
	num := StringToFloat64("0.1010")
	if num != 0.101 {
		t.Error(num)
	}
	t.Log(num)
}

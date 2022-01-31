package main

import "testing"

func TestFindian(t *testing.T) {
	if !findian("ian") {
		t.Errorf("Expected true but got %v", findian("ian"))
	}

	if !findian("Ian") {
		t.Errorf("Expected true but got %v", findian("Ian"))
	}

	if !findian("iuiygaygn") {
		t.Errorf("Expected true but got %v", findian("iuiygaygn"))
	}

	if !findian("I d skd a efju N") {
		t.Errorf("Expected true but got %v", findian("I d skd a efju N"))
	}

	if findian("ihhhhhn") {
		t.Errorf("Expected false but got %v", findian("ihhhhhn"))
	}

	if findian("ina") {
		t.Errorf("Expected false but got %v", findian("ina"))
	}

	if findian("xian") {
		t.Errorf("Expected false but got %v", findian("xian"))
	}
}

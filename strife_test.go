package strife

import (
	"fmt"
	"testing"
)

func TestPadStringLongerThenNumber(t *testing.T) {
	strOrg := "aaaaa"
	str := pad(strOrg, uint32(4))
	if str != strOrg {
		t.Errorf("it should return %v but it returned %v", strOrg, str)
	}
}

func TestPadNumberIsEven(t *testing.T) {
	strOrg := "aaaaa"
	str := pad(strOrg, uint32(7))
	ideal := fmt.Sprintf(" %v ", strOrg)
	if str != ideal {
		t.Errorf("it should return %v but it returned %v", ideal, str)
	}
}

func TestPadNumberIsOdd(t *testing.T) {
	strOrg := "aaaaa"
	str := pad(strOrg, uint32(8))
	ideal := fmt.Sprintf("  %v ", strOrg)
	if str != ideal {
		t.Errorf("it should return %v but it returned %v", ideal, str)
	}
}

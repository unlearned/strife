package strife

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPredictText(t *testing.T) {
	text, err := PredictText(100, "./testdata/interview.json")
	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}
	ideal := idealText()
	if text != ideal {
		t.Errorf("it should return \n%v but it returned %v", ideal, text)
	}

}

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

func TestReadJSONSuccess(t *testing.T) {
	p, err := readJSON("./testdata/interview.json")

	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if reflect.TypeOf(*p).String() != "strife.Recruting" {
		t.Errorf("it should return strife.Recruting but it returned %v", reflect.TypeOf(*p))
	}
}

func TestRound2Success(t *testing.T) {
	r := round2(1.45678)
	if r != 1.46 {
		t.Errorf("it should return 1.46 but it returned %v", r)
	}
}

func idealText() string {
	ideal := "## Past Performance ##\n"
	ideal += "|         phase        |       entory      |   casual meeting  | document screening|technical interview|   2nd interview   |  final interview  |   offer meeting   |       joined      |\n"
	ideal += "|         number       |        1000       |        900        |        800        |        700        |        600        |        500        |        400        |        300        |\n"
	ideal += "|average hours required|         0         |        1.5        |        1.5        |        1.5        |        1.5        |        1.5        |        1.5        |         0         |\n"
	ideal += "\n"
	ideal += "## Prediction ##\n"
	ideal += "|         phase        |       entory      |   casual meeting  | document screening|technical interview|   2nd interview   |  final interview  |   offer meeting   |       joined      |\n"
	ideal += "|         number       |        338        |        304        |        270        |        236        |        202        |        168        |        134        |        100        |\n"
	ideal += "|average hours required|         0         |        1.5        |        1.5        |        1.5        |        1.5        |        1.5        |        1.5        |         0         |\n"
	ideal += "| total hours required |         0         |        456        |        405        |        354        |        303        |        252        |        201        |         0         |\n"
	return ideal
}

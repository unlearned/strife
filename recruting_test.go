package strife

import (
	"reflect"
	"testing"
)

func TestReadJSONSuccess(t *testing.T) {
	p, err := ReadJSON("./testdata/interview.json")

	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if reflect.TypeOf(*p).String() != "strife.Recruting" {
		t.Errorf("it should return strife.Recruting but it returned %v", reflect.TypeOf(*p))
		t.Fatal(err)
	}
}

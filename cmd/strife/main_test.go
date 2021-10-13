package main

import "testing"

func TestConvNumTypeSuccess(t *testing.T) {
	n, err := convNumType("100")
	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if n != 100 {
		t.Errorf("it should return 100 but it returned %v", n)
	}
}

func TestConvNumTypeFailedByString(t *testing.T) {
	_, err := convNumType("a")
	if err == nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}
}

func TestConvNumTypeFailedBy0(t *testing.T) {
	_, err := convNumType("0")
	if err == nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}
}

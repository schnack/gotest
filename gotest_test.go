package gotest

import (
	"fmt"
	"testing"
)

func TestGoTest_Error(t *testing.T) {
	if err := Expect(fmt.Errorf("test")).Error("test"); err != nil {
		t.Error(err)
	}
}

func TestGoTest_True(t *testing.T) {
	if err := Expect(true).True(); err != nil {
		t.Error(err)
	}
}

func TestGoTest_False(t *testing.T) {
	if err := Expect(false).False(); err != nil {
		t.Error(err)
	}
}

func TestGoTest_Eq(t *testing.T) {
	if err := Expect("test").Eq("test"); err != nil {
		t.Error(err)
	}

	if err := Expect([]string{"test1", "test2"}).Eq([]string{"test1", "test2"}); err != nil {
		t.Error(err)
	}

	if err := Expect([]byte{1, 2, 3, 4}).Eq([]byte{1, 2, 3, 4}); err != nil {
		t.Error(err)
	}
}

func TestGoTest_Nil(t *testing.T) {
	var param *GoTest = nil

	if err := Expect(param).Nil(); err != nil {
		t.Error(err)
	}

	var err error
	if err := Expect(err).Nil(); err != nil {
		t.Error(err)
	}
}

func TestGoTest_NotNil(t *testing.T) {
	param := &goTest{}
	if err := Expect(param).NotNil(); err != nil {
		t.Error(err)
	}
}

func TestGoTest_Zero(t *testing.T) {
	var param int
	if err := Expect(param).Zero(); err != nil {
		t.Error(err)
	}
}

func TestGoTest_NotZero(t *testing.T) {
	param := 1
	if err := Expect(param).NotZero(); err != nil {
		t.Error(err)
	}
}

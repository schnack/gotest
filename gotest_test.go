package gotest

import "testing"

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

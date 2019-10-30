package gotest

import (
	"fmt"
	"reflect"
)

// Библиотека для тестирования
// Пример использования
// 		if err := gotest.Expect("string111").Eq("string11"); err != nil { t.Error(err) }

func Expect(expect interface{}) GoTest {
	return &goTest{expect: expect}
}

type GoTest interface {
	Eq(v interface{}) error
	IsZero() error
	NotIsZero() error
}

type goTest struct {
	expect interface{}
}

func (gt *goTest) Eq(v interface{}) error {
	if !reflect.DeepEqual(v, gt.expect) {
		xSample := reflect.ValueOf(v)
		xExpect := reflect.ValueOf(gt.expect)
		expect := fmt.Sprintf("%v", gt.expect)
		sample := fmt.Sprintf("%v", v)
		if xExpect.Kind() == reflect.String {
			expect = fmt.Sprintf(`"%s"`, expect)
		}
		if xSample.Kind() == reflect.String {
			sample = fmt.Sprintf(`"%s"`, sample)
		}
		return fmt.Errorf("\n\nexpected: %v\n     got: %v\n\n(compared using reflect.DeepEqual)\n", sample, expect)
	}
	return nil
}

func (gt *goTest) IsZero() error {
	if reflect.ValueOf(gt.expect).IsZero() {
		return nil
	}
	return fmt.Errorf("\n\nIs not zero\n\n(compared using reflect.IsZero)\n")
}

func (gt *goTest) IsNil() error {
	if reflect.ValueOf(gt.expect).IsNil() {
		return nil
	}
	return fmt.Errorf("\n\nIs not nil\n\n(compared using reflect.IsNil)\n")
}

func (gt *goTest) NotIsZero() error {
	if !reflect.ValueOf(gt.expect).IsZero() {
		return nil
	}
	return fmt.Errorf("\n\nIs zero\n\n(compared using reflect.IsZero)\n")
}

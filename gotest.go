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
	Error(str string) error
	Zero() error
	NotZero() error
	Nil() error
	NotNil() error
	True() error
	False() error
}

type goTest struct {
	expect interface{}
}

func (gt *goTest) Error(str string) error {
	return gt.Eq(fmt.Errorf(str))
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

func (gt *goTest) True() error {
	switch x := gt.expect.(type) {
	case bool:
		if x {
			return nil
		}
		return fmt.Errorf("\n\nIs not true\n")
	default:
		return fmt.Errorf("\n\nIs not bool type\n")
	}
}

func (gt *goTest) False() error {
	switch x := gt.expect.(type) {
	case bool:
		if !x {
			return nil
		}
		return fmt.Errorf("\n\nIs not false\n")
	default:
		return fmt.Errorf("\n\nIs not bool type\n")
	}
}

func (gt *goTest) Zero() error {
	if reflect.ValueOf(gt.expect).IsZero() {
		return nil
	}
	return fmt.Errorf("\n\nIs not zero\n\n(compared using reflect.IsZero)\n")
}

func (gt *goTest) NotZero() error {
	if !reflect.ValueOf(gt.expect).IsZero() {
		return nil
	}
	return fmt.Errorf("\n\nIs zero\n\n(compared using reflect.IsZero)\n")
}

func (gt *goTest) Nil() error {
	v := reflect.ValueOf(gt.expect)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if v.IsNil() {
			return nil
		}
	default:
		if !v.IsValid() {
			return nil
		}
	}
	return fmt.Errorf("\n\nIs not nil\n\n(compared using reflect.IsNil)\n")
}

func (gt *goTest) NotNil() error {
	v := reflect.ValueOf(gt.expect)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if !v.IsNil() {
			return nil
		}
	default:
		if v.IsValid() {
			return nil
		}
	}
	return fmt.Errorf("\n\nIs nil\n\n(compared using reflect.IsNil)\n")
}

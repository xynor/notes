package main

import (
	"errors"
	"fmt"
	"testing"
)

var BaseErr = errors.New("base error")
var UpErr = errors.New("up error")
var Up2Err = errors.New("up2 error")

func TestWrap(t *testing.T) {
	err := upbase()
	fmt.Println(err)
	fmt.Println(errors.Is(err, BaseErr))
}

func TestWrap2(t *testing.T) {
	err := up2base()
	fmt.Println(err)
	fmt.Println(errors.Is(err, BaseErr))
}

func base() error {
	return BaseErr
}

func upbase() error {
	err := base()
	return fmt.Errorf("upbase error[%w]", err)
}

func up2base() error {
	err := upbase()
	return fmt.Errorf("up2base error[%w]", err)
}

func upbaseError() error {
	err := base()
	return errors.Join(UpErr, err)
}

func up2baseError() error {
	err := upbaseError()
	return errors.Join(Up2Err, err)
}

func TestIs(t *testing.T) {
	err := upbaseError()
	fmt.Println(err)
	fmt.Println(errors.Is(err, BaseErr))
}

func TestIs1(t *testing.T) {
	err := upbaseError()
	fmt.Println(err)
	fmt.Println(errors.Is(err, UpErr))
}

func TestIs2(t *testing.T) {
	err := upbaseError()
	fmt.Println(err)
	fmt.Println(errors.Is(err, Up2Err)) //fase
}

func TestIs3(t *testing.T) {
	err := up2baseError()
	fmt.Println(err)
	fmt.Println(errors.Is(err, Up2Err))
}

func TestIs4(t *testing.T) {
	err := up2baseError()
	fmt.Println(err)
	fmt.Println(errors.Is(err, BaseErr))
}

// Struct
type BaseError struct {
	Code int
	Msg  string
}

func (c *BaseError) Error() string {
	return fmt.Sprintf("%s, code=%d", c.Msg, c.Code)
}

// Struct
type UpError struct {
	Code int
	Msg  string
}

func (c *UpError) Error() string {
	return fmt.Sprintf("%s, code=%d", c.Msg, c.Code)
}

func structBaseError() error {
	return &BaseError{
		Code: 1,
		Msg:  "base error",
	}
}

func structUpError() error {
	err := structBaseError()
	return errors.Join(&UpError{
		Code: 2,
		Msg:  "up2 error",
	}, err)
}

func TestStructBase(t *testing.T) {
	err := structBaseError()
	fmt.Println(err)
}

func TestStructUp(t *testing.T) {
	err := structUpError()
	fmt.Println(err)
}

func TestStructAs(t *testing.T) {
	err := structUpError()
	var BaseErr *BaseError
	if errors.As(err, &BaseErr) {
		fmt.Println("BaseErr:", BaseErr)
	}
}

func TestStructAs2(t *testing.T) {
	err := structUpError()
	var up2Err *UpError
	if errors.As(err, &up2Err) {
		fmt.Println("up2Err:", up2Err)
		fmt.Println("err:", err)
	}
}

var bError = structBaseError()
var uError = structUpError()

func TestStructAsWarp(t *testing.T) {
	err := bError
	err = fmt.Errorf("%w:%w", uError, err)
	fmt.Println(err)
	fmt.Println(errors.Is(err, bError))
	fmt.Println(errors.Is(err, uError))
}

func TestStructAsWarp2(t *testing.T) {
	err := structBaseError()
	err = fmt.Errorf("%s:%w", "dsfsdf", err)
	fmt.Println(err)
	fmt.Println(errors.Is(err, bError))
	fmt.Println(errors.Is(err, uError))
	//AS
	var b *BaseError
	if errors.As(err, &b) { //True
		fmt.Println("b:", b)
	}
}

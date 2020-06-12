package assert

import (
	"time"

	"github.com/stretchr/testify/assert"
)

type TestingT = assert.TestingT

func ObjectsAreEqual(actual, expected interface{}) bool {
	return assert.ObjectsAreEqual(expected, actual)
}

func ObjectsAreEqualValues(actual, expected interface{}) bool {
	return assert.ObjectsAreEqualValues(expected, actual)
}

var FailNow = assert.FailNow
var Fail = assert.Fail
var Implements = assert.Implements

func IsType(t TestingT, object interface{}, expectedType interface{}, msgAndArgs ...interface{}) bool {
	return assert.IsType(t, expectedType, object, msgAndArgs...)
}

func Equal(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}

func Same(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.Same(t, expected, actual, msgAndArgs...)
}

func EqualValues(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.EqualValues(t, expected, actual, msgAndArgs...)
}

func Exactly(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.Exactly(t, expected, actual)
}

var NotNil = assert.NotNil
var Nil = assert.Nil
var Empty = assert.Empty
var NotEmpty = assert.NotEmpty
var Len = assert.Len // expected length comes after as expected.
var True = assert.True
var False = assert.False

func NotEqual(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqual(t, expected, actual, msgAndArgs...)
}

func NotEqualValues(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotEqualValues(t, expected, actual, msgAndArgs...)
}

var Contains = assert.Contains
var NotContains = assert.NotContains
var Subset = assert.Subset
var NotSubset = assert.NotSubset
var ElementsMatch = assert.ElementsMatch
var Condition = assert.Condition

type PanicTestFunc = assert.PanicTestFunc

var Panics = assert.Panics

func PanicsWithValue(t TestingT, f PanicTestFunc, expected interface{}, msgAndArgs ...interface{}) bool {
	return assert.PanicsWithValue(t, expected, f, msgAndArgs...)
}

func PanicsWithError(t TestingT, f PanicTestFunc, errString string, msgAndArgs ...interface{}) bool {
	return assert.PanicsWithError(t, errString, f, msgAndArgs...)
}

var NotPanics = assert.NotPanics

func WithinDuration(t TestingT, actual, expected time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	return assert.WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

func InDelta(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return assert.InDelta(t, expected, actual, delta, msgAndArgs...)
}

func InDeltaSlice(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return assert.InDeltaSlice(t, expected, actual, delta, msgAndArgs...)
}

func InDeltaMapValues(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return assert.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...)
}

func InEpsilon(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return assert.InEpsilon(t, expected, actual, delta, msgAndArgs...)
}

func InEpsilonSlice(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return assert.InEpsilonSlice(t, expected, actual, delta, msgAndArgs...)
}

var NoError = assert.NoError
var Error = assert.Error
var EqualError = assert.EqualError // expected error comes after as expected.

func Regexp(t TestingT, str interface{}, rx interface{}, msgAndArgs ...interface{}) bool {
	return assert.Regexp(t, rx, str, msgAndArgs...)
}

func NotRegexp(t TestingT, str interface{}, rx interface{}, msgAndArgs ...interface{}) bool {
	return assert.NotRegexp(t, rx, str, msgAndArgs...)
}

var Zero = assert.Zero
var NotZero = assert.NotZero
var FileExists = assert.FileExists
var NoFileExists = assert.NoFileExists
var DirExists = assert.DirExists
var NoDirExists = assert.NoDirExists

func JSONEq(t TestingT, actual, expected string, msgAndArgs ...interface{}) bool {
	return assert.JSONEq(t, expected, actual, msgAndArgs...)
}

func YAMLEq(t TestingT, actual, expected string, msgAndArgs ...interface{}) bool {
	return assert.YAMLEq(t, expected, actual, msgAndArgs...)
}

var Eventually = assert.Eventually
var Never = assert.Never

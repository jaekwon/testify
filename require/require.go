package require

import (
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestingT = require.TestingT

/*
func ObjectsAreEqual(actual, expected interface{}) {
	require.ObjectsAreEqual(expected, actual)
}

func ObjectsAreEqualValues(actual, expected interface{}) {
	require.ObjectsAreEqualValues(expected, actual)
}
*/

var FailNow = require.FailNow
var Fail = require.Fail
var Implements = require.Implements

func IsType(t TestingT, object interface{}, expectedType interface{}, msgAndArgs ...interface{}) {
	require.IsType(t, expectedType, object, msgAndArgs...)
}

func Equal(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.Equal(t, expected, actual, msgAndArgs...)
}

func Same(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.Same(t, expected, actual, msgAndArgs...)
}

func EqualValues(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.EqualValues(t, expected, actual, msgAndArgs...)
}

func Exactly(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.Exactly(t, expected, actual)
}

var NotNil = require.NotNil
var Nil = require.Nil
var Empty = require.Empty
var NotEmpty = require.NotEmpty
var Len = require.Len // expected length comes after as expected.
var True = require.True
var False = require.False

func NotEqual(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.NotEqual(t, expected, actual, msgAndArgs...)
}

func NotEqualValues(t TestingT, actual, expected interface{}, msgAndArgs ...interface{}) {
	require.NotEqualValues(t, expected, actual, msgAndArgs...)
}

var Contains = require.Contains
var NotContains = require.NotContains
var Subset = require.Subset
var NotSubset = require.NotSubset
var ElementsMatch = require.ElementsMatch
var Condition = require.Condition

type PanicTestFunc = assert.PanicTestFunc

var Panics = require.Panics

func PanicsWithValue(t TestingT, f PanicTestFunc, expected interface{}, msgAndArgs ...interface{}) {
	require.PanicsWithValue(t, expected, f, msgAndArgs...)
}

func PanicsWithError(t TestingT, f PanicTestFunc, errString string, msgAndArgs ...interface{}) {
	require.PanicsWithError(t, errString, f, msgAndArgs...)
}

var NotPanics = require.NotPanics

func WithinDuration(t TestingT, actual, expected time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	require.WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

func InDelta(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDelta(t, expected, actual, delta, msgAndArgs...)
}

func InDeltaSlice(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDeltaSlice(t, expected, actual, delta, msgAndArgs...)
}

func InDeltaMapValues(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...)
}

func InEpsilon(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InEpsilon(t, expected, actual, delta, msgAndArgs...)
}

func InEpsilonSlice(t TestingT, actual, expected interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InEpsilonSlice(t, expected, actual, delta, msgAndArgs...)
}

var NoError = require.NoError
var Error = require.Error
var EqualError = require.EqualError // expected error comes after as expected.

func Regexp(t TestingT, str interface{}, rx interface{}, msgAndArgs ...interface{}) {
	require.Regexp(t, rx, str, msgAndArgs...)
}

func NotRegexp(t TestingT, str interface{}, rx interface{}, msgAndArgs ...interface{}) {
	require.NotRegexp(t, rx, str, msgAndArgs...)
}

var Zero = require.Zero
var NotZero = require.NotZero
var FileExists = require.FileExists
var NoFileExists = require.NoFileExists
var DirExists = require.DirExists
var NoDirExists = require.NoDirExists

func JSONEq(t TestingT, actual, expected string, msgAndArgs ...interface{}) {
	require.JSONEq(t, expected, actual, msgAndArgs...)
}

func YAMLEq(t TestingT, actual, expected string, msgAndArgs ...interface{}) {
	require.YAMLEq(t, expected, actual, msgAndArgs...)
}

var Eventually = require.Eventually
var Never = require.Never

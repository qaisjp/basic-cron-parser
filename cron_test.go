package cron

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseTest(t *testing.T, str string, min, max int) []string {
	result, err := parse(str, min, max)
	assert.NoError(t, err)
	return result
}

func TestParseExample(t *testing.T) {
	assert.Equal(t, []string{"0", "15", "30", "45"}, parseTest(t, "*/15", MinMinutes, MaxMinutes))      // minute
	assert.Equal(t, []string{"0"}, parseTest(t, "0", MinHours, MaxHours))                               // hour
	assert.Equal(t, []string{"1", "15"}, parseTest(t, "1,15", MinDayOfMonth, MaxDayOfMonth))            // day of month
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, parseTest(t, "1-5", MinDayOfWeek, MaxDayOfWeek)) // day of week

	// Months
	expectedMonths := []string{}
	for i := MinMonths; i <= MaxMonths; i++ {
		expectedMonths = append(expectedMonths, strconv.Itoa(i))
	}
	assert.Equal(t, expectedMonths, parseTest(t, "*", MinMonths, MaxMonths)) // months
}

func TestParseMaxRange(t *testing.T) {
	// individual values
	assert.Equal(t, []string{"0"}, parseTest(t, "0", 0, 5))
	assert.Equal(t, []string{"5"}, parseTest(t, "5", 0, 5))

	// comma separated lists
	assert.Equal(t, []string{"0", "2", "5"}, parseTest(t, "0,2,5", 0, 5))

	// asterisks
	assert.Equal(t, []string{"0", "1", "2", "3", "4", "5"}, parseTest(t, "*", 0, 5))

	// range
	assert.Equal(t, []string{"0", "1", "2", "3", "4", "5"}, parseTest(t, "0-5", 0, 5)) // day of week
}

func TestNewExpression(t *testing.T) {
	expected := &CronExpression{
		Minutes:    []string{"0", "15", "30", "45"},
		Hours:      []string{"0"},
		DayOfMonth: []string{"1", "15"},
		Months:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		DayOfWeek:  []string{"1", "2", "3", "4", "5"},
		Command:    "/usr/bin/find arg1 arg2",
	}

	result, err := NewCronExpression("*/15 0 1,15 * 1-5 /usr/bin/find arg1 arg2")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	expectedStr := `minute		0 15 30 45
hour		0
day of month	1 15
month		1 2 3 4 5 6 7 8 9 10 11 12
day of week	1 2 3 4 5
command		/usr/bin/find arg1 arg2`
	assert.Equal(t, expectedStr, result.String())
}

func TestBasicExpressionFailure(t *testing.T) {
	_, err := NewCronExpression("")
	assert.Error(t, err)

	_, err = NewCronExpression("bla1 bla2 bla3 blah4 blah5 blah6 blah7")
	assert.Error(t, err)
}

func TestBasicParseFailure(t *testing.T) {
	_, err := parse("bla1", 4, 5)
	assert.Error(t, err)

	_, err = parse("2", 4, 5)
	assert.Error(t, err)

	_, err = parse("-1", 0, 5)
	assert.Error(t, err)

	_, err = parse("2,5,7", 4, 5)
	assert.Error(t, err)

	_, err = parse("2-7", 4, 5)
	assert.Error(t, err)

	_, err = parse("4-7", 4, 5)
	assert.Error(t, err)

	_, err = parse("3-2", 1, 7)
	assert.Error(t, err)

	_, err = parse("3-2-1", 1, 7)
	assert.Error(t, err)

	_, err = parse("1-F", 1, 7)
	assert.Error(t, err)

	_, err = parse("M-5", 1, 7)
	assert.Error(t, err)

	_, err = parse("*/", 1, 7)
	assert.Error(t, err)

	_, err = parse("*,*", 1, 7)
	assert.Error(t, err)

}

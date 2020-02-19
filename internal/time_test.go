package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestampByMaxTime(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		output int64
	}{
		"success": {output: (1 << 63) - 1},
	}

	for _, t := range tests {
		assert.Equal(t.output, TimestampByMaxTime())
	}
}

func TestStringToTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	tests := map[string]struct {
		input  string
		output time.Time
	}{
		"success": {input: s, output: now},
	}

	for _, t := range tests {
		result := StringToTime(t.input)
		assert.Equal(t.output.Year(), result.Year())
		assert.Equal(t.output.Month(), result.Month())
		assert.Equal(result.Day(), result.Day())
		assert.Equal(result.Hour(), result.Hour())
		assert.Equal(result.Minute(), result.Minute())
		assert.Equal(result.Second(), result.Second())
	}
}

func TestTimeToString(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	tests := map[string]struct {
		input  time.Time
		output string
	}{
		"success": {input: now, output: s},
	}

	for _, t := range tests {
		assert.Equal(t.output, TimeToString(t.input))
	}
}

func TestYearlyStringToTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d", time.Now().Year())

	tests := map[string]struct {
		input  string
		output time.Time
	}{
		"success": {input: s, output: now},
	}

	for _, t := range tests {
		result := YearlyStringToTime(t.input)
		assert.Equal(t.output.Year(), result.Year())
		assert.Equal(1, int(result.Month()))
		assert.Equal(1, result.Day())
		assert.Equal(0, result.Hour())
		assert.Equal(0, result.Minute())
		assert.Equal(0, result.Second())
	}
}

func TestTimeToYearlyStringFormat(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d", time.Now().Year())

	tests := map[string]struct {
		output string
		input  time.Time
	}{
		"success": {output: s, input: now},
	}

	for _, t := range tests {
		result := TimeToYearlyStringFormat(t.input)
		assert.Equal(t.output, result)
	}
}

func TestMonthlyStringToTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d%02d", now.Year(), now.Month())

	tests := map[string]struct {
		input  string
		output time.Time
	}{
		"success": {input: s, output: now},
	}

	for _, t := range tests {
		result := MonthlyStringToTime(t.input)
		assert.Equal(t.output.Year(), result.Year())
		assert.Equal(t.output.Month(), result.Month())
		assert.Equal(1, result.Day())
		assert.Equal(0, result.Hour())
		assert.Equal(0, result.Minute())
		assert.Equal(0, result.Second())
	}
}

func TestTimeToMonthlyStringFormat(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d%02d", now.Year(), now.Month())
	tests := map[string]struct {
		output string
		input  time.Time
	}{
		"success": {output: s, input: now},
	}

	for _, t := range tests {
		result := TimeToMonthlyStringFormat(t.input)
		assert.Equal(t.output, result)
	}
}

func TestDailyStringToTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())

	tests := map[string]struct {
		input  string
		output time.Time
	}{
		"success": {input: s, output: now},
	}

	for _, t := range tests {
		result := DailyStringToTime(t.input)
		assert.Equal(t.output.Year(), result.Year())
		assert.Equal(t.output.Month(), result.Month())
		assert.Equal(t.output.Day(), result.Day())
		assert.Equal(0, result.Hour())
		assert.Equal(0, result.Minute())
		assert.Equal(0, result.Second())
	}

}

func TestTimeToDailyStringFormat(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())
	tests := map[string]struct {
		output string
		input  time.Time
	}{
		"success": {output: s, input: now},
	}

	for _, t := range tests {
		result := TimeToDailyStringFormat(t.input)
		assert.Equal(t.output, result)
	}

}

func TestHourlyStringToTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	s := fmt.Sprintf("%d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour())

	tests := map[string]struct {
		input  string
		output time.Time
	}{
		"success": {input: s, output: now},
	}

	for _, t := range tests {
		result := HourlyStringToTime(t.input)
		assert.Equal(t.output.Year(), result.Year())
		assert.Equal(t.output.Month(), result.Month())
		assert.Equal(t.output.Day(), result.Day())
		assert.Equal(t.output.Hour(), result.Hour())
		assert.Equal(0, result.Minute())
		assert.Equal(0, result.Second())
	}
}

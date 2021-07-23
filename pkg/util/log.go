package util

import (
	"bytes"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var (
	// ColorStatus returns a new function that returns status-colorized (cyan) strings for the
	// given arguments with fmt.Sprint().
	ColorStatus = color.New(color.FgCyan).SprintFunc()

	// ColorWarn returns a new function that returns status-colorized (yellow) strings for the
	// given arguments with fmt.Sprint().
	ColorWarn = color.New(color.FgYellow).SprintFunc()

	// ColorError returns a new function that returns error-colorized (red) strings for the
	// given arguments with fmt.Sprint().
	ColorError = color.New(color.FgRed).SprintFunc()
)

// TextFormat lets use a custom text format.
type TextFormat struct {
	ShowInfoLevel   bool
	ShowTimestamp   bool
	TimestampFormat string
}

// NewTextFormat creates the default text formatter.
func NewTextFormat() *TextFormat {
	return &TextFormat{
		ShowInfoLevel:   false,
		ShowTimestamp:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// Format formats the log statement.
func (f *TextFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	level := strings.ToUpper(entry.Level.String())
	switch level {
	case "INFO":
		if f.ShowInfoLevel {
			b.WriteString(ColorStatus(level))
			b.WriteString(": ")
		}
	case "WARNING":
		b.WriteString(ColorWarn(level))
		b.WriteString(": ")
	case "DEBUG":
		b.WriteString(ColorStatus(level))
		b.WriteString(": ")
	default:
		b.WriteString(ColorError(level))
		b.WriteString(": ")
	}
	if f.ShowTimestamp {
		b.WriteString(entry.Time.Format(f.TimestampFormat))
		b.WriteString(" - ")
	}

	b.WriteString(entry.Message)

	if !strings.HasSuffix(entry.Message, "\n") {
		b.WriteByte('\n')
	}
	return b.Bytes(), nil
}

package utils

import (
	"fmt"
	"time"

	"fyc.com/sprs/models"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func IsNullString(v models.NullString, prefix string, sufix string) string {
	s := string(v)

	if s == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", prefix, string(s), sufix)
}

func FormatLongDate(d string) string {
	const (
		RFC3339  = "2006-01-02"
		layoutID = "02 Jan 2006"
	)

	t, _ := time.Parse(RFC3339, d[0:10])
	return t.Format(layoutID)
}

func FormatShortDate(d string) string {
	const (
		RFC3339  = "2006-01-02"
		layoutID = "02-01-2006"
	)

	t, _ := time.Parse(RFC3339, d[0:10])
	return t.Format(layoutID)
}

func TermString(v int) string {
	if v == 1 {
		return "Cash"
	}
	return "Transfer"
}

func FormatNumber(f float64) string {
	p := message.NewPrinter(language.Indonesian)
	s := p.Sprintf("%0.f", f)
	return s
}

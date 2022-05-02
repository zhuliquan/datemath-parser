package datemath_parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DateMathParserOption func(*DateMathParser) error

func WithFormat(formats []string) DateMathParserOption {
	return func(p *DateMathParser) error {
		var parserFormats = []string{}
		for _, format := range formats {
			if jodaFormats, ok := builtInFormat[format]; ok {
				parserFormats = append(parserFormats, jodaFormats...)
			} else {
				parserFormats = append(parserFormats, format)
			}
		}
		p.Formats = parserFormats
		return nil
	}
}

var TimeZoneOffset = regexp.MustCompile(`(\+|-)(\d{1,2}):(\d{1,2})`)

func WithTimeZone(timeZone string) DateMathParserOption {
	return func(p *DateMathParser) error {
		if loc, err := time.LoadLocation(timeZone); err != nil {
			var timeOffset = timeZone
			if t, ok := builtInTimeZone[strings.ToUpper(timeZone)]; ok {
				timeOffset = t
			}
			var s = TimeZoneOffset.FindStringSubmatch(timeOffset)
			if len(s) != 4 {
				return fmt.Errorf("time zone: %s format is invalid, expect time offset format: (\\+|-)(\\d{1,2}):(\\d{1,2}) or time zone (abbreviation/full name) or IANA format", timeZone)
			} else {
				var flag = 1
				if s[1] == "-" {
					flag = -1
				}
				var hour, _ = strconv.Atoi(s[2])
				if hour > 23 || hour < 0 {
					return fmt.Errorf("time zone: %s is invalid, hour is out of range [0, 23]", timeZone)
				}
				var minute, _ = strconv.Atoi(s[3])
				if minute > 59 || minute < 0 {
					return fmt.Errorf("time zone: %s is invalid, minute is out of range [0, 59]", timeZone)
				}
				p.TimeZone = time.FixedZone("UTC",
					flag*(hour*int(time.Hour)+minute*int(time.Minute))/int(time.Second))
				return nil
			}
		} else {
			p.TimeZone = loc
			return nil
		}
	}
}

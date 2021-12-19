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
		var resF = []string{}
		for _, format := range formats {
			if builtInFormat[format] != nil {
				resF = append(resF, builtInFormat[format]...)
			} else {
				resF = append(resF, format)
			}
		}
		p.Formats = resF
		return nil
	}
}

var TimeZoneOffset, _ = regexp.Compile("(\\+|-)(\\d+):(\\d+)")

func WithTimeZone(timeZone string) DateMathParserOption {
	return func(p *DateMathParser) error {
		if loc, err := time.LoadLocation(timeZone); err != nil {
			timeZone = strings.ToUpper(timeZone)
			if builtInTimeZone[timeZone] != "" {
				timeZone = builtInTimeZone[timeZone]
			}
			var s = TimeZoneOffset.FindStringSubmatch(timeZone)
			if len(s) != 4 {
				return fmt.Errorf("time zone: %s is invalid", timeZone)
			} else {
				var offset = 1
				if s[1] == "-" {
					offset = -1
				}
				var hour, _ = strconv.Atoi(s[2])
				var minute, _ = strconv.Atoi(s[3])
				offset = offset * (hour*int(time.Hour) + minute*int(time.Minute)) / int(time.Second)
				p.TimeZone = time.FixedZone("UTC", offset)
				return nil
			}
		} else {
			p.TimeZone = loc
			return nil
		}
	}
}

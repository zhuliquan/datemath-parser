package datemath_parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DateMathParserOption func(*DateMathParser) error

func changeTimeFormat(format string) (string, error) {
	if strings.Index(format, "ww") != -1 ||
		strings.Index(format, "xxxx") != -1 ||
		strings.Index(format, "DDD") != -1 {
		return "", fmt.Errorf("doesn't support xxxx/ww/DDD")
	}
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "YYYY", "2006")
	format = strings.ReplaceAll(format, "yy", "06")
	format = strings.ReplaceAll(format, "YY", "06")
	format = strings.ReplaceAll(format, "MM", "05")
	format = strings.ReplaceAll(format, "dd", "04")
	format = strings.ReplaceAll(format, "HH", "03")
	format = strings.ReplaceAll(format, "mm", "02")
	format = strings.ReplaceAll(format, "ss", "01")
	format = strings.ReplaceAll(format, "e", "Mon")
	format = strings.ReplaceAll(format, "am", "pm")
	format = strings.ReplaceAll(format, "AM", "PM")
	format = strings.ReplaceAll(format, ".SSSSSSSSS", ".000000000")
	format = strings.ReplaceAll(format, ".SSSSSSSS", ".00000000")
	format = strings.ReplaceAll(format, ".SSSSSSS", ".0000000")
	format = strings.ReplaceAll(format, ".SSSSSS", ".000000")
	format = strings.ReplaceAll(format, ".SSSSS", ".00000")
	format = strings.ReplaceAll(format, ".SSSS", ".0000")
	format = strings.ReplaceAll(format, ".SSS", ".000")
	format = strings.ReplaceAll(format, ".SS", ".00")
	format = strings.ReplaceAll(format, ".S", ".0")
	return format, nil
}

func WithFormat(formats []string) DateMathParserOption {
	return func(p *DateMathParser) error {
		var resF = []string{}
		for _, format := range formats {
			if builtInFormat[format] != nil {
				for _, bf := range builtInFormat[format] {
					if bf, err := changeTimeFormat(bf); err == nil {
						resF = append(resF, bf)
					}
				}
			} else {
				if format, err := changeTimeFormat(format); err == nil {
					resF = append(resF, format)
				}
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

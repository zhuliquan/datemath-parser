package datemath_parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/vjeantet/jodaTime"
)

var durRegexp = regexp.MustCompile(`([\+-]\d*|\/)(y|M|w|d|h|H|m|s)`)

var emptyTime = time.Unix(0, 0)

var units = map[string]time.Duration{
	"y": time.Hour * 365 * 24,
	"M": time.Hour * 30 * 24,
	"w": time.Hour * 7 * 24,
	"d": time.Hour * 24,
	"h": time.Hour,
	"H": time.Hour,
	"m": time.Minute,
	"s": time.Second,
}

type DateMathParser struct {
	Formats  []string
	TimeZone *time.Location
}

func NewDateMathParser(opts ...DateMathParserOption) (*DateMathParser, error) {
	var p = &DateMathParser{}
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (p *DateMathParser) Parse(expr string) (time.Time, error) {
	var res time.Time
	var dur = ""
	if len(expr) >= 3 && expr[0:3] == "now" {
		dur = expr[3:]
		res = time.Now()
	} else {
		var sep = strings.Index(expr, "||")
		var err error
		if sep == -1 {
			dur = ""
			if res, err = p.parseTime(expr); err != nil {
				return emptyTime, err
			}
		} else {
			dur = expr[sep+2:]
			if res, err = p.parseTime(expr[:sep]); err != nil {
				return emptyTime, err
			}
		}
	}
	res = res.UTC()
	if dur == "" {
		return res, nil
	} else {
		return p.evalDur(dur, res)
	}

}

func (p *DateMathParser) parseTime(expr string) (time.Time, error) {
	if len(p.Formats) != 0 {
		for _, format := range p.Formats {
			if format == "epoch_second" {
				if sec, err := strconv.ParseInt(expr, 10, 64); err != nil {
					return emptyTime, fmt.Errorf("failed to parse time, expr: %s, format: epoch_second", expr)
				} else {
					return time.Unix(sec, 0), nil
				}
			} else if format == "epoch_millis" {
				if millis, err := strconv.ParseInt(expr, 10, 64); err != nil {
					return emptyTime, fmt.Errorf("failed to parse time, expr: %s, format: epoch_millis", expr)
				} else {
					return time.Unix(millis/1000, millis%1000), nil
				}
			} else {
				if tim, err := p.parseFormat(expr, format); err == nil {
					return tim, nil
				}
			}
		}
		return emptyTime, fmt.Errorf("failed to parse time, expr: %s, format: %+v", expr, p.Formats)
	} else {
		return p.parseAny(expr)
	}
}

func (p *DateMathParser) parseFormat(expr, format string) (time.Time, error) {
	if p.TimeZone != nil {
		return jodaTime.ParseInLocation(format, expr, p.TimeZone.String())
	} else {
		return jodaTime.Parse(format, expr)
	}
}

func (p *DateMathParser) parseAny(expr string) (time.Time, error) {
	if p.TimeZone != nil {
		return dateparse.ParseIn(expr, p.TimeZone)
	}
	return dateparse.ParseAny(expr)
}

func (p *DateMathParser) evalDur(dur string, tim time.Time) (time.Time, error) {
	var res = tim
	var allMatch = durRegexp.FindAllStringSubmatch(dur, -1)
	if len(allMatch) == 0 {
		return emptyTime, fmt.Errorf(`expect match expression: ([\+-]\d*|\/)(y|M|w|d|h|H|m|s)`)
	}
	for _, s := range allMatch {
		if s[1] == "/" {
			res = res.Round(units[s[2]])
		} else {
			var d int
			if s[1] == "-" {
				d = -1
			} else if s[1] == "+" {
				d = 1
			} else {
				d, _ = strconv.Atoi(s[1])
			}
			res = res.Add(time.Duration(d) * units[s[2]])
		}
	}
	return res, nil
}

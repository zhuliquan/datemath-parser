package datemath_parser

import (
	"testing"
	"time"
)

func TestDateMathParser_evalDur(t *testing.T) {
	var p = &DateMathParser{}
	type testCase struct {
		name string
		in   string
		out  int64
		err  error
	}

	for _, eachCase := range []testCase{
		{
			name: "TestEvalDur01",
			in:   "+7y",
			out:  int64(time.Hour*365*24*7) / 1000,
			err:  nil,
		},
		{
			name: "TestEvalDur02",
			in:   "+y",
			out:  int64(time.Hour*365*24) / 1000,
			err:  nil,
		},
		{
			name: "TestEvalDur03",
			in:   "-y",
			out:  -int64(time.Hour*365*24) / 1000,
			err:  nil,
		},
		{
			name: "TestEvalDur04",
			in:   "-7y",
			out:  -int64(time.Hour*365*24*7) / 1000,
			err:  nil,
		},
		{
			name: "TestEvalDur05",
			in:   "+y+H+m+s/d",
			out:  int64(units["y"]) / 1000,
			err:  nil,
		},
		{
			name: "TestEvalDur06",
			in:   "+y+M+d+H/d",
			out:  int64(units["y"]+units["M"]+units["d"]) / 1000,
			err:  nil,
		},
	} {
		t.Run(eachCase.name, func(t *testing.T) {
			if o, e := p.evalDur(eachCase.in, time.Unix(0, 0)); e != eachCase.err {
				t.Errorf("expect get err: %+v, but get err: %+v", eachCase.err, e)
			} else if o != eachCase.out {
				t.Errorf("expect get res: %+v, but get res: %+v", eachCase.out, o)
			}
		})
	}
}

func TestDateMathParser_timeZone(t *testing.T) {
	type testCase struct {
		name string
		in   string
		out  int64
		err  error
	}

	for _, eachCase := range []testCase{
		{
			name: "TestTimeZone01",
			in:   "+7:00",
			out:  -int64(time.Hour*7) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone02",
			in:   "+7:45",
			out:  -int64(time.Hour*7+time.Minute*45) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone03",
			in:   "-7:00",
			out:  int64(time.Hour*7) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone04",
			in:   "-7:45",
			out:  int64(time.Hour*7+time.Minute*45) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone05",
			in:   "PST",
			out:  int64(time.Hour*8) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone06",
			in:   "Europe/Malta",
			out:  -int64(time.Hour) / int64(time.Second),
			err:  nil,
		},
		{
			name: "TestTimeZone07",
			in:   "Asia/Shanghai",
			out:  -int64(time.Hour*8) / int64(time.Second),
			err:  nil,
		},
	} {
		t.Run(eachCase.name, func(t *testing.T) {
			p, err := NewDateMathParser(WithTimeZone(eachCase.in))
			if err != nil {
				t.Errorf("failed to generate date math parser, err: %+v", err)
			}
			if o, e := p.parseTime("1970-01-01 00:00:00"); e != eachCase.err {
				t.Errorf("expect get err: %+v, but get err: %+v", eachCase.err, e)
			} else if o.Unix() != eachCase.out {
				t.Errorf("expect get res: %+v, but get res: %+v", eachCase.out, o.Unix())
			}
		})
	}
}

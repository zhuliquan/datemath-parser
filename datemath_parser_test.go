package datemath_parser

import (
	"fmt"
	"reflect"
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
			} else if o.UnixNano()/1000 != eachCase.out {
				t.Errorf("expect get res: %+v, but get res: %+v", eachCase.out, o)
			}
		})
	}
}

func TestDateMathParser_timeZone(t *testing.T) {
	type testCase struct {
		name   string
		in     string
		newErr error
		out    int64
		err    error
	}

	for _, eachCase := range []testCase{
		{
			name:   "TestTimeZone01",
			in:     "+7:00",
			newErr: nil,
			out:    -int64(time.Hour*7) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone02",
			in:     "+7:45",
			newErr: nil,
			out:    -int64(time.Hour*7+time.Minute*45) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone03",
			in:     "-7:00",
			newErr: nil,
			out:    int64(time.Hour*7) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone04",
			in:     "-7:45",
			newErr: nil,
			out:    int64(time.Hour*7+time.Minute*45) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone05",
			in:     "PST",
			newErr: nil,
			out:    int64(time.Hour*8) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone06",
			in:     "Europe/Malta",
			newErr: nil,
			out:    -int64(time.Hour) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "TestTimeZone07",
			in:     "Asia/Shanghai",
			newErr: nil,
			out:    -int64(time.Hour*8) / int64(time.Second),
			err:    nil,
		},
		{
			name:   "test_wrong_time_zone_01",
			in:     "+45:00",
			newErr: fmt.Errorf("time zone: +45:00 is invalid, hour is out of range [0, 23]"),
		},
		{
			name:   "test_wrong_time_zone_02",
			in:     "+05:89",
			newErr: fmt.Errorf("time zone: +05:89 is invalid, minute is out of range [0, 59]"),
		},
		{
			name:   "test_wrong_time_zone_03",
			in:     "Asia/Shanghai08",
			newErr: fmt.Errorf("time zone: Asia/Shanghai08 format is invalid, expect time offset format: (\\+|-)(\\d{1,2}):(\\d{1,2}) or time zone (abbreviation/full name) or IANA format"),
		},
	} {
		t.Run(eachCase.name, func(t *testing.T) {
			if p, err := NewDateMathParser(WithTimeZone(eachCase.in)); err != nil {
				if err.Error() != eachCase.newErr.Error() {
					t.Errorf("failed to generate date math parser, err: %+v, expect err: %+v", err, eachCase.newErr)
				}
			} else {
				if o, e := p.parseTime("1970-01-01 00:00:00"); e != eachCase.err {
					t.Errorf("expect get err: %+v, but get err: %+v", eachCase.err, e)
				} else if o.Unix() != eachCase.out {
					t.Errorf("expect get res: %+v, but get res: %+v", eachCase.out, o.Unix())
				}
			}
		})
	}
}

func TestDateMathParser_format(t *testing.T) {
	type testCase struct {
		name   string
		format []string
		parser *DateMathParser
	}

	for _, eachCase := range []testCase{
		{
			name: "test_format",
			format: []string{"epoch_millis", "epoch_second",
				"date_optional_time",
				"strict_date_optional_time_nanos",
				"basic_ordinal_date_time_no_millis",
				"yyyy-DDDTHH"},
			parser: &DateMathParser{
				Formats: []string{"epoch_millis", "epoch_second",
					"yyyy-MM-ddTHH:mm:ss.SSSZ", "yyyy-MM-dd",
					"yyyy-MM-ddTHH:mm:ss.SSSSSSZ", "yyyy-MM-dd",
					"yyyyDDDTHHmmssZ", "yyyy-DDDTHH"},
			},
		},
	} {
		t.Run(eachCase.name, func(t *testing.T) {
			var p, _ = NewDateMathParser(WithFormat(eachCase.format))
			if !reflect.DeepEqual(p.Formats, eachCase.parser.Formats) {
				t.Errorf("expect format: %v, but got: %v", eachCase.parser.Formats, p.Formats)
			}
		})
	}
}

func TestDateMathParser_parseTime(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		p       *DateMathParser
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "TestDateMathParser_parseTime01",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392"},
			want:    time.Unix(1640183392, 0),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_parseTime02",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392001"},
			want:    time.Unix(1640183392, 1000000),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_parseTime03",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22"},
			want:    time.Unix(1640102400, 0),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_parseTime04",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22T10:09:00"},
			want:    time.Unix(1640138940, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.parseTime(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateMathParser.parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateMathParser.parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateMathParser_Parse(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		p       *DateMathParser
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "TestDateMathParser_Parse01",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392||+h/d"},
			want:    time.Unix(1640183392+3600, 0).UTC().Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse01_01",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392||+2d/d"},
			want:    time.Unix(1640183392+(2*24*3600), 0).UTC().Round(time.Hour * 24),
			wantErr: false,
		},

		{
			name:    "TestDateMathParser_Parse02",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392001||+h/d"},
			want:    time.Unix(1640183392+3600, 1000000).UTC().Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse03",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22||+h+M/d"},
			want:    time.Unix(1640102400+3600+3600*24*30, 0).UTC().Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse04",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22T10:09:00||+h+M/d"},
			want:    time.Unix(1640138940+3600+3600*24*30, 0).UTC().Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse05",
			p:       &DateMathParser{TimeZone: time.UTC},
			args:    args{expr: "now/s"},
			want:    time.Now().UTC().Round(time.Second),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse06",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22T10:09:00"},
			want:    time.Unix(1640138940, 0).UTC(),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse07",
			p:       &DateMathParser{Formats: []string{"epoch_millis"}, TimeZone: time.Local},
			args:    args{expr: "1640138940000"},
			want:    time.Unix(1640138940, 0).UTC(),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse07_01",
			p:       &DateMathParser{Formats: []string{"epoch_millis"}, TimeZone: time.Local},
			args:    args{expr: "xx1640138940000"},
			want:    emptyTime,
			wantErr: true,
		},
		{
			name:    "TestDateMathParser_Parse08",
			p:       &DateMathParser{Formats: []string{"epoch_second"}, TimeZone: time.Local},
			args:    args{expr: "1640138940"},
			want:    time.Unix(1640138940, 0).UTC(),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse08_01",
			p:       &DateMathParser{Formats: []string{"epoch_second"}, TimeZone: time.Local},
			args:    args{expr: "xx1640138940000"},
			want:    emptyTime,
			wantErr: true,
		},
		{
			name:    "TestDateMathParser_Parse09",
			p:       &DateMathParser{Formats: []string{"yyyy-MM-dd"}, TimeZone: time.Local},
			args:    args{expr: "1900-02-29"},
			want:    emptyTime,
			wantErr: true,
		},
		{
			name:    "TestDateMathParser_Parse10",
			p:       &DateMathParser{Formats: []string{"yyyy-MM-dd"}, TimeZone: time.Local},
			args:    args{expr: "1900-02-28"},
			want:    time.Date(1900, 02, 28, 0, 0, 0, 0, time.Local).UTC(),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse11",
			p:       &DateMathParser{Formats: []string{"yyyy-MM-dd"}},
			args:    args{expr: "1900-02-28"},
			want:    time.Date(1900, 02, 28, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse12",
			p:       &DateMathParser{Formats: []string{"yyyy-MM-dd"}},
			args:    args{expr: "1900-02-29||+y"},
			want:    emptyTime,
			wantErr: true,
		},
		{
			name:    "TestDateMathParser_Parse13",
			p:       &DateMathParser{Formats: []string{"yyyy-MM-dd"}},
			args:    args{expr: "1900-02-28||+x"},
			want:    emptyTime,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Parse(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateMathParser.parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateMathParser.parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

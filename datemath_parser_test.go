package datemath_parser

import (
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
			want:    time.Unix(1640183392+3600, 0).Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse02",
			p:       &DateMathParser{},
			args:    args{expr: "1640183392001||+h/d"},
			want:    time.Unix(1640183392+3600, 1000000).Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse03",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22||+h+M/d"},
			want:    time.Unix(1640102400+3600+3600*24*30, 0).Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse04",
			p:       &DateMathParser{TimeZone: time.Local},
			args:    args{expr: "2021-12-22T10:09:00||+h+M/d"},
			want:    time.Unix(1640138940+3600+3600*24*30, 0).Round(time.Hour * 24),
			wantErr: false,
		},
		{
			name:    "TestDateMathParser_Parse05",
			p:       &DateMathParser{TimeZone: time.UTC},
			args:    args{expr: "now/s"},
			want:    time.Now().UTC().Round(time.Second),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Parse(tt.args.expr)
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

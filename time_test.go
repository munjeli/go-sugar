package gosugar

import (
	"reflect"
	"testing"
	"time"
)

func TestDateDuration_String(t *testing.T) {
	type fields struct {
		Years   int
		Months  int
		Days    int
		Hours   int
		Minutes int
		Seconds int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "cast to string for query",
			fields: fields{
				Years:   0,
				Months:  0,
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
			want: "0y0M0d0h0m0s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DateDuration{
				Years:   tt.fields.Years,
				Months:  tt.fields.Months,
				Days:    tt.fields.Days,
				Hours:   tt.fields.Hours,
				Minutes: tt.fields.Minutes,
				Seconds: tt.fields.Seconds,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateDuration_ToTime(t *testing.T) {
	tests := []struct {
		name   string
		dd DateDuration
		want   time.Time
	}{
		{
			name: "empty dd",
			dd: DateDuration{
				Years:   0,
				Months:  0,
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
			want: time.Time{},
		},
		{
			name: "set date by ints, still begin time",
			dd: DateDuration{
				Years:   1,
				Months:  1,
				Days:    1,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dd.ToTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateDuration_Truncate(t *testing.T) {
	type fields struct {
		Years   int
		Months  int
		Days    int
		Hours   int
		Minutes int
		Seconds int
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   DateDuration
	}{
		{
			name: "round to days",
			fields: fields{
				Years:   4,
				Months:  0,
				Days:    1,
				Hours:   9,
				Minutes: 3,
				Seconds: 1,
			},
			args: args{s:"d"},
			want: DateDuration{
				Years:   4,
				Months:  0,
				Days:    1,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
		},
		{
			name: "bad key for split, returns og object",
			fields: fields{
				Years:   4,
				Months:  0,
				Days:    1,
				Hours:   9,
				Minutes: 3,
				Seconds: 1,
			},
			args: args{s:"k"},
			want: DateDuration{
				Years:   4,
				Months:  0,
				Days:    1,
				Hours:   9,
				Minutes: 3,
				Seconds: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DateDuration{
				Years:   tt.fields.Years,
				Months:  tt.fields.Months,
				Days:    tt.fields.Days,
				Hours:   tt.fields.Hours,
				Minutes: tt.fields.Minutes,
				Seconds: tt.fields.Seconds,
			}
			if got := d.Truncate(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDateDuration(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantD DateDuration
	}{
		{
			name: "empty string",
			args: args{
				s: "",
			},
			wantD: DateDuration{
				Years:   0,
				Months:  0,
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
		},
		{
			name: "valid string",
			args: args{
				s: "1y2M2d1h2m3s",
			},
			wantD: DateDuration{
				Years:   1,
				Months:  2,
				Days:    2,
				Hours:   1,
				Minutes: 2,
				Seconds: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := ParseDateDuration(tt.args.s); !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("ParseDateDuration() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestSumDurations(t *testing.T) {
	type args struct {
		ds []time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDurations(tt.args.ds); got != tt.want {
				t.Errorf("SumDurations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeAddDateDuration(t *testing.T) {
	type args struct {
		t time.Time
		d DateDuration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeAddDateDuration(tt.args.t, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeAddDateDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSubDateDuration(t *testing.T) {
	type args struct {
		t time.Time
		d DateDuration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeSubDateDuration(tt.args.t, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSubDateDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSubDuration(t *testing.T) {
	type args struct {
		t time.Time
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeSubDuration(tt.args.t, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSubDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

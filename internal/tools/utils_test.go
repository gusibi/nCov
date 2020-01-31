package tools

import "testing"

func TestTimeStampToBJString(t *testing.T) {
	type args struct {
		ts     int64
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "ts to string",
			args: args{
				1580350915,
				"2006-01-02 15:04:05",
			},
			want: "2020-01-30 10:21:55",
		},
		{
			name: "ts to bj string",
			args: args{
				1580383920,
				"2006-01-02 15:04:05",
			},
			want: "2020-01-30 19:32:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BjTimeStampToBJString(tt.args.ts, tt.args.layout); got != tt.want {
				t.Errorf("BjTimeStampToBJString() = %v, want %v", got, tt.want)
			}
		})
	}
}

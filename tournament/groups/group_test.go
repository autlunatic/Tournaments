package groups

import "testing"

func Test_getCompetitorCount(t *testing.T) {
	type args struct {
		in0 []G
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"simple test", args{generate8CompetitorsIn3Groups().grps}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCompetitorCount(tt.args.in0); got != tt.want {
				t.Errorf("getCompetitorCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

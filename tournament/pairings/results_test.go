package pairings

import (
	"reflect"
	"testing"
)

func TestResult_AddPoints(t *testing.T) {
	type fields struct {
		GamePoints1 int
		GamePoints2 int
	}
	type pts struct {
		p1 int
		p2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   pts
		want   pts
	}{
		{"simple test that points are set to correct competitor", fields{1, 2}, pts{3, 4}, pts{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{
				GamePoints1: tt.fields.GamePoints1,
				GamePoints2: tt.fields.GamePoints2,
			}
			r.SetPoints(tt.args.p1, tt.args.p2)
			if !reflect.DeepEqual(r.GamePoints1, tt.want.p1) {
				t.Errorf("CalcPairingsForFinals() = %v, want %v", r.GamePoints1, tt.want.p2)
			}
			if !reflect.DeepEqual(r.GamePoints2, tt.want.p2) {
				t.Errorf("CalcPairingsForFinals() = %v, want %v", r.GamePoints2, tt.want.p2)
			}
		})
	}
}

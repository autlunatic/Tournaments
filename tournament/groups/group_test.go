package groups

import (
	"reflect"
	"testing"
)

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

var grpsForTestGOfCompetitorID = generate8CompetitorsIn3Groups().grps

func TestGOfCompetitorID(t *testing.T) {
	type args struct {
		gs           []G
		competitorID int
	}
	tests := []struct {
		name    string
		args    args
		want    G
		wantErr bool
	}{
		{"simple test no error", args{generate8CompetitorsIn3Groups().grps, 0}, grpsForTestGOfCompetitorID[0], false},
		{"simple test with error", args{generate8CompetitorsIn3Groups().grps, 10}, G{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GOfCompentitorID(tt.args.gs, tt.args.competitorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GOfCompentitorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GOfCompentitorID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroupIDOfCompetitor(t *testing.T) {
	type args struct {
		gs           []G
		competitorID int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"simple test no error", args{generate8CompetitorsIn3Groups().grps, 0}, 0, false},
		{"simple test with error", args{generate8CompetitorsIn3Groups().grps, 10}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroupIDOfCompetitor(tt.args.gs, tt.args.competitorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupIDOfCompetitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGroupIDOfCompetitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

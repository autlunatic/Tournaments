package competitors

import (
	"reflect"
	"testing"
)

var testCompetitors []Competitor

func wantedSlice() []Competitor {
	c := make([]Competitor, 4)
	c[0] = GetCompetitor(testCompetitors, 1)
	c[1] = GetCompetitor(testCompetitors, 2)
	c[2] = GetCompetitor(testCompetitors, 3)
	c[3] = GetCompetitor(testCompetitors, 0)
	return c
}

func TestGetCompetitorsSortedByGroupPoints(t *testing.T) {
	testCompetitors = NewTestCompetitors(4)
	GetCompetitor(testCompetitors, 0).AddPoints(1)
	GetCompetitor(testCompetitors, 1).AddPoints(10)
	GetCompetitor(testCompetitors, 2).AddPoints(9)
	GetCompetitor(testCompetitors, 3).AddPoints(2)

	tests := []struct {
		name string
		want []Competitor
	}{
		{"sort", wantedSlice()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompetitorsSortedByGroupPoints(testCompetitors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitorsSortedByGroupPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		competitorID int
	}
	tests := []struct {
		name            string
		arg             args
		want            int
		competitorToAdd Competitor
	}{
		{name: "Competitor not in items", arg: args{competitorID: 9}, want: 0, competitorToAdd: nil},
		{name: "Competitor one in items", arg: args{competitorID: 3}, want: 1, competitorToAdd: nil},
		{"Competitor two in items", args{competitorID: 1}, 2, NewCompetitor("1", 1)},
	}
	for _, tt := range tests {
		c := NewTestCompetitors(5)
		if tt.competitorToAdd != nil {
			c = append(c, tt.competitorToAdd)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(c, tt.arg.competitorID); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		arg     Competitor
		wantErr bool
	}{
		{"simple add new one", NewCompetitor("Number 6", 6), false},
		{"add one with already given ID", NewCompetitor("Number 1", 1), true},
	}
	for _, tt := range tests {
		c := NewTestCompetitors(5)
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Add(c, tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimpleCompetitor_AddPoints(t *testing.T) {
	type fields struct {
		id          int
		name        string
		GroupPoints int
		drawNumber  int64
	}
	tests := []struct {
		name   string
		fields fields
		p      int
	}{
		{"adding points", fields{1, "Neo", 5, 0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SimpleCompetitor{
				id:          tt.fields.id,
				name:        tt.fields.name,
				GroupPoints: tt.fields.GroupPoints,
				drawNumber:  tt.fields.drawNumber,
			}
			c.AddPoints(tt.p)
			if c.GetPoints() != 8 {
				t.Errorf("points were not added or not read after adding: have %v, want %v", c.GetPoints(), 8)
			}
		})
	}
}

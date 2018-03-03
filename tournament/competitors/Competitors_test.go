package competitors

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	type args struct {
		competitorID int
	}
	tests := []struct {
		name            string
		arg             args
		want            int
		competitorToAdd C
	}{
		{name: "Competitor not in items", arg: args{competitorID: 9}, want: 0, competitorToAdd: nil},
		{name: "Competitor one in items", arg: args{competitorID: 3}, want: 1, competitorToAdd: nil},
		{"Competitor two in items", args{competitorID: 1}, 2, New("1", 1)},
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
		arg     C
		wantErr bool
	}{
		{"simple add new one", New("Number 6", 6), false},
		{"add one with already given ID", New("Number 1", 1), true},
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
		id         int
		name       string
		Points     int
		drawNumber int64
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
				GroupPoints: tt.fields.Points,
				drawNumber:  tt.fields.drawNumber,
				GamePoints:  tt.fields.Points,
			}
			c.AddPoints(tt.p)
			if c.GetPoints() != 8 {
				t.Errorf("points were not added or not read after adding: have %v, want %v", c.GetPoints(), 8)
			}
			c.AddGamePoints(tt.p)
			if c.GetGamePoints() != 8 {
				t.Errorf("gamepoints were not added or not read after adding: have %v, want %v", c.GetGamePoints(), 8)
			}
		})
	}
}

func TestGetCompetitor(t *testing.T) {
	cs := NewTestCompetitors(5)
	tests := []struct {
		name string
		ID   int
		want C
	}{
		{"Get competitor", 1, cs[1]},
		{"invalid Competitor", 9, &EmptyCompetitor{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompetitor(cs, tt.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompetitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearPoints(t *testing.T) {
	cs := NewTestCompetitors(9)
	for i := range cs {
		cs[i].AddPoints(i + 1)
	}
	for i := range cs {
		if cs[i].GetPoints() == 0 {
			t.Error("no points should be Zero")
		}
		cs[i].AddGamePoints(i + 1)
	}
	ClearPoints(cs)
	for i := range cs {
		if cs[i].GetPoints() != 0 {
			t.Error("all points should be Zero after clear")
		}
		if cs[i].GetGamePoints() != 0 {
			t.Error("all points should be Zero after clear")
		}

	}

}

func getWantedTestAddByName1() []C {
	var out = NewTestCompetitors(5)
	out = append(out, New("Hans", 5))
	return out
}
func TestAddByName(t *testing.T) {
	type args struct {
		cs   []C
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []C
		wantErr bool
	}{
		{"add not taken Name", args{NewTestCompetitors(5), "Hans"}, getWantedTestAddByName1(), false},
		{"add Name", args{NewTestCompetitors(5), "Benni"}, NewTestCompetitors(5), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddByName(tt.args.cs, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("AddByName() = %v, want %v", got, tt.want)
			}
			for i, cgot := range got {
				if cgot.Name() != tt.want[i].Name() {
					t.Errorf("AddByName() = %v, want %v", cgot.Name(), tt.want[i].Name())
				}
			}
			for i, cgot := range got {
				if cgot.ID() != tt.want[i].ID() {
					t.Errorf("AddByName() = %v, want %v", cgot.ID(), tt.want[i].ID())
				}
			}
		})
	}
}
func TestNameToID(t *testing.T) {
	type args struct {
		cs   []C
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not containing should return -1", args{NewTestCompetitors(5), "Hugo"}, -1},
		{"Benni should return 0", args{NewTestCompetitors(5), "Benni"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameToID(tt.args.cs, tt.args.name); got != tt.want {
				t.Errorf("ContainsName() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestContainsName(t *testing.T) {
	type args struct {
		cs   []C
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not containing should return false", args{NewTestCompetitors(5), "Hugo"}, false},
		{"containing should return true", args{NewTestCompetitors(5), "Benni"}, true},
		{"containing should ignore case ", args{NewTestCompetitors(5), "benni"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsName(tt.args.cs, tt.args.name); got != tt.want {
				t.Errorf("ContainsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxID(t *testing.T) {
	type args struct {
		cs []C
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"6 test competitors highest id should be 5, as ID begins at 0", args{NewTestCompetitors(6)}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxID(tt.args.cs); got != tt.want {
				t.Errorf("GetMaxID() = %v, want %v", got, tt.want)
			}
		})
	}
}

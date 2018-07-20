package groups

import (
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func TestCalcGroups_OneCompetitor_OneGroup(t *testing.T) {
	c := make([]competitors.C, 0)
	c = append(c,
		competitors.New("Benni", 0))

	result, err := CalcGroups(c, 1)
	TestingUtils.CheckEquals(0, len(result), "len of result should match groupcount", t)
	if err == nil {
		t.Error("Impossible groupcount should return error")
	}
}
func TestCalcGroups_6Competitor_2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(6)
	result, _ := CalcGroups(c, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].Competitors[0].Name() != "Benni" {
		t.Error("first Group first competitor should be Benni")
	}
}
func TestCalcGroups_7Competitor_2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(7)
	result, err := CalcGroups(c, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].Competitors[3].Name() != "Andrea" {
		t.Error("item[3] should be Andrea")
	}
	if result[1].Competitors[2].Name() != "Bernhard" {
		t.Error("second Group last competitor should be Bernhard")
	}
	if err != nil {
		t.Error("error should be Empty")
	}
}
func TestCalcGroups_23Competitor_7Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(23)
	result, _ := CalcGroups(c, 7)
	TestingUtils.CheckEquals(7, len(result), "len of result should match groupcount", t)
	if result[0].Competitors[3].Name() != "Andrea" {
		t.Error("item[3] should be Andrea")
	}
	if result[6].Competitors[2].Name() != "23" {
		t.Error("second Group last competitor should be 7")
	}
	TestingUtils.CheckEquals(7, result[6].ID, "Group id of last Group should be 7", t)
}

func TestCalcGroups_10Competitor_6Groups(t *testing.T) {
	// this test represents an impossible grouping
	c := competitors.NewTestCompetitors(10)
	result, err := CalcGroups(c, 7)
	if err == nil {
		t.Error("impossible groupcount error should be returned")
	} else {
		if err.Error() != "too many groups for this count of competitors!" {
			t.Error("WrongErrorMessage", err.Error())
		}
	}
	TestingUtils.CheckEquals(0, len(result), "len of result should be 0", t)
}

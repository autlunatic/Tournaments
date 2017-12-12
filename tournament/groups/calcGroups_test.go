package groups

import (
	"testing"

	"github.com/autlunatic/TestingUtils"
	"github.com/autlunatic/Tournaments/tournament/competitors"
)

func TestCalcGroups_OneCompetitor_OneGroup(t *testing.T) {
	c := competitors.Competitors{}
	c.Items = append(c.Items,
		competitors.NewCompetitor("Benni"))

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
	if result[0].competitors.Items[0].Name != "1" {
		t.Error("first Group first competitor should be Benni")
	}
}
func TestCalcGroups_7Competitor_2Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(7)
	result, err := CalcGroups(c, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].competitors.Items[3].Name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[1].competitors.Items[2].Name != "7" {
		t.Error("second Group last competitor should be 7")
	}
	if err != nil {
		t.Error("error should be Empty")
	}

}
func TestCalcGroups_23Competitor_7Groups(t *testing.T) {
	c := competitors.NewTestCompetitors(23)
	result, _ := CalcGroups(c, 7)
	TestingUtils.CheckEquals(7, len(result), "len of result should match groupcount", t)
	if result[0].competitors.Items[3].Name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[6].competitors.Items[2].Name != "23" {
		t.Error("second Group last competitor should be 7")
	}
	TestingUtils.CheckEquals(7, result[6].id, "Group id of last Group should be 7", t)
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
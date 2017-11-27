package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"testing"
)

func TestCalcGroups_OneCompetitor_OneGroup(t *testing.T) {
	competitors := Competitors{}
	competitors.items = append(competitors.items,
		newCompetitor("Benni"))

	result, err := calcGroups(competitors, 1)
	TestingUtils.CheckEquals(0, len(result), "len of result should match groupcount", t)
	if err == nil {
		t.Error("Impossible groupcount should return error")
	}
}
func TestCalcGroups_6Competitor_2Groups(t *testing.T) {
	competitors := newTestCompetitors(6)
	result, _ := calcGroups(competitors, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].competitors.items[0].name != "1" {
		t.Error("first group first competitor should be Benni")
	}
}
func TestCalcGroups_7Competitor_2Groups(t *testing.T) {
	competitors := newTestCompetitors(7)
	result,err := calcGroups(competitors, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].competitors.items[3].name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[1].competitors.items[2].name != "7" {
		t.Error("second group last competitor should be 7")
	}
	if err != nil {
		t.Error("error should be Empty")
	}

}
func TestCalcGroups_23Competitor_7Groups(t *testing.T) {
	competitors := newTestCompetitors(23)
	result,_ := calcGroups(competitors, 7)
	TestingUtils.CheckEquals(7, len(result), "len of result should match groupcount", t)
	if result[0].competitors.items[3].name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[6].competitors.items[2].name != "23" {
		t.Error("second group last competitor should be 7")
	}
	TestingUtils.CheckEquals(7, result[6].id, "group id of last group should be 7", t)
}

func TestCalcGroups_10Competitor_6Groups(t *testing.T) {
	// this test represents an impossible grouping
	competitors := newTestCompetitors(10)
	result,err := calcGroups(competitors, 7)
	if err == nil{
		t.Error("impossible groupcount error should be returned")
	}
	TestingUtils.CheckEquals(0, len(result), "len of result should be 0", t)
}

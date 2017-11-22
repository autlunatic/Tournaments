package tournament

import (
	"github.com/autlunatic/TestingUtils"
	"strconv"
	"testing"
)

func TestCalcGroups_OneCompetitor_OneGroup(t *testing.T) {
	competitors := Competitors{}
	competitors.items = append(competitors.items,
		newCompetitor("Benni"))

	result := calcGroups(competitors, 1)
	TestingUtils.CheckEquals(1, len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[0].name != "Benni" {
		t.Error("first group first competitor should be Benni")
	}
}
func TestCalcGroups_6Competitor_2Groups(t *testing.T) {
	competitors := Competitors{}
	for i := 1; i < 7; i++ {
		competitors.items = append(competitors.items, newCompetitor(strconv.Itoa(i)))
	}
	result := calcGroups(competitors, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[0].name != "1" {
		t.Error("first group first competitor should be Benni")
	}
}
func TestCalcGroups_7Competitor_2Groups(t *testing.T) {
	competitors := Competitors{}
	for i := 1; i < 8; i++ {
		competitors.items = append(competitors.items, newCompetitor(strconv.Itoa(i)))
	}
	result := calcGroups(competitors, 2)
	TestingUtils.CheckEquals(2, len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[3].name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[1].Competitors.items[2].name != "7" {
		t.Error("second group last competitor should be 7")
	}
}
func TestCalcGroups_23Competitor_7Groups(t *testing.T) {
	competitors := Competitors{}
	for i := 1; i < 24; i++ {
		competitors.items = append(competitors.items, newCompetitor(strconv.Itoa(i)))
	}
	result := calcGroups(competitors, 7)
	TestingUtils.CheckEquals(7, len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[3].name != "4" {
		t.Error("item[3] should be 4")
	}
	if result[6].Competitors.items[2].name != "23" {
		t.Error("second group last competitor should be 7")
	}
	TestingUtils.CheckEquals(7, result[6].Id, "group id of last group should be 7", t)

}

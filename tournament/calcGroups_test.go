package tournament

import (
	"testing"
	"github.com/autlunatic/TestingUtils"
	"strconv"
	"fmt"
)

func TestCalcGroups_OneCompetitor_OneGroup(t *testing.T) {
	competitors := Competitors{}
	competitors.items = append(competitors.items,
		Competitor{"Benni"})

	result := calcGroups(competitors, 1)
	TestingUtils.CheckEquals(1,len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[0].name != "Benni"{
		t.Error("first group first competitor should be Benni")
	}
}
func TestCalcGroups_6Competitor_2Groups(t *testing.T) {
	competitors := Competitors{}
	for i:=1;i<7;i++{
		competitors.items = append(competitors.items,Competitor{strconv.Itoa(i)})
	}
	result := calcGroups(competitors, 2)
	fmt.Println(result)
	TestingUtils.CheckEquals(2,len(result), "len of result should match groupcount", t)
	if result[0].Competitors.items[0].name != "1"{
		t.Error("first group first competitor should be Benni")
	}
}

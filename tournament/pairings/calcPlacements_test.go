package pairings

import (
	"reflect"
	"testing"
)

func Test_calcPlacements(t *testing.T) {
	var tests []struct {
		name string
		want []placement
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcPlacements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcPlacements() = %v, want %v", got, tt.want)
			}
		})
	}
}

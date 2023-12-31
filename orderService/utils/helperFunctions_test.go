package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/datatypes"
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	assert := require.New(t)

	newValues := []string{"a", "b", "c", "d"}
	presentValues := []string{"a", "b", "c", "d", "e", "f"}

	addedValues := DifferenceSliceStr(newValues, presentValues)

	assert.Equal([]string{}, addedValues)

	deletedValues := DifferenceSliceStr(presentValues, newValues)

	assert.Equal([]string{"e", "f"}, deletedValues)

	newValues = []string{"a", "b", "c", "d"}
	presentValues = []string{"a", "b", "c", "d"}

	addedValues = DifferenceSliceStr(newValues, presentValues)

	assert.Equal([]string{}, addedValues)

	deletedValues = DifferenceSliceStr(presentValues, newValues)

	assert.Equal([]string{}, deletedValues)

	newValues = []string{}
	presentValues = []string{"a", "b", "c", "d"}

	addedValues = DifferenceSliceStr(newValues, presentValues)

	assert.Equal([]string{}, addedValues)

	deletedValues = DifferenceSliceStr(presentValues, newValues)

	assert.Equal([]string{"a", "b", "c", "d"}, deletedValues)
}

func TestFilterNonEmptyStrings(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			name: "only strings",
			args: args{strSlice: []string{"apple", "", "banana", "", "cherry"}},
			want: []string{"apple", "banana", "cherry"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterNonEmptyStrings(tt.args.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterNonEmptyStrings() = %v, want %v", len(got), len(tt.want))
			}
		})
	}
}

func TestMaskSecret(t *testing.T) {
	type args struct {
		config        datatypes.JSON
		secretKeyPath []string
	}
	tests := []struct {
		name string
		args args
		want datatypes.JSON
	}{
		{
			name: "mask secret",
			args: args{
				config: datatypes.JSON(`{"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4"}`),
				secretKeyPath: []string{"key1", "key2"},
			},
			want: datatypes.JSON(fmt.Sprintf(`{"key1": "%s",
					"key2": "%s",
					"key3": "value3",
					"key4": "value4"}`, CredMask, CredMask)),
		},
		{
			name: "mask secret",
			args: args{
				config: datatypes.JSON(`{"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4"}`),
				secretKeyPath: []string{"key1"},
			},
			want: datatypes.JSON(fmt.Sprintf(`{"key1": "%s",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4"}`, CredMask)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskSecret(tt.args.config, tt.args.secretKeyPath...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaskSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

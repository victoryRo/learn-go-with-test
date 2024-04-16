package refle

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	exp := "Victor"
	var actual []string

	x := struct {
		Name string
	}{exp}

	Walk(x, func(input string) {
		actual = append(actual, input)
	})

	if len(actual) != 1 {
		t.Errorf("wrong number of function calls, actual %d want %d", len(actual), 1)
	}
	if actual[0] != exp {
		t.Errorf("Actual %q expected %q", actual[0], exp)
	}
}

func TestWalk2(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				Age  int
			}{"Victor", 35},
			[]string{"Victor"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Angela",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Angela", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string
			Walk2(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("actual %v, expected %v", actual, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk3(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
				Age  int
			}{"Victor", 35},
			[]string{"Victor"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Angela",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Angela", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string
			Walk3(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("actual %v, expected %v", actual, test.ExpectedCalls)
			}
		})
	}
}

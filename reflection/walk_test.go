package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("tests with guaranteed results orders", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "Atlanta"},
				[]string{"Chris", "Atlanta"},
			},
			{
				"Struct with a non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 40},
				[]string{"Chris"},
			},
			{
				"Nested fields",
				Person{
					"Chris",
					Profile{40, "Atlanta"},
				},
				[]string{"Chris", "Atlanta"},
			},
			{
				"Pointers to things",
				&Person{
					"Chris",
					Profile{40, "Atlanta"},
				},
				[]string{"Chris", "Atlanta"},
			},
			{
				"Slices",
				[]Profile{
					{40, "Centennial"},
					{41, "Breckenridge"},
				},
				[]string{"Centennial", "Breckenridge"},
			},
			{
				"Arrays",
				[2]Profile{
					{40, "Centennial"},
					{41, "Breckenridge"},
				},
				[]string{"Centennial", "Breckenridge"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("Got %v want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected '%v' to contain '%s' but it doesn't", haystack, needle)
	}
}

// Person struct
type Person struct {
	Name    string
	Profile Profile
}

// Profile struct
type Profile struct {
	Age  int
	City string
}

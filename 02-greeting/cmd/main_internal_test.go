package main

import "testing"

// func Examplemain() {
// 	main()
// 	// Output:
// 	// Hello, World!
// }

// func TestGreet_English(t *testing.T) {
// 	lang := language("en") // Preparation phase
// 	want := "Hello world"
// 	got := greet(lang) // Execution phase

// 	if got != want { // Decision phase
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

// func TestGreet_French(t *testing.T) {
// 	lang := language("fr") // Preparation phase
// 	want := "Bonjour le monde"
// 	got := greet(lang) // Execution phase

// 	if got != want { // Decision phase
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

// func TestGreet_Akkadian(t *testing.T) {
// 	// Akkadian is not implemented yet!
// 	lang := language("akk") // Preparation phase
// 	want := ""
// 	got := greet(lang) // Execution phase

// 	if got != want { // Decision phase
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}

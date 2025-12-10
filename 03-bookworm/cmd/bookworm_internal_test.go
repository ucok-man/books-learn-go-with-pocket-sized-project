package main

import (
	"testing"
)

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{
		Author: "Margaret Atwood", Title: "Oryx and Crake",
	}
	theBellJar = Book{
		Author: "Sylvia Plath", Title: "The Bell Jar",
	}
	janeEyre = Book{
		Author: "Charlotte Brontë", Title: "Jane Eyre",
	}
)

func TestLoadBookworm(t *testing.T) {
	type TestCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]TestCase{
		"file exist": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Fatal("expected an error to occur, got nil instead")
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Book{oryxAndCrake},
		},
		"two bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{handmaidsTale, theBellJar},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)

			if !equalBooks(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)

			if !equalBooksCount(t, tc.want, got) { // #3
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, found := got[book]
		if !found || targetCount != count {
			return false
		}
	}
	return true
}

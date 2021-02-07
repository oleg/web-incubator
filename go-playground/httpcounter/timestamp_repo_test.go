package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFileTimestampsRepo_LoadAll_empty_file(t *testing.T) {
	newFile := mustCreateNewTempFile()

	repo, err := NewFileTimestampsRepo(newFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	stamps, err := repo.LoadAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(stamps) != 0 {
		t.Errorf("Unexpected stamps count %v, expected %v", len(stamps), 0)
	}
}

func TestFileTimestampsRepo_LoadAll(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedStamps  []time.Time
		errorIsExpected bool
	}{
		{"empty",
			"",
			[]time.Time{},
			false},
		{"correct data",
			`2021-02-07T08:45:40.779056+01:00
2021-02-07T08:45:41.779327+01:00
2021-02-07T08:45:42.779341+01:00
`,
			[]time.Time{
				mustParseTime("2021-02-07T08:45:40.779056+01:00"),
				mustParseTime("2021-02-07T08:45:41.779327+01:00"),
				mustParseTime("2021-02-07T08:45:42.779341+01:00"),
			},
			false},
		{"wrong data",
			"hadfsdf",
			nil,
			true},
		{"wrong data in the middle",
			`2021-02-07T08:45:40.779056+01:00
asdfasdfasdf
2021-02-07T08:45:42.779341+01:00
`,
			nil,
			true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newFile := mustCreateNewTempFile()
			err := ioutil.WriteFile(newFile.Name(), []byte(test.input), 0666)
			if err != nil {
				log.Fatal(err)
			}
			repo, err := NewFileTimestampsRepo(newFile.Name())
			if err != nil {
				log.Fatal(err)
			}

			stamps, err := repo.LoadAll()

			errorIsPresent := err != nil
			if errorIsPresent != test.errorIsExpected {
				t.Errorf("Received error is present:%v, but expected:%v", err, test.errorIsExpected)
			}
			if !reflect.DeepEqual(stamps, test.expectedStamps) {
				t.Errorf("Read wrong stamps %v, expected %v", stamps, test.expectedStamps)
			}
		})
	}
}

func TestFileTimestampsRepo_AppendOne(t *testing.T) {
	newFile := mustCreateNewTempFile()

	repo, err := NewFileTimestampsRepo(newFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = repo.AppendOne(time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC))
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(newFile.Name())
	expected := "2020-01-02T03:04:05.000000006Z\n"
	if string(bytes) != expected {
		t.Errorf("Wrong file content %v, expcted %v", string(bytes), expected)
	}
}

func TestFileTimestampsRepo_StoreAll(t *testing.T) {
	newFile := mustCreateNewTempFile()
	err := ioutil.WriteFile(newFile.Name(), []byte("line-1\nline-2\nline-3\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := NewFileTimestampsRepo(newFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = repo.StoreAll([]time.Time{
		time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC),
		time.Date(2020, 2, 3, 4, 5, 6, 7, time.UTC),
		time.Date(2020, 3, 4, 5, 6, 7, 8, time.UTC),
		time.Date(2020, 4, 5, 6, 7, 8, 9, time.UTC),
		time.Date(2020, 5, 6, 7, 8, 9, 1, time.UTC),
	})
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(newFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	expected := `2020-01-02T03:04:05.000000006Z
2020-02-03T04:05:06.000000007Z
2020-03-04T05:06:07.000000008Z
2020-04-05T06:07:08.000000009Z
2020-05-06T07:08:09.000000001Z
`
	if string(content) != expected {
		t.Errorf("Wrong file content after save all %v expected %v", string(content), expected)
	}
}

func mustParseTime(timeStr string) time.Time {
	parse, err := time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		panic(err)
	}
	return parse
}

func mustCreateNewTempFile() *os.File {
	newFile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	return newFile
}

type testTimestampsRepo struct {
	stamps []time.Time
	err    error
}

func (r *testTimestampsRepo) LoadAll() ([]time.Time, error) {
	return r.stamps, r.err
}

func (r *testTimestampsRepo) AppendOne(time.Time) error {
	return r.err
}

func (r *testTimestampsRepo) StoreAll([]time.Time) error {
	return r.err
}

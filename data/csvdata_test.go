package data_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

func TestSimpleCaseCsv(t *testing.T) {
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name:   "test1",
				Format: "(This|is|test)",
			},
		},
	}

	d := data.NewCsvData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	rows := [][]string{}
	_ = json.Unmarshal(b, &rows)
	for _, row := range rows {
		for _, col := range row {
			if col != "This" && col != "is" && col != "test" {
				t.Fail()
			}
		}
	}
}

func TestSimpleCaseCsvHeader(t *testing.T) {
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name:   "test1",
				Format: "(This|is|test)",
			},
		},
	}

	headers := data.GetColumnHeader(tests.Columns, []string{})
	if headers[0] != "test1" {
		t.Fail()
	}
}

func TestObjectCaseCsv(t *testing.T) {
	child1Regex := `[0-9]{2,3}-[0-9]{5}`
	child2Regex := `[a-z0-9]{8}-[a-z0-9]{8}`
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "test1",
				Columns: []config.Column{
					{
						Name:   "child1",
						Format: child1Regex,
					},
					{
						Name:   "child2",
						Format: child2Regex,
					},
				},
			},
		},
	}

	d := data.NewCsvData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	rows := [][]string{}
	_ = json.Unmarshal(b, &rows)
	for _, row := range rows {
		for _, col := range row {
			testing1 := regexp.MustCompile(child1Regex)
			testing2 := regexp.MustCompile(child2Regex)
			if !testing1.Match([]byte(col)) && !testing2.Match([]byte(col)) {
				t.Fail()
			}
		}
	}
}

func TestObjectCaseCsvHeader(t *testing.T) {
	child1Regex := `[0-9]{2,3}-[0-9]{5}`
	child2Regex := `[a-z0-9]{8}-[a-z0-9]{8}`
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "test1",
				Columns: []config.Column{
					{
						Name:   "child1",
						Format: child1Regex,
					},
					{
						Name:   "child2",
						Format: child2Regex,
					},
				},
			},
		},
	}

	headers := data.GetColumnHeader(tests.Columns, []string{})
	if headers[0] != "child1" {
		t.Fail()
	}

	if headers[1] != "child2" {
		t.Fail()
	}
}

func TestListCaseCsv(t *testing.T) {
	child1Regex := `[0-9]{2,3}-[0-9]{5}`
	child2Regex := `[a-z0-9]{8}-[a-z0-9]{8}`
	child3Regex := `\$[0-9][A-Z]{2}_\$[0-9]{2}[A-Z]{3}`
	test2Regex := `(T|E|S)`
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "test1",
				Rows: &config.Rows{
					Size: 5,
					Columns: []config.Column{
						{
							Name:   "child1",
							Format: child1Regex,
						},
						{
							Name:   "child2",
							Format: child2Regex,
						},
						{
							Name: "child3",
							Rows: &config.Rows{
								Size: 10,
								Columns: []config.Column{
									{
										Name:   "child3",
										Format: child3Regex,
									},
									{
										Name: "child5",
										Rows: &config.Rows{
											Size: 20,
											Columns: []config.Column{
												{
													Name:   "child6",
													Format: child2Regex,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Name:   "test2",
				Format: test2Regex,
			},
		},
	}

	d := data.NewCsvData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	rows := [][]string{}
	_ = json.Unmarshal(b, &rows)
	for _, row := range rows {
		for _, col := range row {
			testing1 := regexp.MustCompile(child1Regex)
			testing2 := regexp.MustCompile(child2Regex)
			testing3 := regexp.MustCompile(child3Regex)
			testing4 := regexp.MustCompile(test2Regex)
			if !testing1.Match([]byte(col)) && !testing2.Match([]byte(col)) && !testing3.Match([]byte(col)) && !testing4.Match([]byte(col)) {
				t.Fail()
			}
		}
	}
}

func TestListCaseCsvHeader(t *testing.T) {
	child1Regex := `[0-9]{2,3}-[0-9]{5}`
	child2Regex := `[a-z0-9]{8}-[a-z0-9]{8}`
	child3Regex := `\$[0-9][A-Z]{2}_\$[0-9]{2}[A-Z]{3}`
	test2Regex := `(T|E|S)`
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "test1",
				Rows: &config.Rows{
					Size: 5,
					Columns: []config.Column{
						{
							Name:   "child1",
							Format: child1Regex,
						},
						{
							Name:   "child2",
							Format: child2Regex,
						},
						{
							Name: "child3",
							Rows: &config.Rows{
								Size: 10,
								Columns: []config.Column{
									{
										Name:   "child3",
										Format: child3Regex,
									},
									{
										Name: "child5",
										Rows: &config.Rows{
											Size: 20,
											Columns: []config.Column{
												{
													Name:   "child6",
													Format: child2Regex,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Name:   "test2",
				Format: test2Regex,
			},
		},
	}

	headers := data.GetColumnHeader(tests.Columns, []string{})
	if headers[0] != "child1" {
		t.Fail()
	}
	if headers[1] != "child2" {
		t.Fail()
	}
	if headers[2] != "child3" {
		t.Fail()
	}
	if headers[3] != "child6" {
		t.Fail()
	}
	if headers[4] != "test2" {
		t.Fail()
	}
}

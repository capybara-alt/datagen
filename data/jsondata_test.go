package data_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

func TestSimpleCase(t *testing.T) {
	testRegex1 := "(This|is|test)"
	testRegex2 := "[a-z]{1,2}"
	testRegex3 := "[0-9]"
	testRegex4 := "[A-Z]{5}"
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name:   "test1",
				Attrs:  map[string]string{"attr1": testRegex2, "attr2": testRegex3},
				Format: testRegex1,
			},
			{
				Name:   "test2",
				Format: testRegex4,
			},
		},
	}

	d := data.NewJsonData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	result := make(map[string]interface{})
	json.Unmarshal(b, &result)
	if !regexp.MustCompile(testRegex1).Match([]byte(result["test1"].(map[string]interface{})["#text"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex4).Match([]byte(result["test2"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex2).Match([]byte(result["test1"].(map[string]interface{})["@attr1"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex3).Match([]byte(result["test1"].(map[string]interface{})["@attr2"].(string))) {
		t.Fail()
	}
}

func TestObjectCase(t *testing.T) {
	testRegex1 := "[0-9]{2,3}-[0-9]{5}"
	testRegex2 := "[a-z0-9]{8}-[a-z0-9]{8}"
	testRegex3 := "[a-z0-9]{8}-[a-z0-9]{8}"
	testRegex4 := "[A-Z]"
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name:  "test1",
				Attrs: map[string]string{"attr": testRegex4},
				Columns: []config.Column{
					{
						Name:   "child1",
						Attrs:  map[string]string{"attr1": testRegex2},
						Format: testRegex1,
					},
					{
						Name:   "child2",
						Format: testRegex3,
					},
				},
			},
		},
	}

	d := data.NewJsonData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	result := make(map[string]interface{})
	json.Unmarshal(b, &result)
	if !regexp.MustCompile(testRegex4).Match([]byte(result["test1"].(map[string]interface{})["@attr"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex1).Match([]byte(result["test1"].(map[string]interface{})["child1"].(map[string]interface{})["#text"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex2).Match([]byte(result["test1"].(map[string]interface{})["child1"].(map[string]interface{})["@attr1"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile(testRegex3).Match([]byte(result["test1"].(map[string]interface{})["child2"].(string))) {
		t.Fail()
	}
}

func TestListCase1(t *testing.T) {
	testRegex1 := "(A|T|R|S)"
	testRegex2 := "z"
	testRegex3 := "ll"
	testRegex4 := "AA_llZZZZ"
	testRegex5 := "T"
	testRegex6 := "(L|I|S|T)"
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name:  "test1",
				Attrs: map[string]string{"attr1": testRegex1},
				Rows: &config.Rows{
					Size: 5,
					Columns: []config.Column{
						{
							Name:   "child1",
							Format: testRegex2,
						},
						{
							Name:   "child2",
							Format: testRegex3,
						},
						{
							Name:   "child3",
							Format: testRegex4,
						},
					},
				},
			},
			{
				Name:   "test2",
				Format: testRegex5,
			},
			{
				Name:  "test3",
				Attrs: map[string]string{"attr1": "A"},
				Columns: []config.Column{
					{
						Name: "val",
						Rows: &config.Rows{
							Size:   10,
							Format: testRegex6,
						},
					},
				},
			},
		},
	}

	d := data.NewJsonData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	result := make(map[string]interface{})
	json.Unmarshal(b, &result)
	for _, value := range result["test1"].([]interface{}) {
		if !regexp.MustCompile(testRegex1).Match([]byte(value.(map[string]interface{})["@attr1"].(string))) {
			t.Fail()
		}
		if !regexp.MustCompile(testRegex2).Match([]byte(value.(map[string]interface{})["child1"].(string))) {
			t.Fail()
		}
		if !regexp.MustCompile(testRegex3).Match([]byte(value.(map[string]interface{})["child2"].(string))) {
			t.Fail()
		}
		if !regexp.MustCompile(testRegex4).Match([]byte(value.(map[string]interface{})["child3"].(string))) {
			t.Fail()
		}
	}
	if !regexp.MustCompile(testRegex5).Match([]byte(result["test2"].(string))) {
		t.Fail()
	}
	if !regexp.MustCompile("A").Match([]byte(result["test3"].(map[string]interface{})["@attr1"].(string))) {
		t.Fail()
	}
	for _, value := range result["test3"].(map[string]interface{})["val"].([]interface{}) {
		if !regexp.MustCompile(testRegex6).Match([]byte(value.(string))) {
			t.Fail()
		}
	}
}

func TestListCase2(t *testing.T) {
	testRegex1 := "(Hello|World|!)"
	tests := &config.Config{
		Rows: &config.Rows{
			Size: 100,
			Columns: []config.Column{
				{
					Name:   "name",
					Format: testRegex1,
				},
			},
		},
	}

	d := data.NewJsonData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	result := []interface{}{}
	json.Unmarshal(b, &result)
	if len(result) < 1 {
		t.Fail()
	}
	for _, value := range result {
		if !regexp.MustCompile(testRegex1).Match([]byte(value.(map[string]interface{})["name"].(string))) {
			t.Fail()
		}
	}
}

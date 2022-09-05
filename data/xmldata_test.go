package data_test

import (
	"regexp"
	"testing"

	"github.com/capybara-alt/datagen/config"
	"github.com/capybara-alt/datagen/data"
)

func TestSimpleXmlCase(t *testing.T) {
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "root",
				Columns: []config.Column{
					{
						Name:   "test1",
						Attrs:  map[string]string{"attr1": "[1-9]", "attr2": "(true|false)"},
						Format: "(This|is|test)",
					},
				},
			},
		},
	}

	d := data.NewXmlData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	matched1, _ := regexp.Match("<root><test1 attr1=\"[1-9]\" attr2=\"(true|false)\">(This|is|test)</test1></root>", b)
	matched2, _ := regexp.Match("<root><test1 attr2=\"(true|false)\" attr1=\"[1-9]\">(This|is|test)</test1></root>", b)
	if !matched1 && !matched2 {
		t.Fail()
	}
}

func TestObjectXmlCase(t *testing.T) {
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "test1",
				Columns: []config.Column{
					{
						Name:   "child1",
						Format: "[0-9]{2,3}-[0-9]{5}",
					},
					{
						Name:   "child2",
						Format: "[a-z0-9]{8}-[a-z0-9]{8}",
					},
				},
			},
		},
	}

	d := data.NewXmlData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	matched, _ := regexp.Match("<test1><child1>[0-9]{2,3}-[0-9]{5}</child1><child2>[a-z0-9]{8}-[a-z0-9]{8}</child2></test1>", b)
	if !matched {
		t.Fail()
	}
}

func TestListXmlCase(t *testing.T) {
	tests := &config.Config{
		Columns: []config.Column{
			{
				Name: "root",
				Columns: []config.Column{
					{
						Name: "test1",
						Rows: &config.Rows{
							Size: 5,
							Columns: []config.Column{
								{
									Name:   "child1",
									Format: "[a-z]",
								},
								{
									Name:   "child2",
									Format: "[0-9]{1,2}",
								},
								{
									Name: "child3-0",
									Rows: &config.Rows{
										Size:   2,
										Format: "[0-9][A-Z]{2}_[0-9]{2}[A-Z]{5}",
									},
								},
							},
						},
					},
					{
						Name:   "test2",
						Format: "(T|E|S)",
					},
				},
			},
		},
	}

	d := data.NewXmlData(tests)
	b, _ := d.GetValue()
	t.Log(string(b))
	xmlStr := "<root>"
	for i := 0; i < 5; i++ {
		xmlStr += "<test1>"
		xmlStr += "<child1>[a-z]</child1>"
		xmlStr += "<child2>[0-9]{1,2}</child2>"
		for j := 0; j < 2; j++ {
			xmlStr += "<child3-0>[0-9][A-Z]{2}_[0-9]{2}[A-Z]{5}</child3-0>"
		}
		xmlStr += "</test1>"
	}
	xmlStr += "<test2>(T|E|S)</test2>"
	xmlStr += "</root>"
	if !regexp.MustCompile(xmlStr).Match(b) {
		t.Fail()
	}
}

package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFixture(t *testing.T) {
	raw, e := os.ReadFile("../../data/contacts.json")
	if e != nil {
		t.Fatal(e)
	}
	var d Dataset
	if e = json.Unmarshal(raw, &d); e != nil {
		t.Fatal(e)
	}
	if len(d.Contacts) != 5 {
		t.Fatalf("contacts=%d", len(d.Contacts))
	}
	for _, c := range d.Contacts {
		if c.Status != "VERIFIED" || c.SourceURL == "" || c.CheckedAt == "" {
			t.Fatalf("unsafe record: %s", c.ID)
		}
	}
}

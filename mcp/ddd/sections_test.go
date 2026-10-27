package main

import "testing"

func TestLoadSectionsKeepsBookOrderAndIDs(t *testing.T) {
	sections, err := loadSections()
	if err != nil {
		t.Fatal(err)
	}
	if sections[0].ID != "preface" {
		t.Fatalf("first section = %q, want preface", sections[0].ID)
	}
	if _, ok := find(sections, "aggregates"); !ok {
		t.Fatal("aggregates section missing")
	}
}

func TestFindUnknownSection(t *testing.T) {
	sections, _ := loadSections()
	if _, ok := find(sections, "nope"); ok {
		t.Fatal("unexpected section")
	}
}

func TestSearchIsCaseInsensitiveAndLimited(t *testing.T) {
	sections, _ := loadSections()
	hits := search(sections, "AGRÉGAT", 2)
	if len(hits) == 0 || len(hits) > 2 {
		t.Fatalf("got %d hits, want 1 or 2", len(hits))
	}
}

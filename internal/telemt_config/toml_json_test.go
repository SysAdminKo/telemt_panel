package telemt_config

import (
	"strings"
	"testing"
)

func TestSectionsToTOMLAndBack(t *testing.T) {
	sections := map[string]interface{}{
		"general":    map[string]interface{}{"use_middle_proxy": true, "ad_tag": "abc"},
		"censorship": map[string]interface{}{"tls_domain": "www.google.com"},
	}
	text, err := SectionsToTOML(sections)
	if err != nil {
		t.Fatalf("SectionsToTOML: %v", err)
	}
	if !strings.Contains(text, "[general]") || !strings.Contains(text, "tls_domain") {
		t.Fatalf("unexpected TOML:\n%s", text)
	}

	back, err := TOMLToSections(text)
	if err != nil {
		t.Fatalf("TOMLToSections: %v", err)
	}
	g, ok := back["general"].(map[string]interface{})
	if !ok || g["ad_tag"] != "abc" {
		t.Fatalf("round-trip lost data: %#v", back)
	}
	if g["use_middle_proxy"] != true {
		t.Fatalf("round-trip lost bool type: %#v", g["use_middle_proxy"])
	}
}

func TestTOMLToSectionsRejectsInvalid(t *testing.T) {
	if _, err := TOMLToSections("this is = = not toml"); err == nil {
		t.Fatal("expected parse error")
	}
}

func TestTOMLToSectionsEmptyReturnsNonNilMap(t *testing.T) {
	m, err := TOMLToSections("")
	if err != nil {
		t.Fatalf("empty doc: %v", err)
	}
	if m == nil {
		t.Fatal("expected non-nil empty map for empty document")
	}
}

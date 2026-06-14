package telemt_config

import (
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

// SectionsToTOML renders a map of config sections (as returned by Telemt's
// GET /v1/config) into TOML text for the editor. Integer underscores are
// stripped to match the format Telemt parses (e.g. 8_443 -> 8443).
func SectionsToTOML(sections map[string]interface{}) (string, error) {
	out, err := toml.Marshal(sections)
	if err != nil {
		return "", fmt.Errorf("marshal sections: %w", err)
	}
	return removeIntegerUnderscores(string(out)), nil
}

// TOMLToSections parses editor TOML text back into a generic section map
// suitable for a PATCH /v1/config body.
func TOMLToSections(content string) (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := toml.Unmarshal([]byte(content), &m); err != nil {
		return nil, fmt.Errorf("parse sections: %w", err)
	}
	if m == nil {
		m = map[string]interface{}{}
	}
	return m, nil
}

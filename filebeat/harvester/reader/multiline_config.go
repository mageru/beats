package reader

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/common/match"
)

type MultilineConfig struct {
	Negate   bool           `config:"negate"`
	Match    string         `config:"match"       validate:"required"`
	MaxLines *int           `config:"max_lines"`
	Pattern  match.Matcher  `config:"pattern"`
	Timeout  *time.Duration `config:"timeout"     validate:"positive"`
}

func (c *MultilineConfig) Validate() error {
	if c.Match != "after" && c.Match != "before" {
		return fmt.Errorf("unknown matcher type: %s", c.Match)
	}
	return nil
}

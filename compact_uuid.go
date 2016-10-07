package cuuid

import (
	"encoding/base64"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")

// CUUID is a compact form of UUID
type CUUID struct {
	UUID uuid.UUID
}

// NewV4 returns a compact version of UUID
func NewV4() *CUUID {
	return &CUUID{uuid.NewV4()}
}

// String returns a string
func (c *CUUID) String() string {
	encoded := base64.RawURLEncoding.EncodeToString(c.UUID.Bytes())
	return escaper.Replace(encoded)
}

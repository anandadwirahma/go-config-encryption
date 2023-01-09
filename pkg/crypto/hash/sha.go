package hash

import (
	"crypto/sha256"
	"fmt"
)

func HashingSHA256(message string) string {
	sha := sha256.Sum256([]byte(message))
	return fmt.Sprintf("%x", sha)
}

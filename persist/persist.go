package persist

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	"github.com/NebulousLabs/Sia/build"
)

var (
	ErrBadVersion = errors.New("incompatible version")
	ErrBadHeader  = errors.New("wrong header")

	HomeFolder = func() string {
		home, err := homedir.Dir()
		if err != nil {
			println("could not find homedir: " + err.Error()) // TODO: better error handling
			return ""
		}

		switch build.Release {
		case "testing":
			return filepath.Join(build.SiaTestingDir, "home")
		case "dev":
			return filepath.Join(home, ".config", "SiaDev")
		default:
			return filepath.Join(home, ".config", "Sia")
		}
	}()
)

// Metadata contains the header and version of the data being stored.
type Metadata struct {
	Header, Version string
}

// RandomSuffix returns a 20 character base32 suffix for a filename. There are
// 100 bits of entropy, and a very low probability of colliding with existing
// files unintentionally.
func RandomSuffix() string {
	randBytes := make([]byte, 20)
	rand.Read(randBytes)
	str := base32.StdEncoding.EncodeToString(randBytes)
	return str[:20]
}

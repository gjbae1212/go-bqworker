package internal

import (
	"cloud.google.com/go/compute/metadata"
	"github.com/gjbae1212/go-ec2meta"
	"os"
	"sync"
)

var (
	hostname     string
	onceHostname = sync.Once{}
)

// GetHostname returns hostname.
func GetHostname() string {
	onceHostname.Do(func() {
		// check aws hostname.
		if name, err := ec2meta.Hostname(); err == nil && len(name) > 0 {
			hostname = name[0]
			return
		}

		// check gcp hostname
		if name, err := metadata.Hostname(); err == nil {
			hostname = name
			return
		}

		// os hostname
		if name, err := os.Hostname(); err == nil {
			hostname = name
			return
		}
		hostname = "localhost"
	})
	return hostname
}

func init() {
	GetHostname()
}

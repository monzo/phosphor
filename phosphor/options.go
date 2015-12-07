package phosphor

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"log"
	"os"
)

type Options struct {
	// basic options
	ID           int64  `flag:"worker-id" cfg:"id"`
	Verbose      bool   `flag:"verbose"`
	HTTPAddress  string `flag:"http-address"`
	HTTPSAddress string `flag:"https-address"`
}

func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	h := md5.New()
	io.WriteString(h, hostname)
	defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)

	return &Options{
		ID: defaultID,
	}
}

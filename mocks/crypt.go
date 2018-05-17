package mocks

import (
	"crypto"
	"crypto/rand"

	"github.com/pkg/errors"
)

type message struct {
	msg []byte
}

func (m *message) decrypt(d crypto.Decrypter) (string, error) {
	r, err := d.Decrypt(rand.Reader, m.msg, nil)
	return string(r), errors.Wrap(err, "can't decrypt")
}

package appregistry

import (
	"github.com/operator-framework/operator-marketplace/pkg/datastore"
	yaml "gopkg.in/yaml.v2"
)

type blobUnmarshaler interface {
	// Unmarshal unmarshals package blob into structured representations
	Unmarshal(in []byte) (*datastore.Manifest, error)
}

type blobUnmarshalerImpl struct{}

func (*blobUnmarshalerImpl) Unmarshal(in []byte) (*datastore.Manifest, error) {
	m := &datastore.Manifest{}
	if err := yaml.Unmarshal(in, m); err != nil {
		return nil, err
	}

	return m, nil
}

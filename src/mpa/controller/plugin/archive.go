package plugin

import (
	"archive/zip"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"strings"
)

var (
	ErrTooLargeMikutterYml = errors.New("mpa/controller/plugin: .mikutter.yml is too large")
	ErrMikutterYmlNotFound = errors.New("mpa/controller/plugin: .mikutter.yml is not found")
)

type Spec struct {
	Slug        string
	Name        string
	Description string
	Version     string
	Dependency  Dependency `yaml:"depends"`
}

type Dependency struct {
	MikutterVersion string   `yaml:"mikutter"`
	Plugins         []string `yaml:"plugin"`
}

const MikutterYmlSizeLimit = 10 * 1024

func LoadSpec(path string) (Spec, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return Spec{}, err
	}
	defer r.Close()

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, ".mikutter.yml") {
			if f.UncompressedSize64 > MikutterYmlSizeLimit {
				return Spec{}, ErrTooLargeMikutterYml
			}
			rc, err := f.Open()
			if err != nil {
				return Spec{}, err
			}
			buf := make([]byte, f.UncompressedSize64)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				return Spec{}, err
			}
			rc.Close()

			s := Spec{}
			err = yaml.Unmarshal(buf, &s)
			return s, err
		}
	}
	return Spec{}, ErrMikutterYmlNotFound
}

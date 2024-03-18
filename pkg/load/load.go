package load

import (
	"github.com/koding/multiconfig"
	"io"
	"strings"
	"sync"
)

var (
	once sync.Once
)

func FileConfig(conf interface{}, fPaths ...string) {
	once.Do(func() {

		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fPath := range fPaths {
			if strings.HasSuffix(fPath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fPath})
			}
			if strings.HasSuffix(fPath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fPath})
			}
			if strings.HasSuffix(fPath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fPath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}

		m.MustLoad(conf)
	})
}

func ReaderConfig(conf interface{}, readerName string, r io.Reader) {
	loaders := []multiconfig.Loader{
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{},
	}

	if strings.HasSuffix(readerName, "toml") {
		loaders = append(loaders, &multiconfig.TOMLLoader{Reader: r})
	}
	if strings.HasSuffix(readerName, "json") {
		loaders = append(loaders, &multiconfig.TOMLLoader{Reader: r})
	}
	if strings.HasSuffix(readerName, "yaml") {
		loaders = append(loaders, &multiconfig.TOMLLoader{Reader: r})
	}

	m := multiconfig.DefaultLoader{
		Loader:    multiconfig.MultiLoader(loaders...),
		Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
	}

	m.MustLoad(conf)
}

package plugin

import (
	"io"
	"mpa/model"
	"mpa/route"
	"os"
	"path/filepath"
)

type DownloadView struct {
	Plugin model.Plugin
}

const StoragePath = "/app/storage"

func (v *DownloadView) Render(ctx *route.Context) error {
	ctx.ResponseWriter.Header().Add("Content-Type", "application/zip")
	path := filepath.Join(StoragePath, v.Plugin.Uuid.String())
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(ctx.ResponseWriter, f)
	return nil
}

package up

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/iyear/tdl/pkg/uploader"
	"os"
	"path/filepath"
)

type iter struct {
	files []string
	cur   int
}

func newIter(files []string) *iter {
	return &iter{
		files: files,
		cur:   -1,
	}
}

func (i *iter) Next(_ context.Context) bool {
	i.cur++

	if i.cur == len(i.files) {
		return false
	}

	return true
}

func (i *iter) Value(_ context.Context) (*uploader.Item, error) {
	cur := i.files[i.cur]

	mime, err := mimetype.DetectFile(cur)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(cur)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &uploader.Item{
		R:    f,
		Name: filepath.Base(f.Name()),
		MIME: mime.String(),
		Size: stat.Size(),
	}, nil
}

func (i *iter) Total(_ context.Context) int {
	return len(i.files)
}
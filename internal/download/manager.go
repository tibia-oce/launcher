// internal\download\manager.go

package download

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
)

type DownloadManager struct {
	totalBytes      int64
	downloadedBytes int64
}

// TODO: Needs app directory path, can't just use dst..
func (d *DownloadManager) DownloadFile(url, dst string, progress bool) error {
	dst = filepath.Join(dst)
	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	var reader io.Reader = resp.Body
	if progress {
		reader = io.TeeReader(reader, d)
	}

	_, err = io.Copy(out, reader)
	if err != nil {
		return err
	}

	atomic.AddInt64(&d.downloadedBytes, 1)

	return nil
}

func (d *DownloadManager) Write(p []byte) (n int, err error) {
	n = len(p)
	atomic.AddInt64(&d.downloadedBytes, int64(n))
	return
}

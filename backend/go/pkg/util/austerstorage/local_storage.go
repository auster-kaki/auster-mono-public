package austerstorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ContentType string

const (
	JPEG ContentType = "image/jpeg"
	PNG  ContentType = "image/png"
	HEIC ContentType = "image/heic"
	HEIF ContentType = "image/heif"
)

// Save はローカルにデータを保存しパスを返す
func Save(content ContentType, path string, data []byte) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path = filepath.Join(pwd, "assets", "images", path)
	switch content {
	case PNG, JPEG:
		// ok
	case HEIC, HEIF:
	default:
		return "", fmt.Errorf("invalid content type: %s", content)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", err
	}

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return path, nil
}

func Get(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

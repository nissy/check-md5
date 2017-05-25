package md5ck

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

type File struct {
	PATH string
	HASH string
}

func (f *File) equals() error {
	h, err := fileHash(f.PATH)

	if err != nil {
		return err
	}

	if h != f.HASH {
		return PluginError{
			Message:  f.PATH + ": does not file match",
			ExitCode: CRITICAL,
		}
	}

	return nil
}

func (f *File) set() (err error) {
	if f.HASH, err = fileHash(f.PATH); err != nil {
		return err
	}

	return nil
}

func fileHash(name string) (string, error) {
	f, err := os.Open(name)

	if err != nil {
		return "", err
	}

	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)[:16]), nil
}

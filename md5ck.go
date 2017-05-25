package md5ck

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

var exitCodeText = map[int]string{
	OK:       "OK",
	WARNING:  "WARNING",
	CRITICAL: "CRITICAL",
	UNKNOWN:  "UNKNOWN",
}

type (
	MD5ck struct {
		Files []*File
	}

	File struct {
		PATH string
		SUM  string
	}

	PluginError struct {
		Message  string
		ExitCode int
	}
)

func New() *MD5ck {
	return &MD5ck{}
}

func (e PluginError) Error() string {
	return fmt.Sprint(e.Message)
}

func ExitCodeText(code int) string {
	return exitCodeText[code]
}

func (ck *MD5ck) Do() error {
	var i int

	for _, v := range ck.Files {
		i++
		if err := v.do(); err != nil {
			return err
		}
	}

	if i == 0 {
		return PluginError{
			Message:  "No check has been done",
			ExitCode: WARNING,
		}
	}

	return nil
}

func (fck *File) do() error {
	f, err := os.Open(fck.PATH)

	if err != nil {
		return err
	}

	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	if hex.EncodeToString(h.Sum(nil)[:16]) != fck.SUM {
		return PluginError{
			Message:  fck.PATH + ": does not file match",
			ExitCode: CRITICAL,
		}
	}

	return nil
}

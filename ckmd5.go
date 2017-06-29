package ckmd5

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
	CkMD5 struct {
		Files []*File
	}

	PluginError struct {
		Message  string
		ExitCode int
	}
)

func New() *CkMD5 {
	return &CkMD5{}
}

func (e PluginError) Error() string {
	return e.Message
}

func ExitCodeText(code int) string {
	return exitCodeText[code]
}

func (ck *CkMD5) Equals() error {
	var i int

	for _, v := range ck.Files {
		i++
		if err := v.equals(); err != nil {
			return err
		}
	}

	if i == 0 {
		return PluginError{
			Message:  "No check has been done",
			ExitCode: UNKNOWN,
		}
	}

	return nil
}

func (ck *CkMD5) Set() error {
	for _, v := range ck.Files {
		if err := v.set(); err != nil {
			return err
		}
	}

	return nil
}

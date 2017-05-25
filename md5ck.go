package md5ck

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

	PluginError struct {
		Message  string
		ExitCode int
	}
)

func New() *MD5ck {
	return &MD5ck{}
}

func (e PluginError) Error() string {
	return e.Message
}

func ExitCodeText(code int) string {
	return exitCodeText[code]
}

func (ck *MD5ck) Equals() error {
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
			ExitCode: WARNING,
		}
	}

	return nil
}

func (ck *MD5ck) Set() error {
	var i int

	for _, v := range ck.Files {
		i++
		if err := v.set(); err != nil {
			return err
		}
	}

	return nil
}

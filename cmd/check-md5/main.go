package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nissy/check-md5"
	"gopkg.in/BurntSushi/toml.v0"
)

const (
	defaultCfgName = "check-md5.conf"
	version        = "0.1"
)

var (
	cfgName   = flag.String("c", defaultCfgName, "set cfgiguration file")
	isSave    = flag.Bool("s", false, "save cfgiguration file md5 hash")
	isHelp    = flag.Bool("h", false, "this help")
	isVersion = flag.Bool("v", false, "show this build version")
)

func main() {
	os.Exit(exitcode(run()))
}

func exitcode(err error) int {
	if err != nil {
		if *isSave {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return 1
		}

		switch err := err.(type) {
		case ckmd5.PluginError:
			fmt.Fprintf(os.Stderr, "%s: %s\n", ckmd5.ExitCodeText(err.ExitCode), err.Message)
			return err.ExitCode
		default:
			fmt.Fprintf(os.Stderr, "%s: %s\n", ckmd5.ExitCodeText(ckmd5.CRITICAL), err.Error())
			return ckmd5.CRITICAL
		}
	}

	return ckmd5.OK
}

func run() error {
	flag.Parse()

	if *isVersion {
		fmt.Println("v" + version)
		return nil
	}

	if *isHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		return nil
	}

	ck := ckmd5.New()

	if _, err := toml.DecodeFile(*cfgName, &ck); err != nil {
		return err
	}

	if *isSave {
		if err := ck.Set(); err != nil {
			return err
		}

		f, err := os.OpenFile(*cfgName, os.O_WRONLY, 0644)

		if err != nil {
			return err
		}

		enc := toml.NewEncoder(f)

		if err := enc.Encode(ck); err != nil {
			return err
		}

		fmt.Println("Save config file " + *cfgName)

		return nil
	}

	if err := ck.Equals(); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s: %s\n", ckmd5.ExitCodeText(ckmd5.OK), "All matched")
	return nil
}

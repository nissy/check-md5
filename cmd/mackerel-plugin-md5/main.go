package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nissy/mackerel-plugin-md5"
	"gopkg.in/BurntSushi/toml.v0"
)

const (
	defaultCfgName = "md5ck.conf"
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
		case md5ck.PluginError:
			fmt.Fprintf(os.Stderr, "%s: %s\n", md5ck.ExitCodeText(err.ExitCode), err.Message)
			return err.ExitCode
		default:
			fmt.Fprintf(os.Stderr, "%s: %s\n", md5ck.ExitCodeText(md5ck.CRITICAL), err.Error())
			return md5ck.CRITICAL
		}
	}

	return md5ck.OK
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

	ck := md5ck.New()

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

	fmt.Fprintf(os.Stdout, "%s: %s\n", md5ck.ExitCodeText(md5ck.OK), "All matched")
	return nil
}

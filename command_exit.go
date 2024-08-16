package main

import "os"

func commandExit(_ *config, _ ...string) error {
	defer os.Exit(0)
	return nil
}

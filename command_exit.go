package main

import "os"

func commandExit(_ *config) error {
	defer os.Exit(0)
	return nil
}

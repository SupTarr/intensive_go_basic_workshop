package main

import (
	"log/slog"
	"strconv"
)

func addStringNumber(a, b string) int {
	ca, err := strconv.Atoi(a)
	if err != nil {
		slog.Info(err.Error())
	}

	cb, err := strconv.Atoi(b)
	if err != nil {
		slog.Info(err.Error())
	}

	return ca + cb
}

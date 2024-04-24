package addstring

import (
	"log/slog"
	"strconv"
	"strings"
)

func AddString(str string) int {
	if len(str) == 0 {
		return 0
	}

	sum := 0
	strs := strings.Split(str, "")
	for _, v := range strs {
		i, err := strconv.Atoi(v)
		if err != nil {
			slog.Info(err.Error())
		}

		sum += i
	}

	return sum
}

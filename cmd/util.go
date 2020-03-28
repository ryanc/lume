package lumecmd

import (
	"fmt"
	"time"

	"git.kill0.net/chill9/lume"
)

func powerColor(s string) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "on" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
}

func statusColor(s lifx.Status) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "ok" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
}

func PrintResults(res []lifx.Result) {
	var length, idWidth, labelWidth, statusWidth int

	for _, r := range res {
		length = len(r.Id)
		if idWidth < length {
			idWidth = length
		}

		length = len(r.Label)
		if labelWidth < length {
			labelWidth = length
		}

		length = len(r.Status)
		if statusWidth < length {
			statusWidth = length
		}
	}

	for _, r := range res {
		fmt.Printf("%*s %*s %*s\n",
			idWidth, r.Id,
			labelWidth, r.Label,
			statusWidth, statusColor(r.Status))
	}
}

func PrintLights(lights []lifx.Light) {
	var length int

	for _, l := range lights {
		length = len(l.Id)
		if idWidth < length {
			idWidth = length
		}

		length = len(l.Location.Name)
		if locationWidth < length {
			locationWidth = length
		}

		length = len(l.Group.Name)
		if groupWidth < length {
			groupWidth = length
		}

		length = len(l.Label)
		if labelWidth < length {
			labelWidth = length
		}

		length = len(l.LastSeen.Local().Format(time.RFC3339))
		if lastSeenWidth < length {
			lastSeenWidth = length
		}

		length = len(l.Power)
		if powerWidth < length {
			powerWidth = length
		}
	}

	fmt.Printf("total %d\n", len(lights))
	for _, l := range lights {
		fmt.Printf(
			"%*s %*s %*s %*s %*s %-*s\n",
			idWidth, l.Id,
			locationWidth, l.Location.Name,
			groupWidth, l.Group.Name,
			labelWidth, l.Label,
			lastSeenWidth, l.LastSeen.Local().Format(time.RFC3339),
			powerWidth, powerColor(l.Power),
		)
	}
}

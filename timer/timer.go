package timer

import (
	"time"

	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

func Start() {
	ticker := time.NewTicker(time.Duration(3800) * time.Second)
	go func() {
		for _ = range ticker.C {
			command.ReadSauna()
		}
	}()
}

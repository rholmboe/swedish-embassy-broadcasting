package timer

import (
	"time"

	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

func Start() {
	go command.ReadSauna()
	ticker := time.NewTicker(time.Duration(1200) * time.Second)
	go func() {
		for _ = range ticker.C {
			command.ReadSauna()
		}
	}()
}

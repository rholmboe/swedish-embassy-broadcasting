package fetch

import (
	"encoding/hex"
	"os/exec"
  "os"
  "path"
  "fmt"
)

type AudioOutput struct {
	Path string
	Info string
  CmdOutput string
}

func DownloadAudio(id string) (AudioOutput, error) {
	prefix := path.Join(os.Getenv("DUMP_PATH"), hex.EncodeToString([]byte(id)))
	cmd := exec.Command(
		"youtube-dl",
		"--extract-audio",
		"--audio-format=mp3",
		"--write-info-json",
		"--max-filesize=30m",
		"--output="+ prefix + ".%(ext)s",
		id)
	o, err := cmd.Output()
	output := AudioOutput{
    Path: prefix + ".mp3",
    Info: prefix + ".info.json",
    CmdOutput: o}
	return output, err
}

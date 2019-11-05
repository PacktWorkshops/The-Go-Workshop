package command_injection
import (
	"bytes"
	"io/ioutil"
	"os/exec"
)
func listFiles2(path string) (resp []string, err error) {
	const fileDir = "test"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return resp, err
	}
	for _, f := range files {
		resp = append(resp, f.Name())
	}
	return resp, nil
}
func listFiles(path string) (string, error) {
	cmd := exec.Command("bash", "-c", "ls"+path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

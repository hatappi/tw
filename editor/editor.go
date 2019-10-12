package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func EditText() ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", "tw")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = os.Remove(tmpfile.Name())
	}()
	defer func() {
		_ = tmpfile.Close()
	}()

	if err := openEditor(tmpfile.Name()); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(tmpfile.Name())
}

func openEditor(filepath string) error {
	cmdName := "vi"
	if e := os.Getenv("EDITOR"); e != "" {
		cmdName = e
	}

	c := exec.Command(cmdName, filepath)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

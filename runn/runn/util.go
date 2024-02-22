package runn

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// RunCommand is a function to run a command and return the output
func RunCommand(command string) (string, error) {
	cmd := exec.Command("/usr/bin/bash", "-c", command)
	fmt.Printf("exec: %s\n", cmd)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return stdout.String(), nil
}

func RunCommandVerbose(command string, args ...string) (string, error) {
	cmd := exec.Command("/usr/bin/bash", "-c", command)
	fmt.Printf("exec: %s\n", cmd)

	var stdoutBuf bytes.Buffer

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(stdoutPipe)

	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			stdoutBuf.Write(scanner.Bytes())
			stdoutBuf.WriteString("\n")
		}
	}()

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}

package main

import (
	"bytes"
	"errors"
	"os/exec"
	"syscall"
)

type response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func restartServer(server_name string) (*response, error) {
	args := []string{server_name}
	out, err := runCommand("echon", args)

	if err != nil {
		return nil, errors.New(err.Message)
	}

	return out, nil
}

func runCommand(command string, args []string) (*response, *response) {
	cmd := exec.Command(command, args...)
	cmdOutput := &bytes.Buffer{}

	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdOutput
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			return nil, &response{StatusCode: waitStatus.ExitStatus(), Message: cmdOutput.String()}
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		return &response{StatusCode: waitStatus.ExitStatus(), Message: cmdOutput.String()}, nil
	}
	return nil, nil
}

package main

import (
	"bytes"
	"os/exec"
	"syscall"
)

type response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func restartServer(server_name string) (*response, error) {
	args := []string{"restart", server_name}
	resp, err := runCommand("supervisorctl", args)
	return resp, err
}

func runCommand(command string, args []string) (*response, error) {
	cmd := exec.Command(command, args...)
	sbuf := &bytes.Buffer{}

	cmd.Stdout = sbuf
	cmd.Stderr = sbuf

	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			return &response{StatusCode: waitStatus.ExitStatus(), Message: sbuf.String()}, nil
		} else {
			return nil, err
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		return &response{StatusCode: waitStatus.ExitStatus(), Message: sbuf.String()}, nil
	}
}

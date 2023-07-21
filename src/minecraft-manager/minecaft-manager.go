package minecraft_manager

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"time"
)

type MinecraftManager struct {
	running   bool
	serverJar string
	process   *exec.Cmd
	stdout    string
	stdin     *os.File
}

func NewMinecraftManager(serverJar string) *MinecraftManager {
	minecraftServer := MinecraftManager{running: false, serverJar: serverJar}
	return &minecraftServer
}

func (server *MinecraftManager) IsServerRunning() bool {
	return server.running
}

func (server *MinecraftManager) Start() {
	cmd := exec.Command("java", "-Xmx2048M", "-Xms1024M", "-jar", server.serverJar, "--nogui")
	outputBuffer := &bytes.Buffer{}

	stdinRead, stdinWrite, err := os.Pipe()
	server.stdin = stdinWrite
	cmd.Stdin = stdinRead
	cmd.Stdout = outputBuffer
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go func() {
		cmd.Wait()
		server.running = false
	}()

	go func() {
		for true {
			readBytes, _ := io.ReadAll(outputBuffer)
			server.stdout = server.stdout + string(readBytes)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	server.process = cmd
	server.running = true
}

func (server *MinecraftManager) GetOutput() string {
	return server.stdout
}

func (server *MinecraftManager) SendCommand(command string) {
	io.WriteString(server.stdin, command+"\n")
}

func (server *MinecraftManager) Stop() {
	server.SendCommand("stop")
}

package minecraft_manager

import (
	"io"
	"os/exec"
)

type MinecraftManager struct {
	running   bool
	serverJar string
	process   *exec.Cmd
	Stdout    *io.ReadCloser
	Stderr    *io.ReadCloser
}

func NewMinecraftManager(serverJar string) MinecraftManager {
	return MinecraftManager{running: false, serverJar: serverJar}
}

func (server MinecraftManager) IsServerRunning() bool {
	return server.running
}

func (server MinecraftManager) Start() {
	cmd := exec.Command("java", "-jar", server.serverJar)

	stdout, error := cmd.StdoutPipe()
	server.Stdout = &stdout
	stderr, error2 := cmd.StderrPipe()

	println(error, error2)
	server.Stderr = &stderr
	server.process = cmd
	error3 := cmd.Start()
	println(error3)
}

func (server MinecraftManager) GetOutput() string {
	bytes, _ := io.ReadAll(*server.Stdout)
	return string(bytes)
}

func (server MinecraftManager) GetError() string {
	bytes, _ := io.ReadAll(*server.Stderr)
	return string(bytes)
}

func (server MinecraftManager) Stop() {

}

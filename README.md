# cube-manager
A simple minecraft server managment gui

### Important: Code is in heavy development and probably not secure. Currently is usage under own risk!

![login](assets/img/examples/login.png?raw=true "login")
![file list](assets/img/examples/filelist.png?raw=true "file list")

# Setup

1. Move the `cube-manager` executable into your minecraft server folder.
2. Create a `.env` file with the following content:
```
  CUBE_USER=<Your Username>
  CUBE_PASSWORD=<Your Password>
```
3. Start the `cube-manager` executable it as deamon/in the background. For example with systemd, screen etc.
4. Log into the Manager with your browser via `http://<YourServerIpOrHostname>:8080`.

# Building
You "need" Go version `go 1.18.*` to build the executable. 

1. Get the sourcecode via `go get github.com/Jnoack331/cube-manager` or `git clone git@github.com:Jnoack331/cube-manager.git`
2. Go into the sourcecode folder.
3. Run `go mod download` to fetch the dependencies.
3. Run `go build` to build the executable.

### Purpose
This repository is a Go devcontainer setup for VS Code. I use it to set up a Go devlopment environment without installing anythign directly to my machine. It can be easily expanded by adding other services (i.e. web server, database, etc) to the `docker-compose.yaml` file to get an application up and running. 

### System Requirements
This requires Docker and VS Code. It has been developed and used on Apple Silicon Macs although it should work seamlessly with other hardware. 

### Customizations
#### Shell 
I use the zsh shell and [Oh-My-Zsh](https://ohmyz.sh/) with the [agnoster](https://github.com/ohmyzsh/ohmyzsh/wiki/Themes#agnoster) theme and some plugins. 

#### VS Code Extensions
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Git Extension Pack](https://marketplace.visualstudio.com/items?itemName=donjayamanne.git-extension-pack)
- [Git Graph](https://marketplace.visualstudio.com/items?itemName=mhutchie.git-graph)
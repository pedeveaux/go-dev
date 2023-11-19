FROM golang:1.21

# Install Oh-My-Zsh with my theme and a few plugins
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- -t agnoster \
 -p git \
 -p https://github.com/zsh-users/zsh-autosuggestions \
 -p https://github.com/zsh-users/zsh-completions \
 -p https://github.com/zsh-users/zsh-syntax-highlighting \

# Set Zsh as the default shell
SHELL ["/bin/zsh", "-1"]

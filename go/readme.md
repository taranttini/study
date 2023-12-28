# GO CONFIG

sample 

PUT_PATH_TO_GO_BIN = /home/tarantini/program

## on bash

`~/.bashrc`

`export PATH=$PATH:/PUT_PATH_TO_GO_BIN/go/bin`

## on zsh

`~/.zshrc`

`export PATH=$PATH:/PUT_PATH_TO_GO_BIN/go/bin`

## on fish

`~/.config/fish/config.fish`

`set -U fish_user_paths $fish_user_paths /PUT_PATH_TO_GO_BIN/go/bin`

inside if content

# run

go run `PATH/file_name`


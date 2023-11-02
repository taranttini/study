systemctl --user enable --now podman.socket

wsl -d podman-machine-default enterns su user

podman machine set --rootful

su user
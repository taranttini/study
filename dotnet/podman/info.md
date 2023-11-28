systemctl --user enable --now podman.socket

wsl -d podman-machine-default enterns su user

podman machine set --rootful

su user


# dotnet commandos

dotnet tool install --global dotnet-ef --version 7.*

# ferramenta 

dotnet ef migrations add initialmigration

~/.dotnet/tools/dotnet-ef migrations add initialmigration


FN2 + 9 or 0 to change keyboard colors ou themes
# Optional: activate your WSL environment
wsl -d Ubuntu

# Install ubuntu requirements:
sudo apt-get update && sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev npm pkg-config && sudo snap install go --classic

# Install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion" # This loads nvm bash_completion
nvm install --lts
nvm use --lts

# Install wails and source it's path
go install github.com/wailsapp/wails/v2/cmd/wails@latest && export PATH=$PATH:$HOME/go/bin && source ~/.bashrc

# Install go requirements
go mod tidy && go mod download && go get -u && go mod verify

# Run the build (for Windows include: --platform windows/amd64)
wails build

# Run the binary
chmod +x build/bin/Slender && ./build/bin/Slender

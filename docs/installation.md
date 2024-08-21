# Installation guide

## Linux (+ WSL)

### Optional: activate your WSL environment
    ```
    wsl -d Ubuntu
    ```

### Install ubuntu requirements:
    ```
    sudo apt-get update && sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev npm pkg-config && sudo snap install go --classic
    ```

### Install nvm
    ```
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" ### This loads nvm
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion" ### This loads nvm bash_completion
    nvm install --lts
    nvm use --lts
    ```

### Install wails and source it's path
    ```
    go install github.com/wailsapp/wails/v2/cmd/wails@latest && export PATH=$PATH:$HOME/go/bin && source ~/.bashrc
    ```


### Install go requirements
    ```
    go mod tidy && go mod download && go get -u && go mod verify && go build -o Slender.exe
    ```

### Install npm requirements
    ```
    cd frontend && npm install --save-dev @tsconfig/svelte && cd ..
    ```


### Run the build (for Windows include: --platform windows/amd64)
    ```
    mkdir -p build/windows && wails build -platform windows/amd64 -o build/windows/Slender.exe
    mkdir -p build/linux && wails build -platform linux/amd64 -o build/linux/Slender
    ```


### Run the binary
    ```
    chmod +x build/bin/Slender && ./build/bin/Slender
    ```


### To create a new release
    ```
    git tag -a v0.0.1 -m "Release version 0.0.1"
    git push origin v0.0.1
    ```

----

## Windows

### Download and run the nvm-windows installer
    ```
    Invoke-WebRequest -Uri "https://github.com/coreybutler/nvm-windows/releases/latest/download/nvm-setup.zip" -OutFile "nvm-setup.zip"
    Expand-Archive -Path "nvm-setup.zip" -DestinationPath "$env:TEMP"
    Start-Process -FilePath "$env:TEMP\nvm-setup.exe" -ArgumentList "/S" -Wait
    ```

### Install LTS version of Node.js
    ```
    nvm install lts
    nvm use lts
    ```

### Install Go from the official website and then:
    ```
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```

### Add Go binaries to PATH (PowerShell profile)
    ```
    $goPath = "$env:USERPROFILE\go\bin"
    [Environment]::SetEnvironmentVariable("PATH", "$env:PATH;$goPath", [EnvironmentVariableTarget]::User)
    ```

### Reload PowerShell profile to include new PATH
    ```
    . $PROFILE
    ```
### Install Go Requirements
    ```
    go mod tidy
    go mod download
    go get -u
    go mod verify
    ```
### Install npm Requirements
    ```
    cd frontend
    npm install --save-dev @tsconfig/svelte
    cd ..
    ```

### Create directories and build
    ```
    New-Item -Path "build\windows" -ItemType Directory -Force
    wails build -platform windows/amd64 -o build\windows\Slender.exe
    New-Item -Path "build\linux" -ItemType Directory -Force
    wails build -platform linux/amd64 -o build\linux\Slender
    ```

### Run the binary (assuming `Slender.exe` is built)
    ```
    .\build\windows\Slender.exe
    ```

#!/bin/bash

# Define paths and default values
BIN_PATH=${1:-"build/bin/Slender"}
LOG_DIR="logs"

# Create the logs directory if it doesn't exist
mkdir -p "$LOG_DIR"

# Run the build
echo -e "\e[01;32m Building the binary... \e[0m"
wails build

# Make the binary executable
chmod +x "$BIN_PATH"

# Loop to start the launcher and handle logging
while true; do
    sleep 10
    echo -e "\e[01;32m Attempting to start the launcher... \e[0m"

    # Create timestamp for log filenames
    TIMESTAMP=$(date +"%F_%H-%M-%S")
    
    # Redirect stderr (warnings) to a separate file and stdout to a log file
    "$BIN_PATH" 2> "$LOG_DIR/${TIMESTAMP}-warnings.log" | tee "$LOG_DIR/${TIMESTAMP}.log"
    
    EXIT_CODE=$?
    if [[ $EXIT_CODE -ne 0 ]]; then
        echo -e "\e[01;31m launcher failed with exit code $EXIT_CODE \e[0m"
    else
        echo -e "\e[01;32m launcher started successfully \e[0m"
    fi
    
    # Check if 'q' was pressed
    read -t 1 -N 1 -r input
    if [[ "$input" == "q" ]]; then
        echo -e "\e[01;33m Stopping the launcher \e[0m"
        break
    fi
done

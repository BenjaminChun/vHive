#!/bin/bash

# Source the setup_node.sh script
source ./scripts/cloudlab/setup_node.sh

# Check if sourcing setup_node.sh was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to source setup_node.sh."
    exit 1
fi

# Change to the ctriface directory
cd ./ctriface

# Check if changing directory was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to change to the ctriface directory."
    exit 1
fi

# Run the make test command
make test

# Check if make test was successful
if [ $? -eq 0 ]; then
    echo "Tests in ctriface ran successfully."
else
    echo "Error: Failed to run tests in ctriface."
    exit 1
fi

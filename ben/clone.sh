#!/bin/bash

# Get the repository URLs from the command line arguments
repo_url1="https://github.com/BenjaminChun/vHive.git"
repo_url2="https://github.com/vhive-serverless/firecracker-containerd.git"
branch_name="ssh-test"

# Clone the first repository
git clone "$repo_url1"

# Check if the cloning was successful
if [ $? -eq 0 ]; then
    echo "First repository cloned successfully. Check the directory."
else
    echo "Error: Failed to clone the first repository."
    exit 1
fi

# Switch to ssh-test branch
cd ./vHive && git checkout "$branch_name" && cd ~

# Check if the checkout branch was successful
if [ $? -eq 0 ]; then
    echo "Switched to $branch_name branch"
else
    echo "Error in switching branch"
    exit 1
fi

# Clone the second repository
git clone "$repo_url2"

# Check if the cloning was successful
if [ $? -eq 0 ]; then
    echo "Second repository cloned successfully. Check the directory."
else
    echo "Error: Failed to clone the second repository."
    exit 1
fi

#!/usr/bin/env bash
# Strict mode
set -euo pipefail
IFS=$'\n\t'

# Variables
DAY_NAME=$1
INPUT_DAY_FOLDER="inputs/day-$DAY_NAME"
GO_FOLDER="go"

# Make input files
mkdir "$INPUT_DAY_FOLDER"
touch "$INPUT_DAY_FOLDER/simple.txt"
touch "$INPUT_DAY_FOLDER/full.txt"
echo "Input files created"

# Make Go files
mkdir "$GO_FOLDER/day$DAY_NAME"
cp "$GO_FOLDER/day03/main.go" "$GO_FOLDER/day$DAY_NAME/main.go"
cp "$GO_FOLDER/day03/main_test.go" "$GO_FOLDER/day$DAY_NAME/main_test.go"
echo "Go files created"
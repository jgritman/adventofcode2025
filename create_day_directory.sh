#!/bin/bash

# Check if a number is provided as an argument
if [ -z "$1" ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

# Get the number and format it as dayXX
number=$1
formatted_number=$(printf "%02d" "$number")
directory_name="day$formatted_number"

# Create the directory
mkdir -p "$directory_name"
mkdir -p "$directory_name/part1"
mkdir -p "$directory_name/part2"

# Create the required files in the directory
cp "template.go" "$directory_name/part1/main.go"
touch "$directory_name/part2/main.go"

git add --all

touch "$directory_name/sample_input.txt"

url=https://adventofcode.com/2025/day/$1/input
curl -o "$directory_name/input.txt" "$url" -H 'Cookie session=53616c7465645f5f60874cd0b6c52d3a789ccae497f6cac1aa8c80757256d34b4f9cdefabda5079d297951e2731a549c7767dd42de102ba59e6f455dd68a58da'

cd $directory_name
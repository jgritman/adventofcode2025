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
curl -o "$directory_name/input.txt" "$url" -H 'Cookie: _ga=GA1.2.1896476050.1764532081; _gid=GA1.2.1819159324.1764532081; session=53616c7465645f5fe7dd26baff672b211fd9f760ef5e917fc2da9981797434ed7f50f40ace480e5111d7f0649255a78305d8c1f41952692cea25bc5de8dcd694; _ga_MHSNPJKWC7=GS2.2.s1764534538$o2$g1$t1764534548$j50$l0$h0'

cd $directory_name

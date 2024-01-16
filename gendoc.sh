#!/bin/bash

# overwrite the headers
cp README_HEADER.md README.md

folders=$(find "." -maxdepth 1 -type d ! -name '.*')

for folder in $folders; do
  echo "Generating docs for $folder"

  make -s -C "$folder" doc

  file="$folder/README.md"
  if [ ! -f "$file" ]; then
    echo "No README.md found in $folder"
    continue
  fi

  title=$(sed -n '1p' "$file")
  title=${title#*# }
  content=$(sed -n '2,/^$/p' "$file")
  # content=$(sed -n '2,/^##.*$/p' "$file")

  echo "- [$title]($file):$content" >>README.md
done

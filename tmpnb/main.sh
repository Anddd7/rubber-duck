# create a temp folder and change into it
tmpnb() {
  current_time=$(date +"%Y%m%d_%H%M%S")
  folder_name=".tmpnb_$current_time"

  mkdir "$folder_name"

  cd "$folder_name"

  echo "Created temp notebook: $folder_name and changed into it."
}

tmpf() {
  current_time=$(date +"%H%M%S")
  file_name="tmpf_$current_time"
  format=$1

  if [ -n "$format" ]; then
    file_name="$file_name.$format"
  fi

  # file_name=$(mktemp $file_name)
  cat - >"$file_name"

  echo "Output written to $file_name"
}

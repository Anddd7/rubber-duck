# @file tempfile
# @brief Create temp notebook(folder) and file
# @description quick create file/folder

# @description create a temp folder to store your works
#
# @example
#   tmpnb
#
# @stdout You'll cd to the new temp folder
tmpnb() {
  current_time=$(date +"%Y%m%d_%H%M%S")
  folder_name=".tmpnb_$current_time"

  mkdir "$folder_name"

  cd "$folder_name"

  echo "Created temp notebook: $folder_name and changed into it."
}

# @description collect the stdin to a temp file
#
# @example
#   echo "test: a" | tmpf
#   echo "test: a" | tmpf yaml
#
# @stdout temp file path
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

# @file base64 shortcuts
# @brief Some shortcuts for base64 encoding and decoding.
# @description
#   * en/decode string
#   * en/decode kubernetes secret file

# @description base64 encode with pipe
#
# @example
#   b64 your-string
#
# @arg $1 string A value to encode
#
# @stdout encoded string
b64() { echo -n "$1" | base64 -w 0; }

# @description base64 decode with pipe
#
# @example
#   b64d your-string
#
# @arg $1 string A value to dencode
#
# @stdout decoded string
b64d() { echo -n "$1" | base64 -d; }

# @description encode the kubernetes secret data
#
# @example
#   b64k input-file output-file
#   b64k8s input-file output-file
#
# @arg
#   $1 src file
#   $2 dest file contains encoded data
b64k8s() {
  local input_file="$1"
  local output_file="$2"

  if [ -z "$input_file" ] || [ -z "$output_file" ]; then
    echo "Usage: encrypt_secret <input_file> <output_file>"
    return 1
  fi

  if [ ! -f "$input_file" ]; then
    echo "Input file not found: $input_file"
    return 1
  fi

  # 提取data部分，进行Base64加密
  cat "$input_file" | yq eval -P '.data |= with_entries(.value |= @base64)' - >"$output_file"

  echo "Secret encrypted and saved to: $output_file"
}

# @description dencode the kubernetes secret data
#
# @example
#   b64dk input-file output-file
#   b64dk8s input-file output-file
#
# @arg
#   $1 src file contains encoded data
#   $2 dest file
b64dk8s() {
  local input_file="$1"
  local output_file="$2"

  if [ -z "$input_file" ] || [ -z "$output_file" ]; then
    echo "Usage: decrypt_secret <input_file> <output_file>"
    return 1
  fi

  if [ ! -f "$input_file" ]; then
    echo "Input file not found: $input_file"
    return 1
  fi

  cat "$input_file" | yq eval -P '.data |= with_entries(.value |= @base64d)' - >"$output_file"

  echo "Secret decrypted and saved to: $output_file"
}

alias b64k="b64k8s"
alias b64dk="b64dk8s"

b64() { echo -n "$1" | base64 -w 0; }
b64d() { echo -n "$1" | base64 -d; }

alias b64k="b64k8s"
alias b64dk="b64dk8s"

# encrypt k8s secret yaml
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

# decrypt k8s secret yaml
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

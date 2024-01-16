# @file kw: kubectl wrapper
# @brief A simple wrapper for kubectl to make it easier to use.
# @description quick export yaml file

# @description provide some shortcuts for your frequently used kubectl command, support kubectl_completion (for zsh)
# 
# @example
#   # export resource to a temp yaml file
#   kw get pod your-pod -oy
#   > kubectl get pod your-pod -o yaml > pod_your-pod_112422.yaml
# 
# @arg $@ kubectl command
kw() {
  local args=("$@")
  local out_args=()
  local out_args_suffix=()

  local resource_type=""
  local resource_name=""

  # set resource type
  if [[ $1 == "logs" ]]; then
    resource_type="pod"
  else
    resource_type="$2"
  fi

  # set resource name if previous command is not start with -
  for ((i = 1; i <= $#; i++)); do
    # echo "$i: ${args[$i]}"
    if [[ ${args[$i - 1]} != -* && ${args[$i]} != -* ]]; then
      resource_name="${args[$i]}"
    fi
  done

  for ((i = 1; i <= $#; i++)); do
    if [[ ${args[$i]} == "-oy" ]]; then
      local timestamp=$(date +%H%M%S)
      local filename="${resource_type}_${resource_name}_${timestamp}.yaml"

      out_args+=("-o" "yaml")
      out_args_suffix+=(">" "$filename")
    else
      out_args+=("${args[$i]}")
    fi
  done

  echo "kw: kubectl ${out_args[@]} ${out_args_suffix[@]}"
  eval "kubectl ${out_args[*]} ${out_args_suffix[*]}"
}

compdef _kubectl kw

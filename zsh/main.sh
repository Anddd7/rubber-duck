# @file zsh tools
# @brief Misc tools for zsh
# @description clear color tags

# @description to remove the color codes from zsh shell output
#
# @example
#   nocolor your-zsh.log
nocolor() {
  sed -i 's/\[[0-9;]*m//g' "$1"
}

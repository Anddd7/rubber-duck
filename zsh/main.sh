# remove the color codes from zsh shell output
nocolor() {
  sed -i 's/\[[0-9;]*m//g' "$1"
}

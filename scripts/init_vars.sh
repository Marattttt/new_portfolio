# This script sets these variables to use as absolute paths for other scripts 
PROTO_DIR=
SERVICES_DIR=

set -euo pipefail

current_path=$(pwd)
last_dir=$(basename "$current_path")

# Use a case statement to match the last segment against acceptable items
case "$last_dir" in
  "new_portfolio")
    PROTO_DIR="$current_path/proto"
    SERVICES_DIR="$current_path/services"
    ;;

  "scripts")
    PROTO_DIR="$(dirname current_path)/proto"
    SERVICES_DIR="$(dirname current_path)/services"
    ;;

  *)
    echo "Error: The last segment of the current directory ('$last_dir') is not new_portfolio or scripts"
    exit 1
    ;;
esac

set +euo pipefail

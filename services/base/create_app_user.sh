#! /usr/bin/sh

usage='
create_app_and_run_users.sh - creates a user named app_user with set to 
APP_USER_PASS or no password if the variable is unset
The use is is unset restricted to RUNTIME_DIR as well as other restrictions

Needs to be executed as root
'

# Exit on any error / non-zero exit code
set -eo

# Check for root user (required to access user management)
if [ "$(id -u)" -ne 0 ]; then 
	echo 'Need root access for user manipulation'
	exit 1
fi;


if [ -z "$RUNTIME_DIR" ]; then 
	RUNTIME_DIR=/app/runtime
	echo "Variable RUNTIME_DIR is not set; default to $RUNTIME_DIR"
else 
	echo "Using runtime dir $RUNTIME_DIR"
fi;

# Create app user with no home directory resticted to $RUNTIME_DIR
useradd -m  -d "$RUNTIME_DIR" app_user
if [ -z "$APP_USER_PASS" ]; then 
	echo 'app_user set up with no password'
	passwd -d app_user
else 
	echo 'app_user set up with a password'
	echo "$APP_USER_PASS" | passwd --stdin app_user
fi

# Verify user creation
if id "app_user" &>/dev/null; then
    echo "User app_user created successfully"
else
    echo "Error creating user app_user"
    exit 1
fi

# Ensure dir is created
mkdir -p "$RUNTIME_DIR"

# Make app_user the ownder of the directory
chown app_user:app_user "$RUNTIME_DIR"

# Configure PATH and go back to dir where the script was started
cd "$RUNTIME_DIR"
echo "cd $RUNTIME_DIR 
export PATH=$RUNTIME_DIR:/usr/local/bin" > .profile

echo 'Wrote app_user .profile file:'
cat .profile
cd -

# All permisisons for app_user and RW for everyone else
rwx_rw_rw=766
chmod "$rwx_rw_rw" "$RUNTIME_DIR"

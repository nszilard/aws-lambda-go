#!/usr/bin/env bash
set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ];
do
	SOURCE="$(readlink "$SOURCE")";
done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

PACKAGES=$(go list ./... | grep -vF function/internal )
COVERAGE=coverage
COVER_MODE=count
COVER_PROFILE="$COVERAGE/coverage.out"

# Creating temp coverage folder
mkdir -p "$COVERAGE"

# Looping through all packages
for pkg in $PACKAGES; do
	f="$COVERAGE/$(echo $pkg | tr / -).cover"
	go test -parallel=20 -covermode="$COVER_MODE" -coverprofile="$f" "$pkg"
done

# Building coverage file
echo "mode: $COVER_MODE" > "$COVER_PROFILE"
grep -h -v "^mode:" "$COVERAGE"/*.cover >> "$COVER_PROFILE"

# Showing coverage in a browser
if [[ -n "$SHOW_HTML" ]];
then
	go tool cover -html=$COVER_PROFILE
fi

# Cleaning up coverage folder
rm -rf "$COVERAGE"
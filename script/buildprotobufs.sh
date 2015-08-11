#!/bin/bash

# Basic path locations
ROOT=$(cd $(dirname -- "$0" ) && cd .. && pwd)
MESSAGEPATH=${ROOT}/proto

# SRCPATH is the path to our src directory - everything from here is fully qualified
# This depends on your storing your code in your GOPATH
# eg. xxx/github.com/mondough/phosphor
SRCPATH=$(cd ${ROOT}/../../.. && pwd)

# Cakes are important. and delicious. and should be given out for success.
function dispatchCake() {
	printf "\n            \033[1;33m*\033[0m  \033[1;33m*\033[0m  \033[1;33m*\033[0m             \n"
	printf "           \033[1;33m*\033[0m\033[0;31m|\033[0m_\033[1;33m*\033[0m\033[0;31m|\033[0m_\033[1;33m*\033[0m\033[0;31m|\033[0m_\033[1;33m*\033[0m           \n"
	printf "       .-'\`\033[0;31m|\033[0m  \033[0;31m|\033[0m  \033[0;31m|\033[0m  \033[0;31m|\033[0m\`'-.       \n"
	printf "       |\`-............-'|       \n"
	printf "       |                |       \n"
	printf "       \   _  .-.   _   /       \n"
	printf "     ,-|'-' '-'  '-' '-'|-,     \n"
	printf "   /\`  \._            _./  \`\   \n"
	printf "   '._    \`\"\"\"\"\"\"\"\"\"\`    _.'\n"
	printf "     \`''--..........--''\`       \n\n"
	printf "        \033[1;5;7;32m GREAT SUCCESS! \033[0m\n\n"
	printf "\n\n"
}

# Show which protobufs were found
printf "\nLocating protobufs...\n"
find ${MESSAGEPATH} -name '*.proto' -exec echo {} \;
echo ""

# Clean out current protos
find ${MESSAGEPATH} -name '*.pb.go' | xargs rm -f

# Try to rebuild all the things
echo "Generating Go protobuf classes..."
find $MESSAGEPATH -name '*.proto' -exec protoc -I${SRCPATH} --go_out=${SRCPATH} {} \;
printf "Complete\n\n"

# GREAT SUCCESS
dispatchCake

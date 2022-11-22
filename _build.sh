#!/bin/bash

# go install fyne.io/fyne/v2/cmd/fyne@latest
#fyne package -os windows -icon fav.png


# go install github.com/fyne-io/fyne-cross@latest
# fyne-cross <command> [options]
# The commands are:
# 	darwin        Build and package a fyne application for the darwin OS
# 	linux         Build and package a fyne application for the linux OS
# 	windows       Build and package a fyne application for the windows OS
# 	android       Build and package a fyne application for the android OS
# 	ios           Build and package a fyne application for the iOS OS
# 	freebsd       Build and package a fyne application for the freebsd OS
# 	version       Print the fyne-cross version information
# Use "fyne-cross <command> -help" for more information about a command.
# fyne-cross windows -arch=amd64,386
rm -rf ./fyne-cross
sudo docker image prune -af
sudo ./cross windows -arch=*
sudo ./cross linux -arch=*
#sudo ./cross darwin -arch=*
#sudo docker image prune -af
#sudo ./cross android
#sudo ./cross freebsd
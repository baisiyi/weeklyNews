#!/bin/bash

# go env -w GO111MODULE=on
# go env -w GOPROXY=https://goproxy.cn,direct

go build -o weeklyNews

chmod +x weeklyNews

./weeklyNews > /dev/null 2>&1 &

# sudo yum install epel-release -y
# sudo rpm --import http://li.nux.ro/download/nux/RPM-GPG-KEY-nux.ro
# sudo rpm -Uvh http://li.nux.ro/download/nux/dextop/el7/x86_64/nux-dextop-release-0-5.el7.nux.noarch.rpm
# sudo yum install ffmpeg ffmpeg-devel -y
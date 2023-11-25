#!/bin/bash

cd /data/release/weeklyNews/d125c8f0c6c74a849988c5457178cd3b/1200 || exit

mkdir -p /data/release/weeklyNews/d125c8f0c6c74a849988c5457178cd3b/result

files=$(ls *.ts | sort -V | tr '\n' '|')

ffmpeg -i "concat:${files%|}" -c copy /data/release/weeklyNews/d125c8f0c6c74a849988c5457178cd3b/result/1200_weeklyNews.mp4
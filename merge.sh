#!/bin/bash


files=$(ls *.ts | sort -V | tr '\n' '|')

ffmpeg -i "concat:${files%|}" -c copy weeklyNews.mp4
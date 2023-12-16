#!/bin/bash

while read -r line; do
  # 获取最后一个斜杠和冒号之间的字符作为镜像名
  image_name=$(echo "$line" | awk -F '[/:]' '{print $(NF-1)}')

  # 创建文件夹（如果不存在）
  mkdir -p "./$image_name"

  # 创建 Dockerfile
  echo "FROM $line" > "./$image_name/Dockerfile"
done < "$1"
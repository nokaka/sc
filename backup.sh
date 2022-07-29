#!/bin/sh

# 先判断是否有旧的备份，有先删除
if [  -f /temp/golden-oldies* ]; then
	rm -rf /temp/golden-oldies*
    echo 已删除旧版本
fi

# 打包备份
zip -r ./golden-oldies.zip /etc/config/*
echo 打包成功
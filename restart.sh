#!/bin/bash

# 查找正在运行的程序的 PID
PID=$(ps aux | grep plus-one | grep -v grep | awk '{print $2}')

# 如果 PID 不为空，则终止进程
if [ -n "$PID" ]; then
    echo "停止进程"
    ps aux | grep plus-one
    kill $PID
else
    echo "无进程"
fi

nohup ./plus-one &
echo "启动进程"
ps aux | grep plus-one

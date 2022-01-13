#!/bin/bash

export REMOTE_MODE=false

mkdir -p $POSTGRE_HOME/logs
nohup $POSTGRE_HOME/mate/postgre_mate >>$POSTGRE_HOME/logs/postgre_mate.stdout.log 2>>$POSTGRE_HOME/logs/postgre_mate.stderr.log


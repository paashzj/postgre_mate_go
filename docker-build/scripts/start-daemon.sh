#!/bin/bash

export REMOTE_MODE=false

nohup $POSTGRE_HOME/mate/postgre_mate >>$POSTGRE_HOME/postgre_mate.stdout.log 2>>$POSTGRE_HOME/postgre_mate.stderr.log
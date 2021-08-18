#!/bin/bash

export REMOTE_MODE=false

nohup $POSTGRE_HOME/mate/postgre_mate >$POSTGRE_HOME/postgre_mate.log 2>$POSTGRE_HOME/postgre_mate_error.log
#!/bin/bash

# postgre need to init as sh user
POSTGRE_DATA=$POSTGRE_HOME/data
POSTGRE_MATE=$POSTGRE_HOME/mate
mkdir -p $POSTGRE_DATA
/usr/pgsql-13/bin/initdb -U postgres -D $POSTGRE_DATA --no-locale --encoding=UTF8
cp $POSTGRE_MATE/conf/postgresql.conf  $POSTGRE_DATA/postgresql.conf
cp $POSTGRE_MATE/conf/pg_hba.conf  $POSTGRE_DATA/pg_hba.conf
/usr/pgsql-13/bin/pg_ctl -D $POSTGRE_DATA start
psql -U postgres -f $POSTGRE_MATE/sql/init_database.sql
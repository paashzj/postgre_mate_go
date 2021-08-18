#!/bin/bash

POSTGRE_MATE=$POSTGRE_HOME/mate
/usr/pgsql-13/bin/initdb -U postgres -D /opt/data/postgre --no-locale --encoding=UTF8
cp $POSTGRE_MATE/conf/postgresql.conf  /opt/data/postgre/postgresql.conf
cp $POSTGRE_MATE/conf/pg_hba.conf  /opt/data/postgre/pg_hba.conf
/usr/pgsql-13/bin/pg_ctl -D /opt/data/postgre start
psql -U postgres -f $POSTGRE_MATE/sql/init_database.sql
#!/bin/bash

chown -R sh:sh /opt/sh

chown -R sh:sh /usr/pgsql-13
mkdir -p /opt/data/postgre
chown -R sh:sh /opt/data/postgre
chown -R sh:sh /var/run/postgresql
su sh -c 'bash $POSTGRE_HOME/mate/scripts/start-postgre-sh.sh'
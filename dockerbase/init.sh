#!/bin/bash

mysql -u root -e "CREATE DATABASE gomicro "

R=/data/deploy/gomicro
cd $R
bash build_local.sh all
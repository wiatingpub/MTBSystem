#!/bin/bash

mysql -u root -e "CREATE DATABASE mtbsystem "

R=/data/deploy/mtbsystem
cd $R
bash build_local.sh all
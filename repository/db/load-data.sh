#!/bin/bash

psql -U admin -d drones -p 5432 -a -q -f /home/shimaa/go/src/drones/repository/db/seeds.sql
psql -U admin -d test_drones -p 5432 -a -q -f /home/shimaa/go/src/drones/repository/db/seeds.sql

#!/bin/bash

git submodule update kernel --remote
cd $(dirname $0)/../kernel && make -j$NUM_JOBS

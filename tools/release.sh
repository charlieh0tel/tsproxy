#!/bin/bash

set -o errexit
set -o xtrace

cp --backup=t -p tsproxy released

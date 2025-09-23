#!/bin/bash

set -o errexit
set -o xtrace

systemctl --user restart tsproxy@pdp1153.service
systemctl --user restart tsproxy@pdp1183.service

#!/usr/bin/env bash
/root/turingchain -f /root/turingchain.toml &
# to wait nginx start
sleep 15
/root/turingchain -f "$PARAFILE"

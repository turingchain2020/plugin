#!/usr/bin/env bash

BUILD=$(cd "$(dirname "$0")" && cd ../../../../build && pwd)
echo "$BUILD"

cd "$BUILD" || return

seed=$(./turingchain-cli seed generate -l 0)
echo "$seed"

./turingchain-cli seed save -p trc123456 -s "$seed"
sleep 1
./turingchain-cli wallet unlock -p trc123456
sleep 1
./turingchain-cli account list

## account
#genesis -- 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli account import_key -k 3990969DF92A5914F7B71EEB9A4E58D6E255F32BF042FEA5318FC8B3D50EE6E8 -l genesis

#A -- 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4
./turingchain-cli account import_key -k 0x6da92a632ab7deb67d38c0f6560bcfed28167998f6496db64c258d5e8393a81b -l A

#B -- 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR
./turingchain-cli account import_key -k 0x19c069234f9d3e61135fefbeb7791b149cdf6af536f26bebb310d4cd22c3fee4 -l B

#C -- 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k
./turingchain-cli account import_key -k 0x7a80a1f75d7360c6123c32a78ecf978c1ac55636f87892df38d8b85a9aeff115 -l C

#D -- 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs
./turingchain-cli account import_key -k 0xcacb1f5d51700aea07fca2246ab43b0917d70405c65edea9b5063d72eb5c6b71 -l D

## config token
./turingchain-cli send config config_tx -c token-finisher -o add -v 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
sleep 1
./turingchain-cli config query -k token-finisher

./turingchain-cli send config config_tx -c token-blacklist -o add -v TRC -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
sleep 1
./turingchain-cli config query -k token-blacklist

## 10亿
./turingchain-cli send token precreate -f 0.001 -i "test ccny" -n "ccny" -a 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs -p 0 -s CCNY -t 1000000000 -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
sleep 1
./turingchain-cli token precreated
./turingchain-cli send token finish -s CCNY -f 0.001 -a 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
sleep 1
./turingchain-cli token created
./turingchain-cli token balance -a 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs -s CCNY -e token

## transfer trc
./turingchain-cli send coins transfer -a 10000 -t 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4 -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send coins transfer -a 10000 -t 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send coins transfer -a 10000 -t 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send coins transfer -a 10000 -t 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli account list

## send trc to execer， 每人10000 trc
./turingchain-cli send coins send_exec -e exchange -a 10000 -k 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4
./turingchain-cli send coins send_exec -e exchange -a 10000 -k 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR
./turingchain-cli send coins send_exec -e exchange -a 10000 -k 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k
./turingchain-cli send coins send_exec -e exchange -a 10000 -k 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs
echo "account balance in execer"
./turingchain-cli account balance -e exchange -a 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4
./turingchain-cli account balance -e exchange -a 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR
./turingchain-cli account balance -e exchange -a 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k
./turingchain-cli account balance -e exchange -a 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs

##transfer token，每人2亿 ccny
./turingchain-cli send token transfer -a 200000000 -s CCNY -t 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4 -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send token transfer -a 200000000 -s CCNY -t 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send token transfer -a 200000000 -s CCNY -t 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs
./turingchain-cli send token transfer -a 200000000 -s CCNY -t 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs -k 1CbEVT9RnM5oZhWMj4fxUrJX94VtRotzvs

## send token to excer
./turingchain-cli send token send_exec -a 200000000 -e exchange -s CCNY -k 1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4
./turingchain-cli send token send_exec -a 200000000 -e exchange -s CCNY -k 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR
./turingchain-cli send token send_exec -a 200000000 -e exchange -s CCNY -k 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k
./turingchain-cli send token send_exec -a 200000000 -e exchange -s CCNY -k 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs
echo "token balance in execer"
./turingchain-cli token balance -e exchange -s CCNY -a "1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4 1JRNjdEqp4LJ5fqycUBm9ayCKSeeskgMKR 1NLHPEcbTWWxxU3dGUZBhayjrCHD3psX7k 1MCftFynyvG2F4ED5mdHYgziDxx6vDrScs"

#!/usr/bin/env bash
# shellcheck disable=SC2128
# shellcheck source=/dev/null
set -e
set -o pipefail

MAIN_HTTP=""
GAME_ID=""

source ../dapp-test-common.sh

pokerbull_PlayRawTx() {
    tx=$(curl -ksd '{"method":"Turingchain.CreateTransaction","params":[{"execer":"pokerbull","actionName":"Play","payload":{"gameId":"pokerbull-abc", "value":"1000000000", "round":1}}]}' ${MAIN_HTTP} | jq -r ".result")
    turingchain_SignAndSendTxWait "$tx" "0x0316d5e33e7bce2455413156cb95209f8c641af352ee5d648c647f24383e4d94" ${MAIN_HTTP} "$FUNCNAME"
}

pokerbull_QuitRawTx() {
    tx=$(curl -ksd '{"method":"Turingchain.CreateTransaction","params":[{"execer":"pokerbull","actionName":"Quit","payload":{"gameId":"'$GAME_ID'"}}]}' ${MAIN_HTTP} | jq -r ".result")
    turingchain_SignAndSendTxWait "$tx" "0x0316d5e33e7bce2455413156cb95209f8c641af352ee5d648c647f24383e4d94" ${MAIN_HTTP} "$FUNCNAME"
}

pokerbull_ContinueRawTx() {
    tx=$(curl -ksd '{"method":"Turingchain.CreateTransaction","params":[{"execer":"pokerbull","actionName":"Continue","payload":{"gameId":"'$GAME_ID'"}}]}' ${MAIN_HTTP} | jq -r ".result")
    turingchain_SignAndSendTxWait "$tx" "0xa26038cbdd9e6fbfb85f2c3d032254755e75252b9edccbecc16d9ba117d96705" ${MAIN_HTTP} "$FUNCNAME"
}

pokerbull_StartRawTx() {
    tx=$(curl -ksd '{"method":"Turingchain.CreateTransaction","params":[{"execer":"pokerbull","actionName":"Start","payload":{"value":"1000000000", "playerNum":"2"}}]}' ${MAIN_HTTP} | jq -r ".result")
    req='{"method":"Turingchain.DecodeRawTransaction","params":[{"txHex":"'"$tx"'"}]}'
    turingchain_Http "$req" ${MAIN_HTTP} '(.result.txs[0].execer != null)' "$FUNCNAME"
    turingchain_SignAndSendTx "$tx" "0x0316d5e33e7bce2455413156cb95209f8c641af352ee5d648c647f24383e4d94" ${MAIN_HTTP}
    GAME_ID=$RAW_TX_HASH
    turingchain_BlockWait 1 "${MAIN_HTTP}"
}

pokerbull_QueryResult() {
    req='{"method":"Turingchain.Query","params":[{"execer":"pokerbull","funcName":"QueryGameByID","payload":{"gameId":"'$GAME_ID'"}}]}'
    resok='(.result.game.gameId == "'"$GAME_ID"'")'
    turingchain_Http "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"

    req='{"method":"Turingchain.Query","params":[{"execer":"pokerbull","funcName":"QueryGameByAddr","payload":{"addr":"14VkqML8YTRK4o15Cf97CQhpbnRUa6sJY4"}}]}'
    turingchain_Http "$req" ${MAIN_HTTP} '(.result != null)' "$FUNCNAME"

    req='{"method":"Turingchain.Query","params":[{"execer":"pokerbull","funcName":"QueryGameByStatus","payload":{"status":"3"}}]}'
    turingchain_Http "$req" ${MAIN_HTTP} '(.result != null)' "$FUNCNAME"
}

init() {
    ispara=$(echo '"'"${MAIN_HTTP}"'"' | jq '.|contains("8901")')
    echo "ipara=$ispara"
    if [ "$ispara" == true ]; then
        pokerbull_addr=$(curl -ksd '{"method":"Turingchain.ConvertExectoAddr","params":[{"execname":"user.p.para.pokerbull"}]}' ${MAIN_HTTP} | jq -r ".result")
    else
        pokerbull_addr=$(curl -ksd '{"method":"Turingchain.ConvertExectoAddr","params":[{"execname":"pokerbull"}]}' ${MAIN_HTTP} | jq -r ".result")
    fi

    local main_ip=${MAIN_HTTP//8901/9671}
    turingchain_ImportPrivkey "0x0316d5e33e7bce2455413156cb95209f8c641af352ee5d648c647f24383e4d94" "14VkqML8YTRK4o15Cf97CQhpbnRUa6sJY4" "pokerbull1" "${main_ip}"
    turingchain_ImportPrivkey "0xa26038cbdd9e6fbfb85f2c3d032254755e75252b9edccbecc16d9ba117d96705" "1MuVM87DLigWhJxLJKvghTa1po4ZdWtDv1" "pokerbull2" "$main_ip"

    local pokerbull1="14VkqML8YTRK4o15Cf97CQhpbnRUa6sJY4"
    local pokerbull2="1MuVM87DLigWhJxLJKvghTa1po4ZdWtDv1"

    if [ "$ispara" == false ]; then
        turingchain_applyCoins "$pokerbull1" 12000000000 "${main_ip}"
        turingchain_QueryBalance "${pokerbull1}" "$main_ip"

        turingchain_applyCoins "$pokerbull2" 12000000000 "${main_ip}"
        turingchain_QueryBalance "${pokerbull2}" "$main_ip"
    else
        turingchain_applyCoins "$pokerbull1" 1000000000 "${main_ip}"
        turingchain_QueryBalance "${pokerbull1}" "$main_ip"

        turingchain_applyCoins "$pokerbull2" 1000000000 "${main_ip}"
        turingchain_QueryBalance "${pokerbull2}" "$main_ip"
        local para_ip="${MAIN_HTTP}"
        turingchain_ImportPrivkey "0x0316d5e33e7bce2455413156cb95209f8c641af352ee5d648c647f24383e4d94" "14VkqML8YTRK4o15Cf97CQhpbnRUa6sJY4" "pokerbull1" "$para_ip"
        turingchain_ImportPrivkey "0xa26038cbdd9e6fbfb85f2c3d032254755e75252b9edccbecc16d9ba117d96705" "1MuVM87DLigWhJxLJKvghTa1po4ZdWtDv1" "pokerbull2" "$para_ip"

        turingchain_applyCoins "$pokerbull1" 12000000000 "${para_ip}"
        turingchain_QueryBalance "${pokerbull1}" "$para_ip"
        turingchain_applyCoins "$pokerbull2" 12000000000 "${para_ip}"
        turingchain_QueryBalance "${pokerbull2}" "$para_ip"
    fi

    turingchain_SendToAddress "$pokerbull1" "$pokerbull_addr" 10000000000 ${MAIN_HTTP}
    turingchain_QueryExecBalance "${pokerbull1}" "pokerbull" "$MAIN_HTTP"
    turingchain_SendToAddress "$pokerbull2" "$pokerbull_addr" 10000000000 ${MAIN_HTTP}
    turingchain_QueryExecBalance "${pokerbull2}" "pokerbull" "$MAIN_HTTP"

    turingchain_BlockWait 1 "${MAIN_HTTP}"
}

function run_test() {
    pokerbull_StartRawTx

    pokerbull_ContinueRawTx

    pokerbull_QuitRawTx

    pokerbull_PlayRawTx

    pokerbull_QueryResult
}

function main() {
    turingchain_RpcTestBegin pokerbull
    MAIN_HTTP="$1"
    echo "ip=$MAIN_HTTP"

    init
    run_test
    turingchain_RpcTestRst pokerbull "$CASE_ERR"
}

turingchain_debug_function main "$1"

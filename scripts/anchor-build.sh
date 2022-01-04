#!/bin/bash

# get this scripts directory
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

CONTRACTS=${SCRIPT_DIR}/../contracts

cd ${CONTRACTS}
anchor build

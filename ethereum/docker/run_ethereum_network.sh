#!/bin/bash

# Exit on any errors
set -e

original_dir=$(pwd)
data_dir=~/.local_geth_network

function delete_old_files() {
    rm -rf $data_dir
}

_term() { # Cleanup when Docker Compose sends a SIGTERM
    echo "Cleaning up..."
    killall geth
    delete_old_files
    exit 0
}

trap _term INT # Cleanup child processes and write log on Ctrl+C
trap _term SIGTERM # Cleanup when Docker Compose kills the container

# Delete old nodes
delete_old_files

# Initialize new node
mkdir -p $data_dir
echo "Creating genesis..."
geth init --datadir $data_dir genesis.json

# Start the node
echo "Starting Geth node..."
geth  --dev --http --http.addr "0.0.0.0" --http.port "5050" --http.vhosts "geth,localhost" --http.corsdomain "0.0.0.0" --datadir $data_dir --http.api "eth,net,web3" --dev.period 1 &
node_pid=$!

# Wait for the node to start
until [ -e "$data_dir/geth.ipc" ]
do
    echo 'Waiting on geth.ipc socket to be present...'
    sleep 1
done

echo "Successfully started network."

wait $node_pid


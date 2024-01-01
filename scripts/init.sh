#!/bin/bash
BINARY=eventchaind
CHAIN_DIR=./data
CHAIN_ID="eventchain"
VAL_MNEMONIC="position hat maze shine swallow short web scheme exist hip quiz rice dirt biology iron left shed panel obey oval minor expire company afraid"
OBSERVER_MNEMONIC_1="minor umbrella weather spin hand addict rabbit vocal sample cargo dash cool soup runway divorce slim mansion divide dinosaur useless glue layer prosper stamp"
OBSERVER_MNEMONIC_2="bachelor detail traffic outside copy speak village swing sun mercy genuine meat trim trim cake direct position song garbage bulb gospel fuel loud shield"
OBSERVER_MNEMONIC_3="rail hole desk six scheme verify gown material belt idle inform lesson great airport ahead quote vehicle brown juice close ladder sun guide recycle"

# Stop if it is already running
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAIN_ID &> /dev/null

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAIN_ID 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAIN_ID..."
$BINARY init test --home $CHAIN_DIR/$CHAIN_ID --chain-id=$CHAIN_ID

echo "Adding genesis accounts..."
echo $VAL_MNEMONIC | $BINARY keys add val --home $CHAIN_DIR/$CHAIN_ID --recover --keyring-backend=test
echo $OBSERVER_MNEMONIC_1 | $BINARY keys add observer1 --home $CHAIN_DIR/$CHAIN_ID --recover --keyring-backend=test
echo $OBSERVER_MNEMONIC_2 | $BINARY keys add observer2 --home $CHAIN_DIR/$CHAIN_ID --recover --keyring-backend=test
echo $OBSERVER_MNEMONIC_3 | $BINARY keys add observer3 --home $CHAIN_DIR/$CHAIN_ID --recover --keyring-backend=test

observer1_address=$("$BINARY" --home "$CHAIN_DIR/$CHAIN_ID" keys show observer1 --keyring-backend test -a)
observer2_address=$("$BINARY" --home "$CHAIN_DIR/$CHAIN_ID" keys show observer2 --keyring-backend test -a)
observer3_address=$("$BINARY" --home "$CHAIN_DIR/$CHAIN_ID" keys show observer3 --keyring-backend test -a)

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAIN_ID keys show val --keyring-backend test -a) 100000000000000000000000000stake --home $CHAIN_DIR/$CHAIN_ID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAIN_ID keys show observer1 --keyring-backend test -a) 100000000000000000000000000stake  --home $CHAIN_DIR/$CHAIN_ID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAIN_ID keys show observer2 --keyring-backend test -a) 100000000000000000000000000stake  --home $CHAIN_DIR/$CHAIN_ID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAIN_ID keys show observer3 --keyring-backend test -a) 100000000000000000000000000stake  --home $CHAIN_DIR/$CHAIN_ID

sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAIN_ID/config/app.toml
sed -i -e 's/pruning = "default"/pruning = "everything"/g' $CHAIN_DIR/$CHAIN_ID/config/app.toml
jq ".app_state.oracle.whitelist = [\"$observer1_address\",\"$observer2_address\",\"$observer3_address\"]" "$CHAIN_DIR/$CHAIN_ID/config/genesis.json" > temp.json && mv temp.json "$CHAIN_DIR/$CHAIN_ID/config/genesis.json"

echo "Creating and collecting gentx..."
$BINARY gentx val 1000000000000000000000stake  --home $CHAIN_DIR/$CHAIN_ID --chain-id $CHAIN_ID --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAIN_ID

$BINARY start --log_level trace --home $CHAIN_DIR/$CHAIN_ID --pruning=everything --minimum-gas-prices=0.00001stake
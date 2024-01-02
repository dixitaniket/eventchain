# eventchain

## To start evm and event chain
```
docker-compose up 
```
- event chain already has 3 observer mnemonics whitelisted 

```
OBSERVER_MNEMONIC_1="minor umbrella weather spin hand addict rabbit vocal sample cargo dash cool soup runway divorce slim mansion divide dinosaur useless glue layer prosper stamp"
OBSERVER_MNEMONIC_2="bachelor detail traffic outside copy speak village swing sun mercy genuine meat trim trim cake direct position song garbage bulb gospel fuel loud shield"
OBSERVER_MNEMONIC_3="rail hole desk six scheme verify gown material belt idle inform lesson great airport ahead quote vehicle brown juice close ladder sun guide recycle"
```
- add these mnemoncic to your local keyring for observers to work

## To deploy contracts
```
npm run deploy
```
## To trigger contract events
```
npm run trigger
```

## To Start observer
- observer has 3 diff. config (config1/2/3) for 3 observer mnemonics
```
cd observer
make build 
./build/observer ./config1.toml
```

## Things I could improve in design
- better verification of heights in msg server for posted data
- currently events are finalised per block, make it customizable such that events could be finalised over a n block basis (n being a params)
- support for different contract events that could be reported
- bonding mechanism for whitelisted operators, whitelisted operators have to deposit some tokens and in case of false events being reported, the tokens could be slashed

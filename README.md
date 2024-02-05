# open-indexer

Open source indexer for avascriptions 

build: go build -o ./indexer ./cmd/main.go  
run: ./indexer

## Data source
You have to provide the data source yourself, this app reads the chained data from mongodb by default.  
We provide an open source program for grabbing data on the chain that can be referenced for you: https://github.com/avascriptions/fetch-chaindata

## Snapshot
Currently the indexed data is not persisted, you must rely on a snapshot to start each time you start, if you do not specify a snapshot, the indexer will start indexing from the initial data.  
Specify the snapshot file command:   
./indexer --snapshot snapshots/snap-xxx.bin  

Of course if you can't run from the genesis block, you can download snapshots from us and the following snapshots are updated regularly.  
01-28 [snap-40953600.bin.zip](https://snapshots.avascriptions.com/snap-40953600.bin.zip) MD5: 73dfea130fe840b81289ac4169f68678  
02-04 [snap-41256000.bin.zip](https://snapshots.avascriptions.com/snap-41256000.bin.zip) MD5: 9e0c2468d0fe44566eafc209bc22f998  

## RPC Interfaces
The indexer implements simple RPC interfaces, the list of interfaces is as follows  
GET /v1/tokens/  
GET /v1/token/:tick  
GET /v1/token/:tick/holders  
GET /v1/address/:addr  
GET /v1/address/:addr/:tick  

 **SEEDS FOR MONGODB**
 mongo proofdex --eval "db.dropDatabase()"
 mongoimport --host localhost --port 27017 --db proofdex --type json --file utils/seed-data/markets.json

**command for mock mockery**
 -name MarketService -output utils/testutils/mocks/ -dir interfaces/ -case underscore

**MONGODB**
 mongod --fork --dbpath ./data/db
 mongod --shutdown
 mongo --host localhost:27017
 
 show dbs
 use proofdex
 show collections
 db.markets.find().limit(10).pretty()
 db.markets.find({"_id": ObjectId("someid")});
 db.markets.find({"_id": ObjectId("someid")}, {field1: 1, field2: 1});
 db.auth("username", "password");
 db.markets.count();
 db.createCollection("collectionName");
 db.markets.insert({field1: "value", field2: "value"})
 db.markets.insert([{field1: "value1"}, {field1: "value2"}])
 db.markets.insertMany([{field1: "value1"}, {field1: "value2"}])
 db.markets.save({"_id": new ObjectId("jhgsdjhgdsf"), field1: "value", field2: "value"});

db.markets.stats()
db.printCollectionStats()
db.markets.dataSize() // Size of the collection
db.markets.storageSize() // Total size of document stored in the collection
db.markets.totalSize() // Total size in bytes for both collection data and indexes
db.markets.totalIndexSize() // Total size of all indexes in the collection
db.markets.latencyStats()

db.markets.drop()

 **APP**
 go run server.go serve

//conn = new Mongo();
//db = conn.getDB("auctions");

/* global use, db */
// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.



//const database = 'admin';
const database = 'auctions';

// Create a new database.
use(database);
db.dropDatabase();

use(database);
//db.getCollectionNames()


// Create collections
db.createCollection('users', { capped: false });
db.createCollection('auctions', { capped: false });
db.createCollection('bids', { capped: false });



// insert collection data
db.users.insertOne({ "_id": "00000001-0000-0000-0000-000000000000", "name": "user1" });
db.users.insertOne({ "_id": "00000002-0000-0000-0000-000000000000", "name": "user2" });
db.users.insertOne({ "_id": "00000003-0000-0000-0000-000000000000", "name": "user3" });

db.getCollection('users').find()
//db.getCollection('auctions').find()
//db.getCollection(collection).find()
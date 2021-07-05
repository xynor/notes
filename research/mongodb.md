## test
````
db.order.insert({"symbol" : "btc/usdt", "price" : 2000, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:33.341Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 1200, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:34.342Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 1000, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:36.343Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 1200, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:38.344Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 1111, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:33.444Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 3333, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:40.544Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 100, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:01.644Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 3123, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:00.644Z")});
db.order.insert({"symbol" : "btc/usdt", "price" : 2000, "qty" : 200, "taker" : 111, "maker" : 222, "time" : ISODate("2021-07-01T12:52:50.644Z")});

````

````
db.order.aggregate([
    { $sort: { "time" : 1 } },
  { 
    "$group": {
      "_id": {
        "$toDate": {
          "$subtract": [{ "$toLong": "$time" },{ "$mod": [ { "$toLong": "$time" }, 1000 * 60 * 1 ] }]
            }
            },
    "open":{"$first":"$$ROOT"},
    "close":{"$last":"$$ROOT"}
    } 
  }
  ])
````
````
db.order.aggregate(
  [
  {"$match": {"time": { "$gte":ISODate("2021-07-01T12:52:00Z"),"$lt": ISODate("2021-07-01T12:53:00Z")}}},
  {"$sort":{"time":1}},
  { "$group": 
    { "_id": ISODate("2021-07-01T12:52:00Z"), 
    "open":{"$first":"$$ROOT"},
    "close":{"$last":"$$ROOT"}
    } 
  }
]
)

db.order.aggregate(
  [
  {"$match": {"time": { "$gte":ISODate("2021-07-01T12:52:00Z"),"$lt": ISODate("2021-07-01T12:53:00Z")}}},
  {"$sort":{"price":1}},
  { "$group": 
    { "_id": ISODate("2021-07-01T12:52:00Z"), 
    "min":{"$first":"$$ROOT"},
    "max":{"$last":"$$ROOT"}
    } 
  }
]
)
````
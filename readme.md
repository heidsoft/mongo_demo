# mongo 

## mongo mac install
- https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/
- brew services start mongodb-community@6.0
- brew services stop mongodb-community@6.0
- mongod --config /usr/local/etc/mongod.conf --fork
- mongod --config /opt/homebrew/etc/mongod.conf --fork
- brew services list

## mongo go driver 
- go get go.mongodb.org/mongo-driver/mongo
package Constants

import "time"

const ErrorFrom = "Api.Calisma"
const DBName = "ordersDB"
const CollectionName = "order"
const MongoConnectionDuration = time.Second*60
const MongoConnectionString = "mongodb://localhost:27017"
const JWTSecretKey = "A.Very.Secret.Key"
const JWTExpTime = time.Minute*5
const CustomerServiceUri = "http://localhost:8000"
const ApplicationName = "OrderService"
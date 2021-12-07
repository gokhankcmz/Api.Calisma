package Constants

import "time"

const ErrorFrom = "Api.Calisma"
const DBName = "customersDB"
const CollectionName = "customer"
const MongoConnectionDuration = time.Second * 60
const MongoConnectionString = "mongodb://localhost:27017"
const JWTSecretKey = "A.Very.Secret.Key"
const JWTExpTime = time.Hour * 5
const ApplicationName = "CustomerService-Api"

const PaginationPerPageDefault = 1
const PaginationPerPageMax = 3
const PaginationOffsetKey = "offset"
const PaginationPageKey = "page"
const PaginationShowAllKey = "showall"
const AcceptableSearchTexts = "name,email"

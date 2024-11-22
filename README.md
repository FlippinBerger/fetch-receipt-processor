# Run Unit tests
from the root directory, run `go test ./...`

# Running the webserver
from the root directory run `go build && ./fetch-receipt-processor`
it will be accessible at `http://localhost:1323`

# Choices I made
I utilized [Echo](https://echo.labstack.com/) as my web framework here. I've
found it to be pretty lightweight, easy to setup, easy to use, and it has some
nice utilities built in so I reach for it to avoid reinventing the wheel for
most things web server setup related (routing, parsing params/request bodies, etc)

I used [Google's uuid pkg](https://github.com/google/uuid) to generate the ids
for my in memory store (simply implemented as a global `map[string]Receipt`). This
is a well trusted and tested package, and again pretty lightweight to remove the
need to come up with a potentially more fragile uuid generation scheme.

My in memory storage solution ended up being a simple global singleton accessed
in the store package. Implemented as a `map[string]receipt` where the key is the uuid.
This is never directly accessed, but rather snagged via a getter function to ensure
it's always initialized once. This would need to be atomically protected in future
versions, but this works fine for the task at hand.

I tried to break things into sensible pkgs based on my own preferences. All the
point calculating or boolean checking methods are collected into a utils package
since I felt like they could potentially be used for more than just this Receipt
type. The method that calls them and spits out the points for the receipt is a
method on the Receipt model itself. This felt like a logical place to house this
method to me, even though some people are against putting logic into their model
pkg.

# Package breakdown

## main
houses the server creation and routing setup logic

## api
houses the controllers called by the routers and handles request/response logic

## model 
houses the data model and the Receipt business logic that calls the helpers and 
calculates the reward points for the Receipt

## store 
houses the singleton storage solution, its init and getter, and the methods used
to get/store Receipts in the underlying map

## utils
houses the methods that parse different strings to generate points or determine 
a condition that, in turn, determines if points should be awarded

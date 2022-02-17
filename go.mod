module everstake/terra-grpc-tx

go 1.17

replace (
	// Use the cosmos-flavored keyring library
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.0
	// latest grpc doesn't work with with our modified proto compiler, so we need to enforce
	// the following version across all dependencies.
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/oasisprotocol/oasis-core/go => github.com/everstake/oasis-core/go v0.2103.9
	google.golang.org/grpc => google.golang.org/grpc v1.33.2

)

require (
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/gogo/protobuf v1.3.3
	github.com/oasisprotocol/oasis-core/go v0.2103.9
	google.golang.org/genproto v0.0.0-20220211171837-173942840c17
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

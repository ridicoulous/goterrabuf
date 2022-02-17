package main

import (
	"context"
	"fmt"

	stakingAPI "github.com/oasisprotocol/oasis-core/go/staking/api"
	"google.golang.org/grpc"
)

func main() {
	// grpcConn, err := grpc.Dial(
	// 	"grpc-terra.everstake.one:9090", // your gRPC server address.
	// 	grpc.WithInsecure(),
	// 	grpc.WithBlock(), // The SDK doesn't support any transport security mechanism.
	// )
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to get grpc client: %s", err))
	// }
	ctx := context.Background()
	// terraClient := tx.NewServiceClient(grpcConn)
	// transaaction, errtx := terraClient.GetTx(ctx, &tx.GetTxRequest{Hash: "F065E474D2B3D606528ED40C55F244503C88834872619BE6AB671D07E89DFFD2"})
	// if errtx != nil {
	// 	fmt.Println(errtx)
	// }
	conn, errc := grpc.Dial(
		"a8.everstake.one:9101", // your gRPC server address.
		grpc.WithInsecure(),
		grpc.WithBlock(), // The SDK doesn't support any transport security mechanism.
	)
	if errc != nil {
		panic(fmt.Sprintf("failed to get grpc client: %s", errc))
	}
	stcl := stakingAPI.NewStakingClient(conn)

	newEpochGenesis, errG := stcl.StateToGenesis(ctx, 7188323)
	if errG != nil {
		fmt.Println(errG)

	}
	fmt.Println(newEpochGenesis)
	// for _, v := range transaaction.TxResponse.Events {
	// 	fmt.Println(v.String())
	// 	for _, event := range v.Attributes {
	// 		fmt.Println(event.String())
	// 	}
	// }
	// fmt.Println(transaaction.TxResponse.Height)
	fmt.Println("Hello, ")
}

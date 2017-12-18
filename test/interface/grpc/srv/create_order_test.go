package srv_test

import (
	"context"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/stretchr/testify/require"
)

func (s *MarketplaceTestSuite) createBidOrder() {
	// arrange
	order := &pb.Order{
		Id:        "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType: pb.OrderType_BID,
		Price:     "777",
		ByuerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",

		Slot: &pb.Slot{
			BuyerRating:    555,
			SupplierRating: 666,
			Resources: &pb.Resources{
				CpuCores: 1,
				GpuCount: pb.GPUCount_SINGLE_GPU,
				RamBytes: 100000000,
				Storage:  1000000000,

				NetworkType:  pb.NetworkType_INCOMING,
				NetTrafficIn: 500000,

				Properties: map[string]float64{
					"hash_rate": 105.7,
				},
			},
		},
	}

	// act
	obtained, err := s.client.CreateOrder(context.Background(), order)

	// assert
	require.NoError(s.T(), err, "cannot create order")
	s.Equal(order.Id, obtained.Id)
}

func (s *MarketplaceTestSuite) createAskOrder() {
	// arrange
	order := &pb.Order{
		Id:         "fc018acd-d9a9-4b8a-a45f-f90456a469c1",
		OrderType:  pb.OrderType_ASK,
		Price:      "777",
		SupplierID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",

		Slot: &pb.Slot{
			SupplierRating: 555,
			Resources: &pb.Resources{
				CpuCores: 1,
				RamBytes: 100000000,
				Storage:  1000000000,

				NetworkType:  pb.NetworkType_INCOMING,
				NetTrafficIn: 500000,

				Properties: map[string]float64{
					"cycles": 42,
					"foo":    1101,
				},
			},
		},
	}

	// act
	obtained, err := s.client.CreateOrder(context.Background(), order)

	// assert
	require.NoError(s.T(), err, "cannot create order")
	s.Equal(order.Id, obtained.Id)
}

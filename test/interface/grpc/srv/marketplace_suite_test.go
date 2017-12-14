package srv_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	// registers grpc gzip encoder/decoder
	_ "google.golang.org/grpc/encoding/gzip"

	"github.com/stretchr/testify/suite"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/test/cli"
)

// TestMarketplaceTestSuite initialize test suit.
func TestMarketplaceTestSuite(t *testing.T) {
	suite.Run(t, new(MarketplaceTestSuite))
}

// MarketplaceTestSuite tests Marketplace service.
type MarketplaceTestSuite struct {
	cli.AppTestSuite

	conn   *grpc.ClientConn
	client pb.MarketClient
}

// SetupTest prepare state for the test.
func (s *MarketplaceTestSuite) SetupTest() {
	// call parent's method.
	s.AppTestSuite.SetupTest()

	conn, err := gRpcClient()
	s.Require().NoError(err, "cannot get grpc client")

	s.conn = conn
	s.client = pb.NewMarketClient(conn)
}

// TearDownTest
func (s *MarketplaceTestSuite) TearDownTest() {
	s.conn.Close()

	s.AppTestSuite.TearDownTest()
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *MarketplaceTestSuite) TestMarketPlace() {

	//s.T().Run("CreateOrder", func(t *testing.T) {
	s.createOrder()
	//})

	//	s.T().Run("GetOrderByID", func(t *testing.T) {
	s.getOrderByID()
	//	})
}

func gRpcClient() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *cli.ListenAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBackoffConfig(grpc.BackoffConfig{MaxDelay: 2 * time.Second}),
	)

	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err.Error())
	}
	return conn, err
}

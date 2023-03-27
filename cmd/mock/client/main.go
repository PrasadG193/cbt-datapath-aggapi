package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/PrasadG193/cbt-datapath/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const URL = "cbt-datapath.cbt-datapath:9000"

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	cert := os.Getenv("CA_CERT")
	pemServerCA, err := ioutil.ReadFile(cert)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	URL := os.Getenv("CBT_URL")
	phClient := NewListVolumeSnapshotDeltas(URL)
	resp, err := phClient.ListVolumeSnapshotDeltas(ctx)
	fmt.Println(resp, err)
}

type Client struct {
	client pb.VolumeSnapshotDeltaServiceClient
}

func NewListVolumeSnapshotDeltas(URL string) Client {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewVolumeSnapshotDeltaServiceClient(conn)
	return Client{client: client}
}

func (c *Client) ListVolumeSnapshotDeltas(ctx context.Context) (*pb.VolumeSnapshotDeltaResponse, error) {
	token := os.Getenv("CBT_TOKEN")

	resp, err := c.client.ListVolumeSnapshotDeltas(ctx, &pb.VolumeSnapshotDeltaRequest{
		SnapshotBase:   stringToPtr("vs-01"),
		SnapshotTarget: "vs-02",
		Token:          token,
	})
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}

func stringToPtr(s string) *string {
	return &s
}

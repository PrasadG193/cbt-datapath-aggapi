package server

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	pb "github.com/PrasadG193/cbt-datapath-aggapi/pkg/grpc"
	"github.com/PrasadG193/cbt-datapath-aggapi/pkg/storage"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var blockDeltas []*pb.ChangedBlockDelta

func init() {
	noOfBlocks, err := strconv.Atoi(os.Getenv(NumberOfBlocksKey))
	if err != nil {
		return
	}
	if noOfBlocks == 0 {
		noOfBlocks = 5
	}
	for i := 0; i < noOfBlocks; i++ {
		blockDeltas = append(blockDeltas, &pb.ChangedBlockDelta{
			Offset:         uint64(i),
			BlockSizeBytes: 524288,
			DataToken: &pb.DataToken{
				Token:        "ieEEQ9Bj7E6XR",
				IssuanceTime: timestamppb.Now(),
				TtlSeconds:   durationpb.New(time.Minute * 180),
			},
		})
	}
}

const (
	NumberOfBlocksKey = "NUMBER_OF_BLOCKS"
)

type Server struct {
	pb.UnimplementedVolumeSnapshotDeltaServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) ListVolumeSnapshotDeltas(
	ctx context.Context,
	req *pb.VolumeSnapshotDeltaRequest,
) (*pb.VolumeSnapshotDeltaResponse, error) {
	var (
		nextToken       = "uXonK48vfznJS"
		volumeSizeBytes = uint64(1073741824)
	)

	if !storage.ValidToken(req.Token) {
		return nil, errors.New("Invalid token")
	}

	return &pb.VolumeSnapshotDeltaResponse{
		BlockDelta: &pb.BlockVolumeSnapshotDelta{
			ChangedBlockDeltas: blockDeltas,
		},
		VolumeSizeBytes: &volumeSizeBytes,
		NextToken:       &nextToken,
	}, nil
}

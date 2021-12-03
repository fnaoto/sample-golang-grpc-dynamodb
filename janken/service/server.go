package service

import (
	"context"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"janken/db"
	"janken/pb"
	"janken/pkg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	db.Init()
}

type JankenService struct {
	tryGames      int32
	wins          int32
	jankenResults []*pb.JankenResult
}

func NewJankenService() *JankenService {
	return &JankenService{
		tryGames:      0,
		wins:          0,
		jankenResults: make([]*pb.JankenResult, 0),
	}
}

func (s *JankenService) PlayJanken(ctx context.Context, req *pb.PlayJankenRequest) (*pb.PlayJankenResponse, error) {
	if req.Hands == pb.Hands_UNKNOWN_HANDS {
		return nil, status.Errorf(codes.InvalidArgument, "Choose GU, PA, or CHOKI.")
	}

	computerHand := pkg.EncodeHands(int32(rand.Intn(3) + 1))

	var result pb.Result
	if req.Hands == computerHand {
		result = pb.Result_DRAW
	} else if (req.Hands.Number()-computerHand.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}

	now := time.Now()

	jankenResult := &pb.JankenResult{
		Id:           int32(now.Nanosecond()),
		MyHand:       req.Hands,
		ComputerHand: computerHand,
		Result:       result,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	s.tryGames = s.tryGames + 1
	if result == pb.Result_WIN {
		s.wins = s.wins + 1
	}

	db.PutItem(jankenResult)

	return &pb.PlayJankenResponse{
		JankenResult: jankenResult,
	}, nil
}

func (s *JankenService) PlayJankenResults(ctx context.Context, req *pb.PlayResultRequest) (*pb.PlayResultResponse, error) {
	jankenResults := db.Scan()
	return &pb.PlayResultResponse{
		Report: &pb.Report{
			TryGames:      s.tryGames,
			Wins:          s.wins,
			JankenResults: jankenResults,
		},
	}, nil
}

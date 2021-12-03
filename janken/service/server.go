package service

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"janken/db"
	pb "janken/pb"
	pkg "janken/pkg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	dynamodbService()
}

func dynamodbService() {
	dynamodbEndpoint := "http://dynamodb:8000"
	dynamodbRegion := "ap-northeast-1"
	dynamodbTable := "table"

	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(dynamodbEndpoint),
		Region:   aws.String(dynamodbRegion),
	})

	if err != nil {
		log.Fatal(err)
	}

	ddbr := &db.DynamoDBConfig{
		Client: dynamodb.New(awsSession, aws.NewConfig()),
		Table:  dynamodbTable,
	}

	ddbr.CreateTable()
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
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	yourHand := pkg.EncodeHands(int32(rand.Intn(3) + 1))

	var result pb.Result
	if req.Hands == yourHand {
		result = pb.Result_DRAW
	} else if (req.Hands.Number()-yourHand.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}

	now := time.Now()

	jankenResult := &pb.JankenResult{
		MyHand:   req.Hands,
		YourHand: yourHand,
		Result:   result,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	s.tryGames = s.tryGames + 1
	if result == pb.Result_WIN {
		s.wins = s.wins + 1
	}
	s.jankenResults = append(s.jankenResults, jankenResult)

	return &pb.PlayJankenResponse{
		JankenResult: jankenResult,
	}, nil
}

func (s *JankenService) PlayJankenResults(ctx context.Context, req *pb.PlayResultRequest) (*pb.PlayResultResponse, error) {

	return &pb.PlayResultResponse{
		Report: &pb.Report{
			TryGames:      s.tryGames,
			Wins:          s.wins,
			JankenResults: s.jankenResults,
		},
	}, nil
}

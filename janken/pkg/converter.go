package pkg

import "janken/pb"

func EncodeHands(n int32) pb.Hands {
	switch n {
	case 1:
		return pb.Hands_GU
	case 2:
		return pb.Hands_PA
	case 3:
		return pb.Hands_CHOKI
	default:
		return pb.Hands_UNKNOWN_HANDS
	}
}


package global

import (
	pb "github.com/moke-game/platform.git/api/gen/analytics"
)

var (
	analyticsClient pb.AnalyticsServiceClient
)

func SetAnalyticsClient(cli pb.AnalyticsServiceClient) {
	analyticsClient = cli
}

func GetAnalyticsSender() pb.AnalyticsServiceClient {
	return analyticsClient
}

package global

import (
	pb "github.com/gstones/platform/api/gen/analytics"
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

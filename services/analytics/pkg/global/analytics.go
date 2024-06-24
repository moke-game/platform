package global

import (
	pb "github.com/moke-game/platform/api/gen/analytics"
)

var (
	analyticsClient pb.AnalyticsServiceClient
)

// SetAnalyticsClient sets the global analytics client
func SetAnalyticsClient(cli pb.AnalyticsServiceClient) {
	analyticsClient = cli
}

// GetAnalyticsSender returns the global analytics client
func GetAnalyticsSender() pb.AnalyticsServiceClient {
	if analyticsClient == nil {
		return &Noop{}
	}
	return analyticsClient
}

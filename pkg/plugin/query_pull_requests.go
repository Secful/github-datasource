package plugin

import (
	"context"

	"github.com/grafana/grafana-github-datasource/pkg/dfutil"
	"github.com/grafana/grafana-github-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func (s *Server) handlePullRequestsQuery(ctx context.Context, q backend.DataQuery) backend.DataResponse {
	query := &models.PullRequestsQuery{}
	if err := UnmarshalQuery(q.JSON, query); err != nil {
		return *err
	}
	return dfutil.FrameResponseWithError(s.Datasource.HandlePullRequestsQuery(ctx, query, q))
}

func (s *Server) HandlePullRequests(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	return &backend.QueryDataResponse{
		Responses: processQueries(ctx, req, s.handlePullRequestsQuery),
	}, nil
}
package bybit

import (
	"context"
	"github.com/google/go-querystring/query"
)

type V5AnnouncementsServiceI interface {
	GetAnnouncement(context.Context, V5GetAnnouncementParam) (*V5GetAnnouncementResponse, error)
}

type V5AnnouncementsService struct {
	client *Client
}

type V5GetAnnouncementResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5Announcement `json:"result"`
}

type V5Announcement struct {
	Total int                  `json:"total"`
	List  []V5AnnouncementItem `json:"list"`
}

type V5AnnouncementItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        struct {
		Title string `json:"title"`
		Key   string `json:"key"`
	} `json:"type"`
	Tags               []string `json:"tags"`
	Url                string   `json:"url"`
	DateTimestamp      int64    `json:"dateTimestamp"`
	StartDateTimestamp int64    `json:"startDateTimestamp"`
	EndDateTimestamp   int64    `json:"endDateTimestamp"`
}

type V5GetAnnouncementParam struct {
	Locale string  `url:"locale"` // required
	Type   *string `url:"type,omitempty"`
	Tag    *string `url:"tag,omitempty"`
	Page   *int    `url:"page,omitempty"`  // default: 1
	Limit  *int    `url:"limit,omitempty"` // default: 20
}

func (s *V5AnnouncementsService) GetAnnouncement(
	ctx context.Context,
	param V5GetAnnouncementParam,
) (res *V5GetAnnouncementResponse, err error) {

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	err = s.client.getV5PrivatelyCtx(ctx, "/v5/announcements/index", queryString, &res)

	return
}

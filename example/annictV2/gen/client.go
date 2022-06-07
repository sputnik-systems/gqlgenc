// Code generated by github.com/sputnik-systems/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"
	"net/http"
	"time"

	"github.com/sputnik-systems/gqlgenc/clientv2"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, interceptors...)}
}

type Query struct {
	Node                Node                    "json:\"node\" graphql:\"node\""
	Nodes               []Node                  "json:\"nodes\" graphql:\"nodes\""
	SearchCharacters    *CharacterConnection    "json:\"searchCharacters\" graphql:\"searchCharacters\""
	SearchEpisodes      *EpisodeConnection      "json:\"searchEpisodes\" graphql:\"searchEpisodes\""
	SearchOrganizations *OrganizationConnection "json:\"searchOrganizations\" graphql:\"searchOrganizations\""
	SearchPeople        *PersonConnection       "json:\"searchPeople\" graphql:\"searchPeople\""
	SearchWorks         *WorkConnection         "json:\"searchWorks\" graphql:\"searchWorks\""
	User                *User                   "json:\"user\" graphql:\"user\""
	Viewer              *User                   "json:\"viewer\" graphql:\"viewer\""
}

type Mutation struct {
	CreateRecord *CreateRecordPayload "json:\"createRecord\" graphql:\"createRecord\""
	CreateReview *CreateReviewPayload "json:\"createReview\" graphql:\"createReview\""
	DeleteRecord *DeleteRecordPayload "json:\"deleteRecord\" graphql:\"deleteRecord\""
	DeleteReview *DeleteReviewPayload "json:\"deleteReview\" graphql:\"deleteReview\""
	UpdateRecord *UpdateRecordPayload "json:\"updateRecord\" graphql:\"updateRecord\""
	UpdateReview *UpdateReviewPayload "json:\"updateReview\" graphql:\"updateReview\""
	UpdateStatus *UpdateStatusPayload "json:\"updateStatus\" graphql:\"updateStatus\""
}

type ViewerFragment struct {
	AvatarURL       *string "json:\"avatar_url\" graphql:\"avatar_url\""
	RecordsCount    int64   "json:\"recordsCount\" graphql:\"recordsCount\""
	WannaWatchCount int64   "json:\"wannaWatchCount\" graphql:\"wannaWatchCount\""
	WatchingCount   int64   "json:\"watchingCount\" graphql:\"watchingCount\""
	WatchedCount    int64   "json:\"watchedCount\" graphql:\"watchedCount\""
}

type WorkFragment struct {
	ID                string       "json:\"id\" graphql:\"id\""
	Title             string       "json:\"title\" graphql:\"title\""
	AnnictID          int64        "json:\"annict_id\" graphql:\"annict_id\""
	SeasonYear        *int64       "json:\"seasonYear\" graphql:\"seasonYear\""
	SeasonName        *SeasonName  "json:\"seasonName\" graphql:\"seasonName\""
	EpisodesCount     int64        "json:\"episodesCount\" graphql:\"episodesCount\""
	OfficialSiteURL   *string      "json:\"officialSiteUrl\" graphql:\"officialSiteUrl\""
	WikipediaURL      *string      "json:\"wikipediaUrl\" graphql:\"wikipediaUrl\""
	ViewerStatusState *StatusState "json:\"viewerStatusState\" graphql:\"viewerStatusState\""
	Episodes          *struct {
		Nodes []*struct {
			ID         string  "json:\"id\" graphql:\"id\""
			AnnictID   int64   "json:\"annictId\" graphql:\"annictId\""
			Title      *string "json:\"title\" graphql:\"title\""
			SortNumber int64   "json:\"sortNumber\" graphql:\"sortNumber\""
		} "json:\"nodes\" graphql:\"nodes\""
	} "json:\"episodes\" graphql:\"episodes\""
}

type HogeCreateRecordMutationPayload struct {
	CreateRecord *struct {
		ClientMutationID *string "json:\"clientMutationId\" graphql:\"clientMutationId\""
	} "json:\"createRecord\" graphql:\"createRecord\""
}

type HogeUpdateStatusMutationPayload struct {
	UpdateStatus *struct {
		ClientMutationID *string "json:\"clientMutationId\" graphql:\"clientMutationId\""
	} "json:\"updateStatus\" graphql:\"updateStatus\""
}

type HogeUpdateWorkStatusPayload struct {
	UpdateStatus *struct {
		ClientMutationID *string "json:\"clientMutationId\" graphql:\"clientMutationId\""
	} "json:\"updateStatus\" graphql:\"updateStatus\""
}

type GetProfile struct {
	Viewer *ViewerFragment "json:\"viewer\" graphql:\"viewer\""
}

type ListWorks struct {
	Viewer *struct {
		Works *struct {
			Edges []*struct {
				Cursor string        "json:\"cursor\" graphql:\"cursor\""
				Node   *WorkFragment "json:\"node\" graphql:\"node\""
			} "json:\"edges\" graphql:\"edges\""
		} "json:\"works\" graphql:\"works\""
	} "json:\"viewer\" graphql:\"viewer\""
}

type ListRecords struct {
	Viewer *struct {
		Records *struct {
			Edges []*struct {
				Node *struct {
					Work    WorkFragment "json:\"work\" graphql:\"work\""
					Episode struct {
						SortNumber int64 "json:\"sortNumber\" graphql:\"sortNumber\""
					} "json:\"episode\" graphql:\"episode\""
					CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
				} "json:\"node\" graphql:\"node\""
			} "json:\"edges\" graphql:\"edges\""
		} "json:\"records\" graphql:\"records\""
	} "json:\"viewer\" graphql:\"viewer\""
}

type ListNextEpisodes struct {
	Viewer *struct {
		Records *struct {
			Edges []*struct {
				Node *struct {
					Episode struct {
						NextEpisode *struct {
							ID          string  "json:\"id\" graphql:\"id\""
							Number      *int64  "json:\"number\" graphql:\"number\""
							NumberText  *string "json:\"numberText\" graphql:\"numberText\""
							Title       *string "json:\"title\" graphql:\"title\""
							NextEpisode *struct {
								ID string "json:\"id\" graphql:\"id\""
							} "json:\"nextEpisode\" graphql:\"nextEpisode\""
						} "json:\"nextEpisode\" graphql:\"nextEpisode\""
						Work struct {
							ID    string "json:\"id\" graphql:\"id\""
							Title string "json:\"title\" graphql:\"title\""
						} "json:\"work\" graphql:\"work\""
					} "json:\"episode\" graphql:\"episode\""
				} "json:\"node\" graphql:\"node\""
			} "json:\"edges\" graphql:\"edges\""
		} "json:\"records\" graphql:\"records\""
	} "json:\"viewer\" graphql:\"viewer\""
}

type GetWork struct {
	SearchWorks *struct {
		Nodes []*WorkFragment "json:\"nodes\" graphql:\"nodes\""
	} "json:\"searchWorks\" graphql:\"searchWorks\""
}

type SearchWorks struct {
	SearchWorks *struct {
		Nodes []*struct {
			ID                string       "json:\"id\" graphql:\"id\""
			Title             string       "json:\"title\" graphql:\"title\""
			AnnictID          int64        "json:\"annict_id\" graphql:\"annict_id\""
			SeasonYear        *int64       "json:\"seasonYear\" graphql:\"seasonYear\""
			SeasonName        *SeasonName  "json:\"seasonName\" graphql:\"seasonName\""
			EpisodesCount     int64        "json:\"episodesCount\" graphql:\"episodesCount\""
			OfficialSiteURL   *string      "json:\"officialSiteUrl\" graphql:\"officialSiteUrl\""
			WikipediaURL      *string      "json:\"wikipediaUrl\" graphql:\"wikipediaUrl\""
			ViewerStatusState *StatusState "json:\"viewerStatusState\" graphql:\"viewerStatusState\""
			Episodes          *struct {
				Nodes []*struct {
					ID         string  "json:\"id\" graphql:\"id\""
					AnnictID   int64   "json:\"annictId\" graphql:\"annictId\""
					Title      *string "json:\"title\" graphql:\"title\""
					SortNumber int64   "json:\"sortNumber\" graphql:\"sortNumber\""
				} "json:\"nodes\" graphql:\"nodes\""
			} "json:\"episodes\" graphql:\"episodes\""
			Work struct {
				Image *struct {
					RecommendedImageURL *string "json:\"recommendedImageUrl\" graphql:\"recommendedImageUrl\""
				} "json:\"image\" graphql:\"image\""
			} "graphql:\"... on Work\""
		} "json:\"nodes\" graphql:\"nodes\""
	} "json:\"searchWorks\" graphql:\"searchWorks\""
}

const CreateRecordMutationQuery = `mutation CreateRecordMutation ($episodeId: ID!) {
	createRecord(input: {episodeId:$episodeId}) {
		clientMutationId
	}
}
`

func (c *Client) CreateRecordMutation(ctx context.Context, episodeID string, interceptors ...clientv2.RequestInterceptor) (*HogeCreateRecordMutationPayload, error) {
	vars := map[string]interface{}{
		"episodeId": episodeID,
	}

	var res HogeCreateRecordMutationPayload
	if err := c.Client.Post(ctx, "CreateRecordMutation", CreateRecordMutationQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const UpdateStatusMutationQuery = `mutation UpdateStatusMutation ($state: StatusState!, $workId: ID!) {
	updateStatus(input: {state:$state,workId:$workId}) {
		clientMutationId
	}
}
`

func (c *Client) UpdateStatusMutation(ctx context.Context, state StatusState, workID string, interceptors ...clientv2.RequestInterceptor) (*HogeUpdateStatusMutationPayload, error) {
	vars := map[string]interface{}{
		"state":  state,
		"workId": workID,
	}

	var res HogeUpdateStatusMutationPayload
	if err := c.Client.Post(ctx, "UpdateStatusMutation", UpdateStatusMutationQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const UpdateWorkStatusQuery = `mutation UpdateWorkStatus ($workId: ID!) {
	updateStatus(input: {state:WATCHING,workId:$workId}) {
		clientMutationId
	}
}
`

func (c *Client) UpdateWorkStatus(ctx context.Context, workID string, interceptors ...clientv2.RequestInterceptor) (*HogeUpdateWorkStatusPayload, error) {
	vars := map[string]interface{}{
		"workId": workID,
	}

	var res HogeUpdateWorkStatusPayload
	if err := c.Client.Post(ctx, "UpdateWorkStatus", UpdateWorkStatusQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetProfileQuery = `query GetProfile {
	viewer {
		... ViewerFragment
	}
}
fragment ViewerFragment on User {
	avatar_url: avatarUrl
	recordsCount
	wannaWatchCount
	watchingCount
	watchedCount
}
`

func (c *Client) GetProfile(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*GetProfile, error) {
	vars := map[string]interface{}{}

	var res GetProfile
	if err := c.Client.Post(ctx, "GetProfile", GetProfileQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListWorksQuery = `query ListWorks ($state: StatusState, $after: String, $n: Int!) {
	viewer {
		works(state: $state, after: $after, first: $n, orderBy: {direction:DESC,field:SEASON}) {
			edges {
				cursor
				node {
					... WorkFragment
				}
			}
		}
	}
}
fragment WorkFragment on Work {
	id
	title
	annict_id: annictId
	seasonYear
	seasonName
	episodesCount
	officialSiteUrl
	wikipediaUrl
	viewerStatusState
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			id
			annictId
			title
			sortNumber
		}
	}
}
`

func (c *Client) ListWorks(ctx context.Context, state *StatusState, after *string, n int64, interceptors ...clientv2.RequestInterceptor) (*ListWorks, error) {
	vars := map[string]interface{}{
		"state": state,
		"after": after,
		"n":     n,
	}

	var res ListWorks
	if err := c.Client.Post(ctx, "ListWorks", ListWorksQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListRecordsQuery = `query ListRecords {
	viewer {
		records {
			edges {
				node {
					work {
						... WorkFragment
					}
					episode {
						sortNumber
					}
					createdAt
				}
			}
		}
	}
}
fragment WorkFragment on Work {
	id
	title
	annict_id: annictId
	seasonYear
	seasonName
	episodesCount
	officialSiteUrl
	wikipediaUrl
	viewerStatusState
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			id
			annictId
			title
			sortNumber
		}
	}
}
`

func (c *Client) ListRecords(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*ListRecords, error) {
	vars := map[string]interface{}{}

	var res ListRecords
	if err := c.Client.Post(ctx, "ListRecords", ListRecordsQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListNextEpisodesQuery = `query ListNextEpisodes {
	viewer {
		records {
			edges {
				node {
					episode {
						nextEpisode {
							id
							number
							numberText
							title
							nextEpisode {
								id
							}
						}
						work {
							id
							title
						}
					}
				}
			}
		}
	}
}
`

func (c *Client) ListNextEpisodes(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*ListNextEpisodes, error) {
	vars := map[string]interface{}{}

	var res ListNextEpisodes
	if err := c.Client.Post(ctx, "ListNextEpisodes", ListNextEpisodesQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetWorkQuery = `query GetWork ($ids: [Int!]) {
	searchWorks(annictIds: $ids) {
		nodes {
			... WorkFragment
		}
	}
}
fragment WorkFragment on Work {
	id
	title
	annict_id: annictId
	seasonYear
	seasonName
	episodesCount
	officialSiteUrl
	wikipediaUrl
	viewerStatusState
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			id
			annictId
			title
			sortNumber
		}
	}
}
`

func (c *Client) GetWork(ctx context.Context, ids []int64, interceptors ...clientv2.RequestInterceptor) (*GetWork, error) {
	vars := map[string]interface{}{
		"ids": ids,
	}

	var res GetWork
	if err := c.Client.Post(ctx, "GetWork", GetWorkQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const SearchWorksQuery = `query SearchWorks ($seasons: [String!]) {
	searchWorks(seasons: $seasons, first: 1, orderBy: {field:WATCHERS_COUNT,direction:DESC}) {
		nodes {
			... WorkFragment
			... on Work {
				image {
					recommendedImageUrl
				}
			}
		}
	}
}
fragment WorkFragment on Work {
	id
	title
	annict_id: annictId
	seasonYear
	seasonName
	episodesCount
	officialSiteUrl
	wikipediaUrl
	viewerStatusState
	episodes(orderBy: {direction:ASC,field:SORT_NUMBER}) {
		nodes {
			id
			annictId
			title
			sortNumber
		}
	}
}
`

func (c *Client) SearchWorks(ctx context.Context, seasons []string, interceptors ...clientv2.RequestInterceptor) (*SearchWorks, error) {
	vars := map[string]interface{}{
		"seasons": seasons,
	}

	var res SearchWorks
	if err := c.Client.Post(ctx, "SearchWorks", SearchWorksQuery, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

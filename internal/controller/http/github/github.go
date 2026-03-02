package github

import repositoryv1 "github.com/vibino-xyz/protos/contracts/build/go/repository/v1"

type GitHubWebhookPayload struct {
	Hook struct {
		Type   string   `json:"type"`
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
	} `json:"hook"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	} `json:"repository"`
	Url      string `json:"url"`
	CloneUrl string `json:"clone_url"`
}

func ToProtos(payload GitHubWebhookPayload) *repositoryv1.RepositoryEventMessage {
	return &repositoryv1.RepositoryEventMessage{
		Provider:      repositoryv1.Provider_GITHUB,
		EventType:     repositoryv1.EventType_FULL_INDEX,
		RepositoryId:  int64(payload.Repository.ID),
		RepoFullName:  payload.Repository.FullName,
		DefaultBranch: "main",
		CloneUrl:      payload.CloneUrl,
	}
}

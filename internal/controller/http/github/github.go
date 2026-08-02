package github

import contracts "github.com/vibino-xyz/protos/contracts/build"

type PushPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		Id            int    `json:"id"`
		FullName      string `json:"full_name"`
		Private       bool   `json:"private"`
		CloneURL      string `json:"clone_url"`
		SSHURL        string `json:"ssh_url"`
		DefaultBranch string `json:"default_branch"`
	} `json:"repository"`
	Commits []struct {
		Id       string   `json:"id"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		Id       string   `json:"id"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"head_commit"`
	Installation struct {
		Id string `json:"id"`
	} `json:"installation"`
}

func ToProto(payload PushPayload) *contracts.RepositoryEventMessage {
	return &contracts.RepositoryEventMessage{
		Provider:  contracts.Provider_GITHUB,
		EventType: contracts.EventType_INCREMENTAL_INDEX,
		Repository: &contracts.Repository{
			Id:            int64(payload.Repository.Id),
			FullName:      payload.Repository.FullName,
			Private:       payload.Repository.Private,
			CloneUrl:      payload.Repository.CloneURL,
			SshUrl:        payload.Repository.SSHURL,
			DefaultBranch: payload.Repository.DefaultBranch,
		},
		Commits: make([]*contracts.Commit, 0),
		HeadCommit: &contracts.Commit{
			Id:       payload.HeadCommit.Id,
			Added:    make([]string, 0),
			Removed:  make([]string, 0),
			Modified: make([]string, 0),
		},
	}
}

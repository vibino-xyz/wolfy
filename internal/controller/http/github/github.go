package github

import contracts "github.com/vibino-xyz/protos/contracts/build"

type PushPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		Id            int64  `json:"id"`
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
		Id int64 `json:"id"`
	} `json:"installation"`
}

func ToProto(payload PushPayload) *contracts.RepositoryEventMessage {
	commits := make([]*contracts.Commit, 0, len(payload.Commits))
	for _, c := range payload.Commits {
		commits = append(commits, &contracts.Commit{
			Id:       c.Id,
			Added:    c.Added,
			Removed:  c.Removed,
			Modified: c.Modified,
		})
	}

	return &contracts.RepositoryEventMessage{
		Provider:  contracts.Provider_GITHUB,
		EventType: contracts.EventType_INCREMENTAL_INDEX,
		Repository: &contracts.Repository{
			Id:            payload.Repository.Id,
			FullName:      payload.Repository.FullName,
			Private:       payload.Repository.Private,
			CloneUrl:      payload.Repository.CloneURL,
			SshUrl:        payload.Repository.SSHURL,
			DefaultBranch: payload.Repository.DefaultBranch,
		},
		Commits: commits,
		HeadCommit: &contracts.Commit{
			Id:       payload.HeadCommit.Id,
			Added:    payload.HeadCommit.Added,
			Removed:  payload.HeadCommit.Removed,
			Modified: payload.HeadCommit.Modified,
		},
		InstallationId: payload.Installation.Id,
	}
}

package webhooks

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

type GithubWebhookEvent struct {
}

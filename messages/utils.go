package messages

func formatUsername(username string) string {
	if username == "" {
		return "anonymous"
	}
	if (username)[0] == '@' {
		return username
	}
	return "@" + username
}

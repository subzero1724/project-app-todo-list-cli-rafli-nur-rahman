package utils

func ColorStatus(status string) string {
	switch status {
	case "completed":
		return "[32m" + status + "\033[0m" // Green
	case "pending":
		return "[31m" + status + "\033[0m" // Red
	case "progress":
		return "[33m" + status + "\033[0m" // Yellow
	case "new":
		return "[36m" + status + "\033[0m" // Cyan
	default:
		return status
	}
}

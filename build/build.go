package build

import "strings"

var (
	// Build build time
	Build string
	// GitCommitHash git commit hash
	GitCommitHash string
)

// Version 生成版本信息
func Version(prefix string) string {
	var buf strings.Builder

	if prefix != "" {
		buf.WriteString(prefix)
	}

	if Build != "" {
		buf.WriteByte('\n')
		buf.WriteString("Build: ")
		buf.WriteString(Build)
	}
	if GitCommitHash != "" {
		buf.WriteByte('\n')
		buf.WriteString("Commit: ")
		buf.WriteString(GitCommitHash)
	}
	return buf.String()
}

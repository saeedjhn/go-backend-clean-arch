package article

type ModerationStatus string

const (
	CommentPending  = ModerationStatus("pending")
	CommentApproved = ModerationStatus("approved")
	CommentRejected = ModerationStatus("rejected")
)

var moderationStatusStrings = map[ModerationStatus]string{ //nolint:gochecknoglobals // nothing
	CommentPending:  "pending",
	CommentApproved: "approved",
	CommentRejected: "rejected",
}

func (a ModerationStatus) IsValidStatus() bool {
	_, ok := moderationStatusStrings[a]

	return ok
}

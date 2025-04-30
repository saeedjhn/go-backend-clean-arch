package article

type CommentStatus string

const (
	CommentPending  = CommentStatus("pending")
	CommentApproved = CommentStatus("approved")
	CommentRejected = CommentStatus("rejected")
)

var commentStatusStrings = map[CommentStatus]string{ //nolint:gochecknoglobals // nothing
	CommentPending:  "pending",
	CommentApproved: "approved",
	CommentRejected: "rejected",
}

func (a CommentStatus) IsValidStatus() bool {
	_, ok := commentStatusStrings[a]

	return ok
}

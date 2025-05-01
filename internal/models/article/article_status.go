package article

type PublicationStatus string

const (
	ArticleDraft     = PublicationStatus("draft")
	ArticlePublished = PublicationStatus("published")
	ArticleArchived  = PublicationStatus("archived")
)

var articleStatusStrings = map[PublicationStatus]string{ //nolint:gochecknoglobals // nothing
	ArticleDraft:     "draft",
	ArticlePublished: "published",
	ArticleArchived:  "archived",
}

func (a PublicationStatus) IsValidStatus() bool {
	_, ok := articleStatusStrings[a]

	return ok
}

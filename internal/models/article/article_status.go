package article

type ArticleStatus string

const (
	ArticleDraft     = ArticleStatus("draft")
	ArticlePublished = ArticleStatus("published")
	ArticleArchived  = ArticleStatus("archived")
)

var articleStatusStrings = map[ArticleStatus]string{ //nolint:gochecknoglobals // nothing
	ArticleDraft:     "draft",
	ArticlePublished: "published",
	ArticleArchived:  "archived",
}

func (a ArticleStatus) IsValidStatus() bool {
	_, ok := articleStatusStrings[a]

	return ok
}

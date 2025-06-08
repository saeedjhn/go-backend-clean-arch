package article

import (
	"errors"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

type Article struct {
	ID      types.ID
	Title   string
	Slug    string
	Content string
	Status  PublicationStatus
	// Structure Optimization:
	// Separate Comments from Article:
	// If the volume of comments is very high, manage them in a separate aggregate (using Event Sourcing or CQRS).
	Comments   []Comment
	AuthorID   types.ID
	CategoryID []types.ID
	TagIDs     []types.ID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// IsDeleted  bool
}

func (a *Article) AddComment(comment Comment) {
	a.Comments = append(a.Comments, comment)
}

func (a *Article) RemoveComment(commentID types.ID) {
	for i, c := range a.Comments {
		if c.ID == commentID {
			a.Comments = append(a.Comments[:i], a.Comments[i+1:]...)
			break
		}
	}
}

func (a *Article) Publish() error {
	if a.Status == ArticleArchived {
		return errors.New("cannot publish archived article")
	}
	a.Status = ArticlePublished
	return nil
}

type Comment struct {
	ID        types.ID
	Text      string
	Status    ModerationStatus
	AuthorID  types.ID
	CreatedAt time.Time
}

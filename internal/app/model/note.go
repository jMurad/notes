package model

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jMurad/notes/pkg/yaspeller"
)

// Note ...
type Note struct {
	ID       int       `json:"id"`
	Created  time.Time `json:"created"`
	Content  string    `json:"content"`
	AuthorID int       `json:"author_id"`
}

// Validate ...
func (n *Note) Validate() error {
	errs := url.Values{}

	res, err := yaspeller.CheckText(n.Content)
	switch {
	case n.Content == "":
		errs.Add("content", "the content field is required!")
	case len([]rune(n.Content)) < 20 || len([]rune(n.Content)) > 500:
		errs.Add("content", "the content field must be between 20-500 chars!")
	case err != nil:
		errs.Add("content", "the content could not be verified!")
	case !res.IsCorrect():
		hint := res.RightText(n.Content)
		errs.Add("content", fmt.Sprintf("the content contains spelling errors (hint=%s)", hint))
	}

	if len(errs) != 0 {
		e := ""
		for k, v := range errs {
			e += fmt.Sprintf("%s: %s; ", k, strings.Join(v, ", "))
		}
		err = fmt.Errorf("%s", e)
	}

	return err
}

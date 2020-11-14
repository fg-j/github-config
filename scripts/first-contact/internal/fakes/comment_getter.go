package fakes

import (
	"sync"

	"github.com/paketo-buildpacks/github-config/scripts/first-contact/internal"
)

type CommentGetter struct {
	GetCreatedAtCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			String string
		}
		Stub func() string
	}
	GetFirstReplyCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Comment internal.Comment
			Error   error
		}
		Stub func() (internal.Comment, error)
	}
}

func (f *CommentGetter) GetCreatedAt() string {
	f.GetCreatedAtCall.Lock()
	defer f.GetCreatedAtCall.Unlock()
	f.GetCreatedAtCall.CallCount++
	if f.GetCreatedAtCall.Stub != nil {
		return f.GetCreatedAtCall.Stub()
	}
	return f.GetCreatedAtCall.Returns.String
}
func (f *CommentGetter) GetFirstReply() (internal.Comment, error) {
	f.GetFirstReplyCall.Lock()
	defer f.GetFirstReplyCall.Unlock()
	f.GetFirstReplyCall.CallCount++
	if f.GetFirstReplyCall.Stub != nil {
		return f.GetFirstReplyCall.Stub()
	}
	return f.GetFirstReplyCall.Returns.Comment, f.GetFirstReplyCall.Returns.Error
}

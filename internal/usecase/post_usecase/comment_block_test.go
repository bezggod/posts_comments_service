package post_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_comments_service/internal/domain/models"
	"testing"
)

func TestPostUseCase_CommentBlock(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		id           models.PostID
		commentBlock bool
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Post
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{id: 1, commentBlock: false},
			want: &models.Post{
				ID:           1,
				CommentBlock: false,
			},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().CommentBlock(mock.Anything, args.id, args.commentBlock).
					Return(nil)
			},
		},
		{
			name:    "repo returns error",
			args:    args{id: 2, commentBlock: true},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().
					CommentBlock(mock.Anything, args.id, args.commentBlock).
					Return(errTest)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			usecase, mocks := makeService(t)

			if tc.before != nil {
				tc.before(mocks, tc.args)
			}
			err := usecase.CommentBlock(context.Background(), tc.args.id, tc.args.commentBlock)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)

		})
	}
}

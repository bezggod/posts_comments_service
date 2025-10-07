package post_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestPostUseCase_Create(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		userID       models.UserID
		title        string
		body         string
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
			args: args{
				userID:       1,
				title:        "test title",
				body:         "test body",
				commentBlock: false,
			},
			want: &models.Post{
				UserID:       1,
				Title:        "test title",
				Body:         "test body",
				CommentBlock: false,
			},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().Create(mock.Anything, args.userID, args.title, args.body, args.commentBlock).Return(&models.Post{
					UserID:       args.userID,
					Title:        args.title,
					Body:         args.body,
					CommentBlock: args.commentBlock,
				}, nil)
			},
		},
		{
			name: "error on postRepo.Create",
			args: args{
				userID:       1,
				title:        "bad title",
				body:         "bad body",
				commentBlock: false,
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().Create(mock.Anything, args.userID, args.title, args.body, args.commentBlock).Return(nil, errTest)
			},
		},
		{
			name: "empty title",
			args: args{
				userID:       1,
				title:        "",
				body:         "test body",
				commentBlock: false,
			},
			wantErr: true,
			before:  nil,
		},
		{
			name: "empty body",
			args: args{
				userID:       1,
				title:        "test title",
				body:         "",
				commentBlock: false,
			},
			wantErr: true,
			before:  nil,
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
			post, err := usecase.Create(context.Background(), tc.args.userID, tc.args.title, tc.args.body, tc.args.commentBlock)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, post)
		})
	}
}

package comment_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestCommentUseCase_CreateReplyComment(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		postID          models.PostID
		userID          models.UserID
		parentCommentID models.CommentID
		text            string
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Comment
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				postID:          1,
				userID:          1,
				parentCommentID: 3,
				text:            "ok",
			},
			want: &models.Comment{
				Text: "ok",
			},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().GetByID(mock.Anything, args.postID).Return(&models.Post{
					ID:           args.postID,
					CommentBlock: false,
				}, nil)
				m.commentRepo.EXPECT().CreateReplyComment(mock.Anything, args.postID, args.userID, args.parentCommentID, args.text).Return(&models.Comment{Text: "ok"}, nil)
			},
		},
		{
			name: "error on postRepo.GetByID",
			args: args{
				postID:          1,
				userID:          2,
				parentCommentID: 3,
				text:            "ok",
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().GetByID(mock.Anything, args.postID).Return(nil, errTest)
			},
		}, {
			name: "error postRepo.CommentBlock",
			args: args{
				postID:          1,
				userID:          2,
				parentCommentID: 3,
				text:            "ok",
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().GetByID(mock.Anything, args.postID).Return(&models.Post{CommentBlock: true}, nil)
			},
		},
		{
			name: "empty text",
			args: args{
				postID:          1,
				userID:          2,
				parentCommentID: 3,
				text:            "",
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
			comment, err := usecase.CreateReplyComment(context.Background(), tc.args.postID, tc.args.userID, tc.args.parentCommentID, tc.args.text)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, comment)
		})
	}
}

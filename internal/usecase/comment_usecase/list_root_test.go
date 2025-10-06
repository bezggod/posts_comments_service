package comment_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestCommentUseCase_ListRoots(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		postID models.PostID
		limit  int
		lastID *models.CommentID
	}

	tests := []struct {
		name    string
		args    args
		want    []*models.Comment
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				postID: 1,
				limit:  10,
				lastID: nil,
			},
			want: []*models.Comment{{
				Text: "ok",
			},
			},
			before: func(m mockService, args args) {
				m.commentRepo.EXPECT().ListRoots(mock.Anything, args.postID, args.limit, args.lastID).Return([]*models.Comment{{Text: "ok"}}, nil, nil)
			},
		},
		{
			name: "error on commentRepo.ListRoots",
			args: args{
				postID: 1,
				limit:  10,
				lastID: nil,
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.commentRepo.EXPECT().ListRoots(mock.Anything, args.postID, args.limit, args.lastID).Return(nil, nil, errTest)
			},
		},
		{
			name: "empty postID",
			args: args{
				postID: 0,
				limit:  10,
				lastID: nil,
			},
			wantErr: true,
			before:  nil,
		},
		{
			name: "limit<0",
			args: args{
				postID: 1,
				limit:  -5,
				lastID: nil,
			},
			want: []*models.Comment{},
			before: func(m mockService, args args) {
				m.commentRepo.EXPECT().ListRoots(mock.Anything, args.postID, 0, args.lastID).Return([]*models.Comment{}, nil, nil)
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
			comment, _, err := usecase.ListRoots(context.Background(), tc.args.postID, tc.args.limit, tc.args.lastID)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, comment)
		})
	}
}

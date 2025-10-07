package post_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestPostUseCase_ListPosts(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		limit  int
		lastID *models.PostID
	}

	tests := []struct {
		name    string
		args    args
		want    []*models.Post
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				limit:  10,
				lastID: nil,
			},
			want: []*models.Post{
				{
					Title: "ok",
				},
			},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().ListPosts(mock.Anything, args.limit, args.lastID).Return([]*models.Post{{Title: "ok"}}, nil, nil)
			},
		},
		{
			name: "error on postRepo.ListPosts",
			args: args{
				limit:  10,
				lastID: nil,
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().ListPosts(mock.Anything, args.limit, args.lastID).Return(nil, nil, errTest)
			},
		},
		{
			name: "limit<0",
			args: args{
				limit:  -5,
				lastID: nil,
			},
			want: []*models.Post{},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().ListPosts(mock.Anything, 0, args.lastID).Return([]*models.Post{}, nil, nil)
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
			post, _, err := usecase.ListPosts(context.Background(), tc.args.limit, tc.args.lastID)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, post)
		})
	}
}

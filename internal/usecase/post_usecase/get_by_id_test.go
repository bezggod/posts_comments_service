package post_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestPostUseCase_GetByID(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		id models.PostID
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
			args: args{id: 1},
			want: &models.Post{
				ID:           1,
				UserID:       2,
				Title:        "test title",
				Body:         "test body",
				CommentBlock: false,
			},
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().GetByID(mock.Anything, args.id).Return(&models.Post{
					ID:           1,
					UserID:       2,
					Title:        "test title",
					Body:         "test body",
					CommentBlock: false,
				}, nil)
			},
		},
		{
			name:    "repo error",
			args:    args{id: 1},
			wantErr: true,
			before: func(m mockService, args args) {
				m.postRepo.EXPECT().GetByID(mock.Anything, args.id).Return(nil, errTest)
			},
		},
		{
			name:    "epmty id",
			args:    args{id: 0},
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
			comment, err := usecase.GetByID(context.Background(), tc.args.id)
			if tc.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(tc.want, comment)
		})
	}
}

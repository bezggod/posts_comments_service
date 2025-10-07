package comment_usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"posts_commets_service/internal/domain/models"
	"testing"
)

func TestCommentUseCase_GetByID(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type args struct {
		id models.CommentID
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
			args: args{id: 1},
			want: &models.Comment{
				ID:   1,
				Text: "test comment",
			},
			before: func(m mockService, args args) {
				m.commentRepo.EXPECT().GetByID(mock.Anything, args.id).Return(&models.Comment{
					ID:   1,
					Text: "test comment",
				}, nil)
			},
		},
		{
			name:    "error on GetByID",
			args:    args{id: 99},
			wantErr: true,
			before: func(m mockService, args args) {
				m.commentRepo.EXPECT().GetByID(mock.Anything, args.id).Return(nil, errTest)
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

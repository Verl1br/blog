package repository

import (
	"database/sql"
	"testing"

	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

// TODO: change error test

func TestComment_CreateComment(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewCommentRepository(db)

	type args struct {
		item model.Comment
	}

	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				item: model.Comment{
					UserId: 1,
					PostId: 1,
					Text:   "test",
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO posts_comments").
					WithArgs(args.item.UserId, args.item.PostId, args.item.Text).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.CreateComment(tt.input.item)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestComment_GetByIdComment(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewCommentRepository(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		want    model.Comment
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "post_id", "comment"}).
					AddRow(1, 1, 1, "test")

				mock.ExpectQuery("SELECT (.+) FROM posts_comments WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: model.Comment{Id: 1, UserId: 1, PostId: 1, Text: "test"},
		},
		/*{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id"})

				mock.ExpectQuery("SELECT (.+) FROM posts WHERE (.+)").
					WithArgs(404, 1).WillReturnRows(rows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetComment(1)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestComment_GetComments(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewCommentRepository(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []model.Comment
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "post_id", "comment"}).
					AddRow(1, 1, 1, "test").
					AddRow(2, 1, 1, "test").
					AddRow(3, 1, 1, "test")

				mock.ExpectQuery("SELECT (.+) FROM posts_comments WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			want: []model.Comment{
				{Id: 1, UserId: 1, PostId: 1, Text: "test"},
				{Id: 2, UserId: 1, PostId: 1, Text: "test"},
				{Id: 3, UserId: 1, PostId: 1, Text: "test"},
			},
			input: args{
				id: 1,
			},
		},
		/*{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id"})

				mock.ExpectQuery("SELECT (.+) FROM posts WHERE (.+)").
					WithArgs(404).WillReturnRows(rows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetComments(1)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestComment_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewCommentRepository(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM posts_comments WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM posts_comments WHERE (.+)").
					WithArgs(404).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteComment(tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestComment_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewCommentRepository(db)

	type args struct {
		postId int
		input  model.UpdateComment
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK_AllFields",
			mock: func() {
				mock.ExpectExec("UPDATE posts_comments SET (.+) WHERE (.+)").
					WithArgs("test", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				postId: 1,
				input: model.UpdateComment{
					Text: stringPointer("test"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.UpdateComment(1, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

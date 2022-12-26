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

func TestPost_CreatePost(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewPostRepository(db)

	type args struct {
		item model.Post
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
				item: model.Post{
					Title:   "Test title",
					Content: "Test content",
					UserId:  1,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO posts").
					WithArgs(args.item.Title, args.item.Content, args.item.UserId).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.CreatePost(tt.input.item)
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

func TestPost_GetByIdPost(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewPostRepository(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		want    model.Post
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id"}).
					AddRow(1, "test", "test", 1)

				mock.ExpectQuery("SELECT (.+) FROM posts WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want: model.Post{Id: 1, Title: "test", Content: "test", UserId: 1},
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

			got, err := r.GetPost(tt.input.id, 1)
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

func TestPost_GetPosts(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewPostRepository(db)

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []model.Post
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id"}).
					AddRow(1, "test", "test", 1).
					AddRow(2, "test", "test", 1).
					AddRow(3, "test", "test", 1)

				mock.ExpectQuery("SELECT (.+) FROM posts WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			want: []model.Post{
				{Id: 1, Title: "test", Content: "test", UserId: 1},
				{Id: 2, Title: "test", Content: "test", UserId: 1},
				{Id: 3, Title: "test", Content: "test", UserId: 1},
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

			got, err := r.GetPosts(tt.input.id)
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

func TestPost_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewPostRepository(db)

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
				mock.ExpectExec("DELETE FROM posts WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM posts WHERE (.+)").
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

			err := r.DeletePost(tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := repository.NewPostRepository(db)

	type args struct {
		postId int
		userId int
		input  model.UpdatePost
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
				mock.ExpectExec("UPDATE posts SET (.+) WHERE (.+)").
					WithArgs("test", "test", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				postId: 1,
				userId: 1,
				input: model.UpdatePost{
					Title:   stringPointer("test"),
					Content: stringPointer("test"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.UpdatePost(1, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func stringPointer(s string) *string {
	return &s
}

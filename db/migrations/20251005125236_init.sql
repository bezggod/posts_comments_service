-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
    (
        id bigserial primary key,
        name text not null
);

CREATE TABLE posts
(
    id bigserial primary key,
    user_id bigint not null references users(id),
    title varchar(300) not null,
    body varchar(2000) not null,
    comment_block boolean default false,
    created_at timestamptz  default now()
);

CREATE TABLE comments
(
    id bigserial primary key,
    post_id bigint not null references posts(id),
    user_id bigint not null references users(id),
    text varchar(2000) not null,
    parent_comment_id bigint null references comments(id),
    first_comment_id bigint null,
    created_at timestamptz default now()
);

CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_parent_comment_id ON comments(parent_comment_id);
CREATE INDEX IF NOT EXISTS idx_comments_first_comment_id ON comments(first_comment_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments, posts, users CASCADE;
-- +goose StatementEnd

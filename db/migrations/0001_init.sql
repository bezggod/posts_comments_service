CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key,
    name text not null
);

CREATE TABLE IF NOT EXISTS posts (
    id bigserial primary key,
    user_id bigint not null references users(id),
    title varchar(300) not null,
    body varchar(2000) not null,
    comment_block boolean default false,
    created_at timestamptz default now()
    );

CREATE TABLE IF NOT EXISTS comments (
    id bigserial primary key,
    post_id bigint not null references posts(id),
    user_id bigint not null references users(id),
    text varchar(2000) not null,
    parent_comment_id bigint null references comments(id),
    first_comment_id bigint null,
    created_at timestamptz default now()
    );
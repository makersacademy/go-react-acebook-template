DROP TABLE IF EXISTS comments;
DROP SEQUENCE IF EXISTS comments_id_seq;
DROP TABLE IF EXISTS posts;
DROP SEQUENCE IF EXISTS posts_id_seq;
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;

CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE users (
    id text PRIMARY KEY,
    email text,
    username text,
    password text
);

-- Then the table with the foreign key second.
CREATE SEQUENCE IF NOT EXISTS posts_id_seq;
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content text,
    likes int,
-- The foreign key name is always {other_table_singular}_id
    user_id text,
    constraint fk_user foreign key(user_id)
        references users(id)
        on delete cascade
);

CREATE SEQUENCE IF NOT EXISTS comments_id_seq;
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content text,
    likes int,
-- The foreign key name is always {other_table_singular}_id
    post_id int,
    constraint fk_post foreign key(post_id)
        references posts(id)
        on delete cascade
);

-- Seed the table with data 
-- users first

INSERT INTO users (id, email, username, password)
VALUES
    ('9dc86cat-37bb-4c02-aab9-108ed9b2a261', 'cat@catmail.com', 'SuperCat', 'cat.123'),
    ('9dc86dog-37bb-4c02-aab9-108ed9b2a262', 'dog@dogmail.com', 'LoyalDog', 'dog.456'),
    ('9dc86rab-37bb-4c02-aab9-108ed9b2a263', 'rabbit@rabmail.com', 'BBunny', 'looney.123');

INSERT INTO posts (content, likes, user_id)
VALUES
    ('Post of a cat', 64, '9dc86cat-37bb-4c02-aab9-108ed9b2a261'),
    ('Post of a dog', 16, '9dc86dog-37bb-4c02-aab9-108ed9b2a262'),
    ('Another post of a dog', 134, '9dc86dog-37bb-4c02-aab9-108ed9b2a262'),
    ('A rabbit post', 92, '9dc86rab-37bb-4c02-aab9-108ed9b2a263');

INSERT INTO comments (content, likes, post_id)
VALUES
    ('Comment to a cat post', 14, 1),
    ('One more comment to the cat post', 54, 1);


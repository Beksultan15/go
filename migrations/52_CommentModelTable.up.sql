CREATE TABLE IF NOT EXISTS comments(
    id bigserial PRIMARY KEY,
    user_id int NOT NULL REFERENCES users ON DELETE CASCADE,
    product_id int NOT NULL REFERENCES products ON DELETE CASCADE,
    comment_body text NOT NULL,
    rating int NOT NULL DEFAULT 0
    );

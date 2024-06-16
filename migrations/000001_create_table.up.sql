CREATE TABLE IF NOT EXISTS book(
   id uuid PRIMARY key DEFAULT gen_random_uuid(),
    title VARCHAR(100),
    author_id uuid REFERENCES author(id),
    genre_id uuid REFERENCES genre(id),
    summary text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE IF NOT EXISTS author(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    biography text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE IF NOT EXISTS genre(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);

CREATE TABLE IF NOT EXISTS borrower(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    book_id uuid REFERENCES book(id),
    borrow_date TIMESTAMP,
    return_date TIMESTAMP  DEFAULT null,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at bigint DEFAULT 0
);
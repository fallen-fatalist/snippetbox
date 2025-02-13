-- Create a `snippets` table.
CREATE TABLE snippets (
    snippet_id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires TIMESTAMP NOT NULL
);




INSERT INTO snippets (title, content, created_at, expires) VALUES (
    'An old silent pond',
    E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    NOW(),
    NOW() + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created_at, expires) VALUES (
    'Over the wintry forest',
    E'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    NOW(),
    NOW() + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created_at, expires) VALUES (
    'First autumn morning',
    E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    NOW(),
    NOW() + INTERVAL '7 days'
);

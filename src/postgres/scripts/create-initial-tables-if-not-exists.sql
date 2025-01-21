SELECT 'CREATE DATABASE notes_database'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'notes_database')\gexec;

CREATE IF NOT EXISTS TABLE notes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
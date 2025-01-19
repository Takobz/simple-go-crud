SELECT 'CREATE DATABASE notes_database'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'notes_database')\gexec;


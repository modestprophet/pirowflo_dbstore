-- +goose Up
CREATE ROLE pirowdbuser WITH LOGIN PASSWORD 'secure_password'; -- Replace with actual password

GRANT CONNECT ON DATABASE plumbus TO pirowdbuser;
GRANT USAGE ON SCHEMA fitness TO pirowdbuser;
GRANT SELECT, INSERT, UPDATE, DELETE ON fitness.waterrower TO pirowdbuser;

-- +goose Down
REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA fitness FROM pirowdbuser;
REVOKE USAGE ON SCHEMA fitness FROM pirowdbuser;
REVOKE CONNECT ON DATABASE plumbus FROM pirowdbuser;
DROP ROLE pirowdbuser;
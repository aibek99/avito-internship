-- +goose Up
-- +goose StatementBegin
CREATE TYPE status_type AS ENUM (
    'CREATED',
    'PUBLISHED',
    'CLOSED'
);

CREATE TABLE tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type VARCHAR(100) NOT NULL,
    status status_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,
    version INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tender_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type VARCHAR(100) NOT NULL,
    status status_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,
    version INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE offer (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status status_type,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,
    version INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE offer_history (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   name VARCHAR(100) NOT NULL,
   description TEXT,
   status status_type,
   tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
   organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
   creator_username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,
   version INTEGER DEFAULT 1,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tender;
DROP TABLE tender_history;
DROP TABLE offer;
DROP TABLE offer_history;
-- +goose StatementEnd

CREATE TABLE places (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    type INTEGER REFERENCES place_types(id) NOT NULL,
    distributor_id UUID REFERENCES distributors(id) ON DELETE CASCADE NOT NULL,
    street_id UUID NOT NULL,
    house_number VARCHAR(8) NOT NULL,
    location GEOGRAPHY(POINT, 4326) NOT NULL,
    total_score INT DEFAULT 0 CHECK (total_score >= 0) NOT NULL,
    reviews_count INT DEFAULT 0 CHECK (reviews_count >= 0) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE distributors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    owner_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE distributors_staff (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    distributors_id UUID NOT NULL REFERENCES distributors(id) ON DELETE CASCADE,
    users_id UUID NOT NULL,
    role TEXT CHECK (role IN ('manager', 'staff', 'owner')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE places_staff (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    place_id UUID NOT NULL REFERENCES places(id) ON DELETE CASCADE,
    users_id UUID NOT NULL,
    role TEXT CHECK (role IN ('manager', 'staff', 'owner')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE place_types (
    id INTEGER PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE
);
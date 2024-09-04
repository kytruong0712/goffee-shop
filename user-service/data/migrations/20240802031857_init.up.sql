CREATE TABLE IF NOT EXISTS users (
    id                       BIGINT                                 NOT NULL PRIMARY KEY,
    iam_id                   BIGINT                                 NOT NULL CONSTRAINT users_iam_id_check CHECK (iam_id > 0 :: BIGINT),
    full_name                TEXT                                   NOT NULL CONSTRAINT users_full_name_check CHECK (full_name <> '' :: TEXT),
    hashed_password          TEXT                                   NOT NULL CONSTRAINT users_hashed_password_check CHECK (hashed_password <> '' :: TEXT),
    status                   TEXT                                   NOT NULL CONSTRAINT users_status_check CHECK (status <> '' :: TEXT),
    hashed_otp               TEXT,
    phone_number             TEXT                                   NOT NULL CONSTRAINT users_phone_number_check CHECK (phone_number <> '' :: TEXT),
    phone_number_verified    BOOL                                   NOT NULL DEFAULT FALSE,
    created_at               TIMESTAMP WITH TIME ZONE               NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMP WITH TIME ZONE               NOT NULL DEFAULT NOW(),
    otp_expiry_time          TIMESTAMP WITH TIME ZONE,
    CONSTRAINT users_full_name_phone_number_uindex UNIQUE (full_name, phone_number)
);
CREATE INDEX IF NOT EXISTS users_iam_id_index ON users(iam_id);


CREATE TABLE IF NOT EXISTS user_profiles (
    id                  BIGINT                                 NOT NULL PRIMARY KEY,
    user_id             BIGINT                                 NOT NULL CONSTRAINT user_id_fkey REFERENCES users (id) UNIQUE,
    email               TEXT,
    gender              TEXT,
    date_of_birth       TIMESTAMP,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT user_profiles_email_uindex UNIQUE (email)
);
CREATE INDEX IF NOT EXISTS user_profiles_email_index ON user_profiles(email);

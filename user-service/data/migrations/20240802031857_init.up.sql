CREATE TABLE IF NOT EXISTS users (
    id                  BIGINT                                 NOT NULL PRIMARY KEY,
    iam_id              BIGINT                                 NOT NULL CONSTRAINT users_iam_id_check CHECK (iam_id > 0 :: BIGINT),
    full_name           TEXT                                   NOT NULL CONSTRAINT users_full_name_check CHECK (full_name <> '' :: TEXT),
    phone_number        TEXT                                   NOT NULL CONSTRAINT users_phone_number_check CHECK (phone_number <> '' :: TEXT),
    password_hashed     TEXT                                   NOT NULL CONSTRAINT users_password_hashed_check CHECK (password_hashed <> '' :: TEXT),
    status              TEXT                                   NOT NULL CONSTRAINT users_status_check CHECK (status <> '' :: TEXT),
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT users_full_name_phone_number_index UNIQUE (full_name, phone_number)
);
CREATE INDEX IF NOT EXISTS users_iam_id_index ON users(iam_id);


CREATE TABLE IF NOT EXISTS user_profiles (
    id                  BIGINT                                 NOT NULL PRIMARY KEY,
    user_id             BIGINT                                 NOT NULL CONSTRAINT user_id_fkey REFERENCES users (id),
    email               TEXT,
    gender              TEXT,
    date_of_birth       TIMESTAMP,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    CONSTRAINT user_profiles_email_index UNIQUE (email)
);

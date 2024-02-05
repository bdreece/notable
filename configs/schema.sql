CREATE TABLE users (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    first_name varchar(63) NOT NULL,
    last_name varchar(63) NOT NULL,
    bio text NULL,
    email_address varchar(127) NOT NULL,
    email_verified boolean NOT NULL DEFAULT false,
    hash bytea NOT NULL
);

CREATE INDEX IX_users_name ON users (last_name ASC, first_name ASC);
CREATE UNIQUE INDEX IX_users_email_address ON users (email_address ASC);

CREATE TABLE devices (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    mac_address macaddr NOT NULL,
    ip_address inet NULL,
    connected boolean NOT NULL,
    user_id uuid NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE UNIQUE INDEX IX_devices_mac_address ON devices (mac_address ASC);
CREATE INDEX IX_devices_ip_address ON devices (ip_address ASC);
CREATE INDEX IX_devices_connected ON devices (user_id ASC, connected ASC);

CREATE TABLE directories (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    name varchar(127) NOT NULL,
    description text NULL,
    parent_id uuid NULL,
    user_id uuid NOT NULL,

    FOREIGN KEY (parent_id) REFERENCES directories (id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE INDEX IX_directories_users_name ON directories (user_id ASC, name ASC);
CREATE INDEX IX_directories_parent ON directories (parent_id ASC);

CREATE TABLE notes (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    title varchar(127) NOT NULL,
    content jsonb NOT NULL,
    directory_id uuid NOT NULL,
    user_id uuid NOT NULL,

    FOREIGN KEY (directory_id) REFERENCES directories (id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,

    FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE INDEX IX_notes_users_directories_title ON notes (user_id ASC, directory_id ASC, title ASC);

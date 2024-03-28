CREATE TABLE internal__fields (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL,

    type VARCHAR(255) NOT NULL,
    is_unique BOOLEAN NOT NULL,
    is_primary BOOLEAN NOT NULL,
    is_autoincrement BOOLEAN NOT NULL,
    is_nullable BOOLEAN NOT NULL,
    is_foreignkey BOOLEAN NOT NULL,
    reference VARCHAR(255) NOT NULL,
    related VARCHAR(255) NOT NULL,
    on_delete VARCHAR(255) NOT NULL,

    table_id INTEGER NOT NULL,

    CONSTRAINT fk_internal__table_field FOREIGN KEY (table_id) REFERENCES internal__tables(id)
);


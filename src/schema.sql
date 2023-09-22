CREATE TABLE template (
    id SERIAL PRIMARY KEY,
    app_id INTEGER NOT NULL,
    level_count INTEGER NOT NULL DEFAULT 0,
    condition VARCHAR(5)
);

CREATE TABLE approver (
    id INTEGER NOT NULL,
    email VARCHAR(255),
    template_id INTEGER REFERENCES template(id) ON DELETE CASCADE,
    PRIMARY KEY(id, template_id)
);

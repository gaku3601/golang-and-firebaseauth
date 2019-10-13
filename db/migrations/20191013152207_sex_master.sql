
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE sex_master (
    id serial,
    logical_value varchar(255),
    PRIMARY KEY(id)
);
INSERT INTO sex_master (logical_value) VALUES('男');
INSERT INTO sex_master (logical_value) VALUES('女');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE sex_master;

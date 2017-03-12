-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE pilots (
  id serial primary key NOT NULL,
  name text NOT NULL,
  hobbies text[] NOT NULL,
  created_at TIMESTAMP not NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE jets (
  id serial primary key NOT NULL,
  pilot_id integer NOT NULL,
  age integer NOT NULL,
  name text NOT NULL,
  color text NOT NULL,
  created_at TIMESTAMP not NULL,
  updated_at TIMESTAMP NOT NULL,

  foreign key(pilot_id) REFERENCES pilots (id)
);

CREATE TABLE languages (
  id serial primary key NOT NULL,
  language text NOT NULL,
  created_at TIMESTAMP not NULL,
  updated_at TIMESTAMP NOT NULL
);

-- Join table
CREATE TABLE pilot_languages (
  pilot_id integer NOT NULL,
  language_id integer NOT NULL,

  PRIMARY KEY (pilot_id, language_id),

  FOREIGN KEY (pilot_id) REFERENCES pilots (id),
  FOREIGN KEY (language_id) REFERENCES languages (id)
);

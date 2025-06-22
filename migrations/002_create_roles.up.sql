CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE
);

CREATE TABLE user_roles (
  user_id INT REFERENCES users(id),
  role_id INT REFERENCES roles(id),
  PRIMARY KEY (user_id, role_id)
);
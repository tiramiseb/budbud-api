PRAGMA foreign_keys = ON;

-- Authn

CREATE TABLE user (
  email TEXT PRIMARY KEY,
  passhash TEXT,
  passsalt TEXT
);

CREATE TABLE token (
  token TEXT PRIMARY KEY,
  user_email TEXT,
  FOREIGN KEY(user_email) REFERENCES user(email)
);

-- Ownership

CREATE TABLE workspace (
  id INTEGER PRIMARY KEY,
  owner_email TEXT,
  name TEXT,
  UNIQUE(owner_email, name),
  FOREIGN KEY(owner_email) REFERENCES user(email)
);

CREATE TABLE workspace_guest (
  workspace_id TEXT NOT NULL,
  user_email TEXT NOT NULL,
  FOREIGN KEY(workspace_id) REFERENCES workspace(id),
  FOREIGN KEY(user_email) REFERENCES user(email)
);

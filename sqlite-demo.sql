PRAGMA foreign_keys = ON;

INSERT INTO user VALUES ("test@example.com", "X", "X");
INSERT INTO user VALUES ("foo@example.com", "X", "X");

INSERT INTO workspace VALUES (1, "test@example.com", "My workspace");
INSERT INTO workspace VALUES (2, "foo@example.com", "Foo workspace");

INSERT INTO workspace_guest VALUES (2, "test@example.com");
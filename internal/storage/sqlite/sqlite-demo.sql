PRAGMA foreign_keys = ON;

INSERT INTO user VALUES ("test@example.com", "X", "X");
INSERT INTO user VALUES ("foo@example.com", "X", "X");

INSERT INTO workspace VALUES (1, "test@example.com", "My workspace");
INSERT INTO workspace VALUES (2, "foo@example.com", "Foo workspace");

INSERT INTO workspace_guest VALUES (2, "test@example.com");

INSERT INTO supercategory VALUES (1, "Income", 1);
INSERT INTO supercategory VALUES (2, "Leisure", 1);
INSERT INTO supercategory VALUES (3, "Videogames", 2);
INSERT INTO supercategory VALUES (4, "Regular expenses", 1);

INSERT INTO category VALUES (1, "Salary", 1);
INSERT INTO category VALUES (2, "Restaurant", 2);
INSERT INTO category VALUES (3, "School", 4);
INSERT INTO category VALUES (4, "Xbox", 3);
INSERT INTO category VALUES (5, "Switch", 3);

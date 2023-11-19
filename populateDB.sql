-- USERS
INSERT INTO users (id, name, password_hash) VALUES (1, 'Fer', '$2a$10$S9A7.aQAeK.cnB8PdQ7DIOZrd768freWDP5fZVBIZ5Yu3Zq9su.wu');
INSERT INTO users (id, name, password_hash) VALUES (2, 'Toggl', '$2a$10$I6Shy2xb3osQhoFxghG6wOQbHUZxlwEIiTjhYy136h/n4vxWhmDae');


-- QUESTIONS
INSERT INTO questions (id, body, user_id) VALUES (1, 'First question', 1);
INSERT INTO questions (id, body, user_id) VALUES (2, 'Second question', 1);
INSERT INTO questions (id, body, user_id) VALUES (3, 'Third question', 1);
INSERT INTO questions (id, body, user_id) VALUES (4, 'Forth question', 1);
INSERT INTO questions (id, body, user_id) VALUES (5, 'Fifth question', 1);

INSERT INTO questions (id, body, user_id) VALUES (6, 'Sixth question', 2);
INSERT INTO questions (id, body, user_id) VALUES (7, 'Seventh question', 2);
INSERT INTO questions (id, body, user_id) VALUES (8, 'Eighth question', 2);
INSERT INTO questions (id, body, user_id) VALUES (9, 'Nineth question', 2);
INSERT INTO questions (id, body, user_id) VALUES (10, 'Tenth question', 2);


-- OPTIONS
INSERT INTO options (id, body, correct, question_id) VALUES (1, 'One', true, 1);
INSERT INTO options (id, body, correct, question_id) VALUES (2, 'Two', true, 1);
INSERT INTO options (id, body, correct, question_id) VALUES (3, 'Three', false, 1);
INSERT INTO options (id, body, correct, question_id) VALUES (4, 'Four', false, 3);
INSERT INTO options (id, body, correct, question_id) VALUES (5, 'Five', true, 3);
INSERT INTO options (id, body, correct, question_id) VALUES (6, 'Six', true, 4);

INSERT INTO options (id, body, correct, question_id) VALUES (7, 'Seven', true, 6);
INSERT INTO options (id, body, correct, question_id) VALUES (8, 'Eight', false, 6);
INSERT INTO options (id, body, correct, question_id) VALUES (9, 'Nine', false, 6);
INSERT INTO options (id, body, correct, question_id) VALUES (10, 'Ten', true, 7);
INSERT INTO options (id, body, correct, question_id) VALUES (11, 'Eleven', true, 7);
INSERT INTO options (id, body, correct, question_id) VALUES (12, 'Twelve', true, 8);
CREATE TABLE IF NOT EXISTS `users`
(
	`id` int PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(255) UNIQUE,
	`real_name` varchar(255),
	`email` varchar(255),
	`created_at` datetime,
	`role_id` int,
	`pass_hash` varchar(255)
);

CREATE TABLE IF NOT EXISTS `privileges`
(
	`id` int PRIMARY KEY,
	`name` varchar(255)
);

CREATE TABLE IF NOT EXISTS `roles`
(
	`id` int PRIMARY KEY,
	`name` varchar(255)
);

CREATE TABLE IF NOT EXISTS `privileges_by_role`
(
	`priv_id` int,
	`role_id` int
);

ALTER TABLE `privileges_by_role` ADD FOREIGN KEY (`priv_id`) REFERENCES `privileges` (`id`);
ALTER TABLE `privileges_by_role` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

-- Create the data
INSERT INTO `roles` (id,name) VALUES (0,'Administrator');
INSERT INTO `roles` (id,name) VALUES (1,'User');

INSERT INTO `privileges` (id,name)
VALUES (0,'CAN_EDIT_TIMETABLE');
INSERT INTO `privileges` (id,name)
VALUES (1,'CAN_MAKE_USERS');
INSERT INTO `privileges` (id,name)
VALUES (2,'CAN_ADD_LECTURES');
INSERT INTO `privileges` (id,name)
VALUES (3,'CAN_ADD_SPEAKERS');

-- Administrator permissions
INSERT INTO `privileges_by_role` (role_id,priv_id)
VALUES (0,0);
INSERT INTO `privileges_by_role` (role_id,priv_id)
VALUES (0,1);
INSERT INTO `privileges_by_role` (role_id,priv_id)
VALUES (0,2);
-- User permissions
INSERT INTO `privileges_by_role` (role_id,priv_id)
VALUES (1,2);
INSERT INTO `privileges_by_role` (role_id,priv_id)
VALUES (1,3);

INSERT INTO `users` ( name, real_name, email, created_at, role_id, pass_hash)
VALUES ('admin','admin','null@nospam.com',NOW(),0,'admin');
INSERT INTO `users` ( name, real_name, email, created_at, role_id, pass_hash)
VALUES ('user','user','null@nospam.com',NOW(),1,'user');

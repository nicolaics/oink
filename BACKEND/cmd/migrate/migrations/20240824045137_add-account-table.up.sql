CREATE TABLE `account` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `balance` float NOT NULL DEFAULT 0,

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
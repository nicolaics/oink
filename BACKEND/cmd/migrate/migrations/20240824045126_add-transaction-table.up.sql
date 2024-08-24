CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` float NOT NULL,
  `transaction_time` timestamp NOT NULL DEFAULT (now()),

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
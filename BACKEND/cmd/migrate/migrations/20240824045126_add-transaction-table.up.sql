CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `receiver_id` int NOT NULL,
  `sender_id` int NOT NULL,
  `amount` float NOT NULL,
  `transaction_time` timestamp NOT NULL DEFAULT (now()),

  FOREIGN KEY (`receiver_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`)
);
CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` float NOT NULL,
  `src_acc` varchar(256) NOT NULL,
  `dest_acc` varchar(256) NOT NULL,
  `visible` boolean,
  `transaction_time` timestamp NOT NULL DEFAULT (now()),

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
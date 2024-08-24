CREATE TABLE `loan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `debtor_id` int NOT NULL,
  `amount` float NOT NULL,
  `start_date` timestamp NOT NULL DEFAULT (now()),
  `end_date` timestamp NOT NULL,
  `duration` varchar(255) NOT NULL,

  FOREIGN KEY (`debtor_id`) REFERENCES `users` (`id`)
);
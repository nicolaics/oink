CREATE TABLE `loan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `debtor_id` int NOT NULL,
  `amount` float NOT NULL,
  `amount_paid` float NOT NULL DEFAULT 0.0,
  `start_date` timestamp NOT NULL DEFAULT (now()),
  `end_date` timestamp NOT NULL,
  `duration` varchar(255) NOT NULL,
  `active` boolean DEFAULT FALSE,

  FOREIGN KEY (`debtor_id`) REFERENCES `users` (`id`)
);
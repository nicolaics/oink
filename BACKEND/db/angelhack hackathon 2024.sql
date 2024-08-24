CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(512) NOT NULL,
  `password` text NOT NULL,
  `email` varchar(256) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `receiver_id` int NOT NULL,
  `sender_id` int NOT NULL,
  `amount` float NOT NULL,
  `transaction_time` timestamp NOT NULL DEFAULT (now()),

  FOREIGN KEY (`receiver_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `loan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `debtor_id` int NOT NULL,
  `amount` float NOT NULL,
  `start_date` timestamp NOT NULL DEFAULT (now()),
  `end_date` timestamp NOT NULL,
  `duration` varchar(255) NOT NULL,

  FOREIGN KEY (`debtor_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `account` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `balance` float NOT NULL DEFAULT 0,

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `savings_account` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` float NOT NULL DEFAULT 0,

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `pig_race` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `pig_stamina` float NOT NULL DEFAULT 0,
  `final_distance_to_goal` float,

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

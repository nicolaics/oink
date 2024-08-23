CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(512) UNIQUE NOT NULL,
  `password` text NOT NULL,
  `balance` float NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `transaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `receiver_id` int NOT NULL,
  `sender_id` int NOT NULL,
  `amount` float NOT NULL,
  `transaction_time` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `loan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `debtor_id` int NOT NULL,
  `amount` float NOT NULL,
  `start_date` timestamp NOT NULL DEFAULT (now()),
  `end_date` timestamp NOT NULL,
  `duration` varchar(255) NOT NULL
);

CREATE TABLE `savings_account` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` float NOT NULL DEFAULT 0,
  `debitted_for_loan` float
);

CREATE TABLE `pig_race` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `pig_stamina` float NOT NULL DEFAULT 0,
  `final_distance_to_goal` float
);

ALTER TABLE `transaction` ADD FOREIGN KEY (`receiver_id`) REFERENCES `users` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`);

ALTER TABLE `loan` ADD FOREIGN KEY (`debtor_id`) REFERENCES `users` (`id`);

ALTER TABLE `savings_account` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `pig_race` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

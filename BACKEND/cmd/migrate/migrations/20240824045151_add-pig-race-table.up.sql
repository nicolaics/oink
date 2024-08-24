CREATE TABLE `pig_race` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `pig_stamina` float NOT NULL DEFAULT 0,
  `final_distance_to_goal` float,

  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
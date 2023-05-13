drop database if exists `PersonalizedRecommendationSystem`;

create database `PersonalizedRecommendationSystem`;

use  `PersonalizedRecommendationSystem`;

-- 创建用户表
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `users` (`username`, `email`, `password`) VALUES 
('John Doe', 'john.doe@example.com', 'mypassword'),
('Jane Smith', 'jane.smith@example.com', 'mypassword'),
('Bob Johnson', 'bob.johnson@example.com', 'mypassword');


Select * from `users`;

-- 创建学习资源表
CREATE TABLE `resources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `url`  varchar(255) NOT NULL,
  `tag`  varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `resources` (`name`, `description`, `url`, `tag`)
VALUES ('Introduction to Python', 'A beginner-friendly Python course that covers the basics of programming with Python.', 'https://example.com/python-intro', "Python"),
       ('Node.js Fundamentals', 'An online course for learning Node.js', 'https://example.com/nodejs-fundamentals', "NodeJs"),
       ('Data Structures and Algorithms in Java', 'Learn about common data structures and algorithms using Java programming language.', 'https://example.com/data-structures-algorithms-java', "Java"),
       ('Data Structures and Algorithms in Rust', 'Learn about algorithms using Java programming language.', 'https://example.com/data-structures-algorithms-rust', "Rust"),
       ('Masterclass', 'A comprehensive course ', 'https://example.com/css-masterclass', "HTML"),
       ('Vue', 'A comprehensive course ', 'https://example.com/vue-masterclass', "Vue"),
       ('React', 'A comprehensive course ', 'https://example.com/react-masterclass', "React"),
       ('CSS Masterclass', 'A comprehensive course on mastering CSS', 'https://example.com/css-masterclass', "CSS");


Select * from `resources`;


-- 创建用户行为表
CREATE TABLE `user_actions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `action_type` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`) VALUES 
(1, 3, 'collect' ),
(1, 3, 'collect' ),
(1, 3, 'collect' ),
(1, 3, 'collect' ),
(1, 3, 'collect' ),
(2, 3, 'collect' ),
(3, 3, 'collect' ),
(4, 3, 'collect' );


INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`) VALUES 
(1, 1, 'browse'),
(2, 1, 'browse'),
(1, 2, 'browse'),
(3, 2, 'browse'),
(1, 3, 'browse'),
(2, 3, 'browse'),
(1, 4, 'browse'),
(2, 4, 'browse'),
(3, 4, 'browse'),
(1, 5, 'browse');


INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`) VALUES 
(1, 1, 'like'),
(2, 1, 'like'),
(1, 2, 'like'),
(3, 2, 'like'),
(3, 2, 'like');

INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`) VALUES 
(1, 1, 'save'),
(2, 1, 'save'),
(3, 1, 'save'),
(1, 2, 'save');

INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`) VALUES 
(1, 1, 'click'),
(2, 1, 'click'),
(3, 1, 'click'),
(4, 1, 'click'),
(1, 2, 'click');

Select * from `user_actions`;


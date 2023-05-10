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


-- 创建学习资源表
CREATE TABLE `resources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `cover_image` varchar(255),
  `url` varchar(255) NOT NULL,
  `type` varchar(50) NOT NULL,
  `difficulty` varchar(50),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `resources` (`name`, `description`, `cover_image`, `url`, `type`, `difficulty`) VALUES 
('Learn Python', 'A video course about Python programming language', 'python.jpg', 'https://example.com/python', 'video', 'beginner'),
('Java for Beginners', 'Get started with Java programming', 'java.jpg', 'https://example.com/java', 'article', 'beginner'),
('Web Development Bootcamp', 'Learn how to build web applications from scratch', 'webdev.jpg', 'https://example.com/webdev', 'course', 'intermediate'),
('Data Science Essentials', 'Introduction to Data Science concepts and tools', 'datascience.jpg', 'https://example.com/datascience', 'ebook', 'advanced'),
('Machine Learning Mastery', 'Master Machine Learning algorithms and techniques', 'ml.jpg', 'https://example.com/ml', 'course', 'advanced');


-- 创建用户行为表
CREATE TABLE `user_actions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `action_type` varchar(50) NOT NULL,
  `action_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `resource_id` (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`, `action_time`) VALUES 
(1, 1, 'collect', NOW()),
(2, 1, 'collect', NOW()),
(1, 2, 'collect', NOW()),
(3, 2, 'collect', NOW()),
(1, 3, 'collect', NOW());


INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`, `action_time`) VALUES 
(1, 1, 'browse', NOW()),
(2, 1, 'browse', NOW()),
(1, 2, 'browse', NOW()),
(3, 2, 'browse', NOW()),
(1, 3, 'browse', NOW()),
(2, 3, 'browse', NOW()),
(1, 4, 'browse', NOW()),
(2, 4, 'browse', NOW()),
(3, 4, 'browse', NOW()),
(1, 5, 'browse', NOW());

INSERT INTO `user_actions` (`user_id`, `resource_id`, `action_type`, `action_time`) VALUES 
(1, 1, 'like', NOW()),
(2, 1, 'like', NOW()),
(1, 2, 'like', NOW());


-- 根据用户ID获取用户信息
SELECT * FROM users WHERE id = 1;

-- 根据资源类型和难度级别获取学习资源列表
SELECT * FROM resources WHERE type = 'video' AND difficulty = 'advanced';

-- 根据用户浏览记录推荐相关学习资源
SELECT r.* FROM user_actions ua
LEFT JOIN resources r ON ua.resource_id = r.id
WHERE ua.user_id = 1 AND ua.action_type = 'browse'
ORDER BY ua.action_time DESC
LIMIT 10;

-- 根据用户收藏记录推荐相关学习资源
SELECT r.* FROM user_actions ua
LEFT JOIN resources r ON ua.resource_id = r.id
WHERE ua.user_id = 1 AND ua.action_type = 'collect'
ORDER BY ua.action_time DESC
LIMIT 10;

-- 根据用户点赞记录推荐相关学习资源
SELECT r.* FROM user_actions ua
LEFT JOIN resources r ON ua.resource_id = r.id
WHERE ua.user_id = 1 AND ua.action_type = 'like'
ORDER BY ua.action_time DESC
LIMIT 10;

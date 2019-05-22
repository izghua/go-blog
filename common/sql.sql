
CREATE TABLE `z_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名',
  `display_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类别名',
  `seo_desc` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'seo描述',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父类ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `z_categories_display_name_index` (`display_name`),
  KEY `z_categories_parent_id_index` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_links` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '友链名',
  `link` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '友链链接',
  `order` int(11) NOT NULL COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `z_migrations` WRITE;
/*!40000 ALTER TABLE `z_migrations` DISABLE KEYS */;

INSERT INTO `z_migrations` (`id`, `migration`, `batch`)
VALUES
	(2,'2014_10_12_100000_create_password_resets_table',1),
	(3,'2018_12_26_144328_create_users_table',1),
	(4,'2018_12_26_145106_create_posts_table',1),
	(5,'2018_12_26_145124_create_categories_table',1),
	(6,'2018_12_26_145200_create_tags_table',1),
	(7,'2018_12_26_145222_create_post_tag_table',1),
	(8,'2018_12_26_145240_create_post_cate_table',1),
	(9,'2018_12_26_145258_create_post_views_table',1),
	(10,'2018_12_26_145340_create_systems_table',1),
	(11,'2018_12_26_145355_create_links_table',1);

/*!40000 ALTER TABLE `z_migrations` ENABLE KEYS */;
UNLOCK TABLES;



CREATE TABLE `z_password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  KEY `z_password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



CREATE TABLE `z_post_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) NOT NULL COMMENT '文章ID',
  `cate_id` int(11) NOT NULL COMMENT '分类ID',
  PRIMARY KEY (`id`),
  KEY `z_post_cate_post_id_index` (`post_id`),
  KEY `z_post_cate_cate_id_index` (`cate_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_post_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) NOT NULL COMMENT '文章ID',
  `tag_id` int(11) NOT NULL COMMENT '标签ID',
  PRIMARY KEY (`id`),
  KEY `z_post_tag_post_id_index` (`post_id`),
  KEY `z_post_tag_tag_id_index` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_post_views` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) NOT NULL COMMENT '文章ID',
  `num` int(11) NOT NULL COMMENT '阅读次数',
  PRIMARY KEY (`id`),
  KEY `z_post_views_post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'uid',
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `summary` char(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '摘要',
  `original` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原文章内容',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章密码',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `z_posts_user_id_index` (`user_id`),
  KEY `z_posts_uid_index` (`uid`),
  KEY `z_posts_title_index` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `z_systems` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `theme` tinyint(4) NOT NULL DEFAULT '0' COMMENT '主题',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '网站title',
  `keywords` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '网站关键字',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '网站描述',
  `record_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备案号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签名',
  `display_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签别名',
  `seo_desc` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'seo描述',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '被使用次数',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `z_tags_display_name_index` (`display_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `z_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '用户状态 0创建,1正常',
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `z_users_email_unique` (`email`),
  KEY `z_users_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



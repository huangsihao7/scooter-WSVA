CREATE TABLE `videos`  (
                           `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                           `author_id` int UNSIGNED NOT NULL COMMENT '上传用户Id',
                           `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
                           `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面url',
                           `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放url',
                           `favorite_count` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
                           `comment_count` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数目',
                           `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NULL DEFAULT NULL,
                           `deleted_at` datetime NULL DEFAULT NULL,
                           `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频分类',
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `author_id`(`author_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;
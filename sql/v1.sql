CREATE DATABASE blog;
USE blog;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
     `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
     `tag_id` int(10) UNSIGNED DEFAULT 0 COMMENT '标签ID',
     `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '文章标题',
     `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '简述',
     `content` text CHARACTER SET utf8 COLLATE utf8_general_ci,
     `created_on` int(11) DEFAULT NULL,
     `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '创建人',
     `modified_on` int(10) UNSIGNED DEFAULT 0 COMMENT '修改时间',
     `modified_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '修改人',
     `deleted_on` int(10) UNSIGNED DEFAULT 0,
     `state` tinyint(3) UNSIGNED DEFAULT 1 COMMENT '状态 0为禁用1为启用',
     PRIMARY KEY (`id`) USING BTREE
);

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (2, 6, 'test-edit1', 'test-desc-edit', 'test-content-edit', 0, 'yugu', 1679124676, 'test-created-edit', 0, 1);
INSERT INTO `blog_article` VALUES (3, 6, 'test', 'test desc', 'test function is can used', 0, 'yugu', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (4, 6, 'test', 'test desc', 'test function is can used', 0, 'yugu', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (5, 6, 'test', 'test desc', 'test function is can used', 0, 'yugu', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (6, 6, 'test', 'test desc', 'test function is can used', 0, 'yugu', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (7, 666, 'newArticle1', 'testCreate', '', 0, '', 0, '', 0, 1);

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'admin', '123456');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
     `id` int(10) NOT NULL AUTO_INCREMENT,
     `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '标签名称',
     `created_on` int(10) DEFAULT 0 COMMENT '创建时间',
     `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '创建人',
     `modified_on` int(10) DEFAULT 0 COMMENT '修改时间',
     `modified_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '修改人',
     `deleted_on` int(10) DEFAULT 0,
     `state` tinyint(3) DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
     PRIMARY KEY (`id`) USING BTREE
) ;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (3, 't1', 1679046852, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (4, 't11', 1679046864, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (5, 't111', 1679046866, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (6, 't2', 1679046870, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (7, 't22', 1679046872, 'test', 0, '', 0, 1);

SET FOREIGN_KEY_CHECKS = 1;

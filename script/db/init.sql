CREATE TABLE `schools` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `name` varchar(255) NOT NULL,
                           `brief_introduction` text,
                           `school_code` varchar(100) DEFAULT NULL,
                           `master_point` int DEFAULT NULL,
                           `phd_point` int DEFAULT NULL,
                           `research_project` int DEFAULT NULL,
                           `title_double_first_class` tinyint(1) DEFAULT NULL,
                           `title_985` tinyint(1) DEFAULT NULL,
                           `title_211` tinyint(1) DEFAULT NULL,
                           `title_college` tinyint(1) DEFAULT NULL,
                           `title_undergraduate` tinyint(1) DEFAULT NULL,
                           `region` varchar(255) DEFAULT NULL,
                           `website` varchar(255) DEFAULT NULL,
                           `recruitment_phone` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `email` varchar(100) DEFAULT NULL,
                           `promotion_rate` varchar(50) DEFAULT NULL,
                           `abroad_rate` varchar(50) DEFAULT NULL,
                           `employment_rate` varchar(50) DEFAULT NULL,
                           `double_first_class_disciplines` text,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `name` (`name`),
                           UNIQUE KEY `school_code` (`school_code`)
);

CREATE TABLE `scores` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `school_id` int DEFAULT NULL,
                          `location` int DEFAULT NULL,
                          `year` int DEFAULT NULL,
                          `type_id` int DEFAULT NULL,
                          `tag` varchar(50) DEFAULT NULL,
                          `lowest` int DEFAULT NULL,
                          `lowest_rank` int DEFAULT NULL,
                          `sg_name` varchar(50) DEFAULT NULL,
                          `batch_name` varchar(50) DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `scores_school_id_IDX` (`school_id`,`location`,`type_id`,`year`,`tag`,`sg_name`,`batch_name`) USING BTREE,
                          CONSTRAINT `scores_ibfk_1` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`)
);

create table majors (
                        id int auto_increment primary key, -- 主键ID
                        name varchar(255) not null, -- 名称
                        national_feature boolean, -- 国家特色
                        level varchar(100), -- 层次
                        discipline_category varchar(100), -- 学科门类
                        major_category varchar(100), -- 专业类别
                        duration int, -- 学制
                        school_id int, -- 所属院校ID
                        foreign key (school_id) references schools(id) -- 外键关联院校表
);

create table major_scores (
                        special_id int,
                        location varchar(50),
                        year     int,
                        kelei    varchar(50),
                        recruitment_number int, -- 招收人数
                        highest_score int, -- 最高分
                        lowest_score int, -- 最低分
                        highest_rank int, -- 最高名次
                        lowest_rank int -- 最低名次
);

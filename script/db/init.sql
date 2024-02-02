create table schools (
                         id int auto_increment primary key, -- 主键ID
                         name varchar(255) unique not null, -- 名称
                         brief_introduction text ,  -- 简介
                         school_code varchar(100) unique, -- 院校代码
                         master_point int, -- 硕士点
                         phd_point int, -- 博士点
                         research_project int, -- 科研项目
                         title_double_first_class boolean, -- 双一流
                         title_985 boolean, -- 985
                         title_211 boolean, -- 211
                         title_college boolean, -- 专科
                         title_undergraduate boolean, -- 本科
                         region varchar(255), -- 地区
                         website varchar(255), -- 官网
                         recruitment_phone varchar(50), -- 招生电话
                         email varchar(100), -- 邮箱
                         promotion_rate varchar(50), -- 升学率
                         abroad_rate varchar(50), -- 出国率
                         employment_rate varchar(50), -- 就业率
                         double_first_class_disciplines text -- 双一流学科
);

create table scores (
        id int auto_increment primary key, -- 主键ID
        school_id int,
        location int,
        year int,
        type_id int,
        tag  varchar(50), -- 批次
        lowest  int,
        lowest_rank int, -- 最低名次
        foreign key (school_id) references schools(id) -- 外键关联院校表
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

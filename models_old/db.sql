CREATE TABLE `address` 
(
  `city` varchar(128) NOT NULL ,
  `state` varchar(128) NOT NULL ,
  `country` varchar(128) NOT NULL ,
  `line_1` varchar(128) NOT NULL ,
  `line_1_check` varchar(128) NOT NULL ,
  `line_2` varchar(128) NOT NULL ,
  `zip` varchar(128) NOT NULL ,
  `zip_check` varchar(128) NOT NULL ,
  
) engine=innodb DEFAULT charset=utf8mb4;
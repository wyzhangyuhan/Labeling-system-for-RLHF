CREATE TABLE `users` (
  `id` bigint PRIMARY KEY,
  `username` varchar(25) UNIQUE NOT NULL,
  `password` varchar(100),
  `CreatTime` datetime NOT NULL,
  `UpdateTime` datetime NOT NULL
);

CREATE TABLE `dataCate` (
  `id` int PRIMARY KEY,
  `parentid` int,
  `catename` text,
  `CreatTime` datetime NOT NULL,
  `UpdateTime` datetime NOT NULL
);

CREATE TABLE `dataSet` (
  `id` int PRIMARY KEY,
  `cateid` int,
  `dataname` varchar(255),
  `CreatTime` datetime NOT NULL,
  `UpdateTime` datetime NOT NULL
);

CREATE TABLE `Model` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `CreatTime` datetime NOT NULL,  
  `UpdateTime` datetime NOT NULL
);

CREATE TABLE `dataItem` (
  `MessageId` int PRIMARY KEY,
  `SessionId` int,
  `Query` text,
  `Answer` text,
  `Edit` text,
  `CreatTime` datetime NOT NULL,
  `UpdateTime` datetime NOT NULL
);

CREATE TABLE `SessionItem` (
  `SessionId` int PRIMARY KEY,
  `DatasetId` int,
  `ModelId` int,
  `userid` int,
  `CreatTime` datetime NOT NULL,
  `UpdateTime` datetime NOT NULL
);

ALTER TABLE `dataSet` ADD FOREIGN KEY (`cateid`) REFERENCES `dataCate` (`id`);

ALTER TABLE `dataSet` ADD FOREIGN KEY (`id`) REFERENCES `SessionItem` (`DatasetId`);

ALTER TABLE `Model` ADD FOREIGN KEY (`id`) REFERENCES `SessionItem` (`ModelId`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `SessionItem` (`userid`);

ALTER TABLE `dataItem` ADD FOREIGN KEY (`MessageId`) REFERENCES `SessionItem` (`SessionId`);

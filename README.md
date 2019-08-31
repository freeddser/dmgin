
这是用GIN简单封装的一个脚手架

    common 公共方法

    flags 一些常量

    handlers 用来处理请求层

    services 业务逻辑层

    repository 数据层，cache

    models 数据层model

    request 请求层model

    response 返回层model

    route 路由

    main.go 主函数

    config.toml 配置文件


# Example
# 运行服务的预先要求
    *  1. Redis          (v. 4.0.9+)
    *  2. Mysql       (v. 5.6+)


    CREATE TABLE `rs_driver` (
    `driver_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `driver_name` varchar(255) NOT NULL DEFAULT ' ',
    `driver_mobile` varchar(31) NOT NULL DEFAULT ' ',
    `vehicle_id` bigint(20) NOT NULL DEFAULT '0',
     `vehicle_no` varchar(12) NOT NULL DEFAULT ' ',
      `status` char(8) NOT NULL DEFAULT '',
     `created_at` datetime NOT NULL,
     `updated_at` datetime NOT NULL,
    PRIMARY KEY (`driver_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2003 DEFAULT CHARSET=latin1

    mysql> select * from rs_driver;
    +-----------+-------------+---------------+------------+------------+--------+---------------------+---------------------+
    | driver_id | driver_name | driver_mobile | vehicle_id | vehicle_no | status | created_at          | updated_at          |
    +-----------+-------------+---------------+------------+------------+--------+---------------------+---------------------+
    |      2001 | Bob         | 89015566      |       3001 | BY81X      | 0      | 2019-01-01 11:00:00 | 2019-01-01 11:00:00 |
    |      2002 | Ken         | 81011166      |       3002 | GA110      | 0      | 2019-01-01 11:00:00 | 2019-01-01 11:00:00 |
    +-----------+-------------+---------------+------------+------------+--------+---------------------+---------------------+
    2 rows in set (0.01 sec)

###POST example:

    localhost:8080/login

    request:
    {
    "name":"gavin",
	"password":"123456"
    }

    response:

    {
    "status": "you are logged in"
    }

    {
     "status": "unauthorized"
    }

### Get example:

    localhost:8080/driver/2001

    {
     "msg": "success",
     "response": {
         "driver_id": 2001,
         "driver_name": "Bob",
         "driver_mobile": "89015566",
         "vehicle_id": "3001",
         "vehicle_no": "BY81X",
         "status": 0,
         "created_at": "2019-01-01 11:00:00",
         "updated_at": "2019-01-01 11:00:00"
     }
    }
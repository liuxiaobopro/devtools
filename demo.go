package main

type lxbUser struct {
	Id       int    `json:"id" xorm:"int(11) unsigned NOT NULL"`
	Nickname string `json:"nickname" xorm:"varchar(50) NOT NULL COMMENT"`
	Username string `json:"username" xorm:"varchar(50) NOT NULL COMMENT"`
	Password string `json:"password" xorm:"varchar(100) NOT NULL COMMENT"`
	Phone    string `json:"phone" xorm:"char(11) DEFAULT '' COMMENT"`
	RoleId   int    `json:"role_id" xorm:"int(11) NOT NULL DEFAULT '-1' COMMENT"`
}

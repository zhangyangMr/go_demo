package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	host := "192.168.0.151"
	port := 3307
	user := "root"
	dbname := "fabric"
	password := "root"
	Type := "mysql"
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		user, password, host, port, dbname)
	//dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
	//	user, password, host, port, dbname)
	db, err := gorm.Open(Type, dbDSN)
	if err != nil {
		fmt.Printf("open db err : %v\n", err)
	}
	err = db.DB().Ping()
	if err != nil {
		fmt.Printf("db ping err : %v\n", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	type TxEvent struct {
		TableName   string `json:"-" gorm:"-"`
		Id          int64  `json:"id" gorm:"primary_key;column:id"`
		TxKey       string `json:"tx_key" gorm:"column:tx_key"`
		TxID        string `json:"tx_id" gorm:"column:tx_id"`
		DataIndex   string `json:"data_index" gorm:"column:data_index"`
		FuncName    string `json:"func_name" gorm:"column:func_name"`
		Status      int    `json:"status" gorm:"column:status"`
		Data        string `json:"data" gorm:"size:65534"`
		Time        string `json:"time" gorm:"column:time"`
		BlockHeight int    `json:"block_height" gorm:"column:block_height"`
		NewTxId     string `json:"new_txid" gorm:"column:new_txid"`
	}
	migrateErr := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&TxEvent{}).Error
	if migrateErr != nil {
		fmt.Printf("migrate err : %v\n", migrateErr)
	}

	data := `{"business_id":"789425898231356077","business_system_name":"CAMS","operator_id":"517689777039941719","operation_time":"2020-12-10 11:05:50","meta_data":"{\"bos\":[{\"boDefId\":\"280097099663181116\",\"formKey\":\"168ybgcysjl\",\"data\":{\"btgcmc\":\"华润电力内黄县润电20MW分散式风电项目\",\"btgcmcx\":\"隐蔽工程质量签证报验单\",\"bhz\":\"CRNE—NHRH—A26\",\"bhy\":\"HNGC-A26-013\",\"bydgcmc\":\"华润电力内黄县润电20MW分散式风电项目\",\"z\":\"河北鸿泰融华润电力内黄县润电20MW分散式风电项目监理部\",\"fxgc\":\"基础接地\",\"gcmc\":\"华润电力内黄县润电20MW分散式风电项目\",\"bh\":\"2FJ010003000101001\",\"dwgcmc\":\"1号风力发电机组\",\"fxgcmc\":\"基础接地\",\"sgdw\":\"中国电建集团河南工程有限公司\",\"jldw\":\"河北鸿泰融华润电力内黄县润电20MW分散式风电项目监理部\",\"hjwd\":\"12\",\"ysrq\":\"2020-04-15T00:00:00\",\"bw1\":\"/\",\"czgg1\":\"/\",\"jgxs\":\"/\",\"wgjc1\":\"/\",\"ljqk1\":\"/\",\"ffcl1\":\"/\",\"bw2\":\"风机基础\",\"czgg\":\"𠃋50*4\",\"zd1\":\"140\",\"fsff\":\"垂直与水平混合敷设\",\"wgjc2\":\"良好\",\"ljqk2\":\"按规范焊接，搭接长度满足要求，焊 接质量良好\",\"ffcl2\":\"良好\",\"bw3\":\"风机基础\",\"czgg3\":\"𠃋50*4\",\"zd2\":\"14\",\"ljqk3\":\"良好\",\"dxkgd\":\"/\",\"jdtxh\":\"/\",\"ffcl3\":\"良好\",\"ID_\":\"789425898231356078\"}}]}","operation_status":1}`
	txevent := &TxEvent{
		TxKey:       "13083_0",
		TxID:        "b3106472b56715a1e82a4a0660f39d4913e354555e1151f662f2d3985195b3ea",
		DataIndex:   "HandlerForm_517689777039941719_1_2020-12-10 11:05:50",
		FuncName:    "HandlerForm",
		Status:      1,
		Data:        data,
		Time:        "2020-12-10 11:05:50",
		BlockHeight: 13083,
		NewTxId:     "",
	}
	err = db.Model(&TxEvent{}).Table("tx_event").Create(txevent).Error
	if err != nil {
		fmt.Printf("insert err : %v\n", err)
	}
}

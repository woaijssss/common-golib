package client

import (
	"testing"
	"time"
)

type PublishORM struct {
	ID                 int       `gorm:"id" json:"id"`
	CompanyID          int       `gorm:"company_id" json:"company_id"`
	EstateGroupID      int       `gorm:"estate_group_id" json:"estate_group_id"`
	EstateGroupName    string    `gorm:"estate_group_name" json:"estate_group_name"`
	ParkingID          int       `gorm:"parking_id" json:"parking_id"`
	ParkingName        string    `gorm:"parking_name" json:"parking_name"`
	ParkingSpaceID     int       `gorm:"parking_space_id" json:"parking_space_id"`
	ParkingSpaceName   string    `gorm:"parking_space_name" json:"parking_space_name"`
	ParkingAreaID      int       `gorm:"parking_area_id" json:"parking_area_id"`
	ParkingAreaName    string    `gorm:"parking_area_name" json:"parking_area_name"`
	OwnerUserID        int       `gorm:"owner_user_id" json:"owner_user_id"`
	OwnerUserName      string    `gorm:"owner_user_name" json:"owner_user_name"`
	OwnerUserMobile    string    `gorm:"owner_user_mobile" json:"owner_user_mobile"`
	TenantryUserID     int       `gorm:"tenantry_user_id" json:"tenantry_user_id"`
	TenantryUserName   string    `gorm:"tenantry_user_name" json:"tenantry_user_name"`
	TenantryUserMobile string    `gorm:"tenantry_user_mobile" json:"tenantry_user_mobile"`
	TenantryUserNo     string    `gorm:"tenantry_user_no" json:"tenantry_user_no"`
	PublishState       int       `gorm:"publish_state" json:"publish_state"`
	OrderState         int       `gorm:"order_state" json:"order_state"`
	Type               int       `gorm:"type" json:"type"`
	StartDate          time.Time `gorm:"start_date" json:"start_date"`
	EndDate            time.Time `gorm:"end_date" json:"end_date"`
	StartTime          time.Time `gorm:"start_time" json:"start_time"`
	EndTime            time.Time `gorm:"end_time" json:"end_time"`
	Rent               float64   `gorm:"rent" json:"rent"`
	TenantryTime       time.Time `gorm:"tenantry_time" json:"tenantry_time"`
	PaymentTIme        time.Time `gorm:"payment_time" json:"payment_time"`
	ConfirmTime        time.Time `gorm:"confirm_time" json:"confirm_time"`
	CreateBy           int       `gorm:"create_by" json:"create_by"`
	CreateTime         time.Time `gorm:"create_time" json:"create_time"`
}

func TestMysqlSetUp(T *testing.T) {
	//SetEnv("dev")
	//var p PublishORM
	//err := GetParkingDB().Table("fw_parking_publish").Where("id1 in (?) AND create_time <= ?", []int{1, 2}, time.Now()).Take(&p).Error
	//if err != nil {
	//	fmt.Println("================")
	//	T.Error()
	//}
	//
	//fmt.Println("===========================")
	//fmt.Println(p.StartTime)
	//fmt.Println("===========================")
}

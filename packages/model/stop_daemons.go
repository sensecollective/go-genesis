package model

type StopDaemon struct {
	StopTime int32 `gorm:"not null"`
}

func (sd *StopDaemon) TableName() string {
	return "stop_daemons"
}

func (sd *StopDaemon) Create() error {
	return DBConn.Create(sd).Error
}

func (sd *StopDaemon) Delete() error {
	return DBConn.Delete(&StopDaemon{}).Error
}

func (sd *StopDaemon) Get() error {
	return DBConn.First(sd).Error
}
package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewMerchantData() *MerchantData {
	return &MerchantData{}
}

type MerchantData struct {
}

func (s MerchantData) List() ([]*po.MerchantList, error) {
	list := make([]*po.MerchantList, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("MerchantData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s MerchantData) Info(id int64) (*po.MerchantList, error) {
	info := &po.MerchantList{}
	err := conn.GetMerchantDB().First(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("MerchantData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MerchantData) Add(merc *po.MerchantList) error {
	merc.Id = 0
	err := conn.GetMerchantDB().Model(&po.MerchantList{}).Where("id = ?", merc.Id).Create(merc).Error
	if err != nil {
		logrus.Errorf("MerchantData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MerchantData) Edit(merc *po.MerchantList) error {
	err := conn.GetMerchantDB().Model(&po.MerchantList{}).Where("id = ?", merc.Id).Updates(merc).Error
	if err != nil {
		logrus.Errorf("MerchantData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MerchantData) Delete(merc *po.MerchantList) error {
	err := conn.GetMerchantDB().Where("id = ?", merc.Id).Delete(merc).Error
	if err != nil {
		logrus.Errorf("MerchantData Delete Err: %s", err.Error())
		return err
	}
	return nil
}

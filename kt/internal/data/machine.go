// package data

// import (
// 	"context"
// 	"kt/internal/biz"

// 	"time"

// 	"github.com/go-kratos/kratos/v2/log"
// 	"gorm.io/gorm"
// )

// // var _ biz.MachineRepo = (*machineRepo)(nil)
// var _ biz.MachineRepo = (*machineRepo)(nil)

// type machineRepo struct {
// 	data *Data
// 	log  *log.Helper
// }

// type Machine struct {
// 	gorm.Model
// 	ID         string
// 	PropertyId string
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }

// // func (u *Machine) BeforeCreate(tx *gorm.DB) (err error) {
// // 	u.ID = ulid.GenerateUlid()

// // 	u.CreatedAt = time.Now().UTC()
// // 	u.UpdatedAt = time.Now().UTC()
// // 	return nil
// // }

// func NewMachineRepo(data *Data, logger log.Logger) biz.MachineRepo {
// 	return &machineRepo{
// 		data: data,
// 		log:  log.NewHelper(log.With(logger, "module", "Machine-service/data/machine")),
// 	}
// }

// // GetMachine implements biz.MachineRepo
// func (u *machineRepo) GetMachine(ctx context.Context, pId int32) (*biz.Machine, error) {
// 	var machines []Machine
// 	result := u.data.db.Find(&machines)
// 	return result, nil
// }

// func modelToResponse(machine Machine) biz.Machine {
// 	machineInfoRsp := biz.Machine{
// 		Id:         machine.ID,
// 		PropertyId: machine.PropertyId,
// 		CreatedAt:  machine.CreatedAt,
// 		UpdatedAt:  machine.UpdatedAt,
// 	}
// 	return machineInfoRsp
// }

package data

import (
	"context"
	"kt/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type machineRepo struct {
	data *Data
	log  *log.Helper
}

func NewMachineRepo(data *Data, logger log.Logger) biz.MachineRepo {
	return &machineRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data")),
	}
}

func (r *machineRepo) GetMachine(ctx context.Context, machineID string) (*biz.Machine, error) {
	var machine biz.Machine
	if err := r.data.db.Where("id = ?", machineID).First(&machine).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &machine, nil
}

package postgresrepository

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
	"gorm.io/gorm"
)

type ConsultaPostgresRepository struct {
	db *gorm.DB
}

func NewConsultaPostgresRepository(repo *gorm.DB) *ConsultaPostgresRepository {
	return &ConsultaPostgresRepository{
		db: repo,
	}
}

func (u *ConsultaPostgresRepository) CreateConsulta(ctx context.Context, consulta domain.Consulta) error {
	err := u.db.Create(&consulta).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *ConsultaPostgresRepository) FindByConsultaId(ctx context.Context, consulta_id string) (domain.Consulta, error) {

	var consulta domain.Consulta
	err := u.db.Where("id = ?", consulta_id).First(&consulta).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Consulta{}, nil
		}
		return domain.Consulta{}, err
	}
	if consulta.ID == "" {
		return domain.Consulta{}, nil
	}
	return consulta, nil
}

func (u *ConsultaPostgresRepository) GetTotalConsultas(ctx context.Context) (int64, error) {

	var totalConsultas int64
	err := u.db.Model(&domain.Consulta{}).Count(&totalConsultas).Error
	if err != nil {
		return 0, err
	}
	return totalConsultas, nil
}

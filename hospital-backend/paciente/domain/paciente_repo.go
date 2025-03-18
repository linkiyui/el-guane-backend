package domain

import "context"

type IPacienteRepository interface {
	CreatePaciente(ctx context.Context, paciente Paciente) error
	GetAllPacientes(ctx context.Context) ([]Paciente, error)
}

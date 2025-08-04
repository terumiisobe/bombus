package mapper

import (
	"bombus/domain"
	"bombus/dto"
)

func ToDTO(colmeia domain.Colmeia) dto.Colmeia {
	species := colmeia.Species.GetCommonName()
	status := colmeia.Status.String()
	startingDate := colmeia.StartingDate.Format("2006-01-02")

	return dto.Colmeia{
		ID:           &colmeia.ID,
		ColmeiaID:    colmeia.ColmeiaID,
		QRCode:       colmeia.QRCode,
		Species:      &species,
		StartingDate: &startingDate,
		Status:       &status,
	}
}

// ToDTOList converts slice of domain Colmeia to DTO slice
func ToDTOList(colmeias []domain.Colmeia) []dto.Colmeia {
	dtos := make([]dto.Colmeia, len(colmeias))
	for i, colmeia := range colmeias {
		dtos[i] = ToDTO(colmeia)
	}
	return dtos
}

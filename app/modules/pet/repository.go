package pet

import (
	"context"
	"golang-docker-demo/app/server/generated"
	"golang-docker-demo/db/generated/model"
	"golang-docker-demo/db/generated/query"
)

type Repository struct {
}

func NewRepository() RepositoryInterface {
	return &Repository{}
}

func (r *Repository) FindPetByStatus(ctx context.Context, tx *query.Query, params generated.FindPetsByStatusParams) (pets []*model.MPet, err error) {
	query := tx.MPet

	err = query.WithContext(ctx).Where(
		tx.MPet.Status.Eq(string(params.Status[0])),
	).Scan(&pets)

	if err != nil {
		return nil, err
	}

	return pets, nil
}

// AddPet implements RepositoryInterface.
func (r *Repository) AddPet(ctx context.Context, tx *query.Query, body *model.MPet) (pet *model.MPet, err error) {
	err = tx.WithContext(ctx).MPet.Create(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

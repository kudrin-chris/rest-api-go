package store

import "fileServer/internal/app/model"

type RecorderRepository struct {
	store *Store
}

func AddBD() error {
	r := model.Record{
		1,
		1,
		"str",
		"str",
		"str1",
		"str",
		"str",
		"str",
		"str",
		1,
		"str",
		"str",
		"str",
		"str",
		"str",
	}
	t := RecorderRepository{store: New(NewConfig())}
	p, err := t.Create(&r)
	p.Bit_ = "p"
	if err != nil {
		return nil
	}
	return nil
}
func (r *RecorderRepository) Create(rec *model.Record) (*model.Record, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO record (Id, N, Mqtt, Invid, Unit_guid, Msg_id, Text, Context, Class, Level, Area, Block, Type_, Bit_, Invert_bit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id",
		rec.Id,
		rec.N,
		rec.Mqtt,
		rec.Invid,
		rec.Unit_guid,
		rec.Msg_id,
		rec.Text,
		rec.Context,
		rec.Class,
		rec.Level,
		rec.Area,
		rec.Block,
		rec.Type_,
		rec.Bit_,
		rec.Invert_bit,
	).Scan(&rec.Id); err != nil {
		return nil, err
	}
	return rec, nil
}
func CreteInfoError() {

}

func (r *RecorderRepository) FindByUnitQuid(unit_guid string) (*model.Record, error) {

	return nil, nil
}

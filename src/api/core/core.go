package core

func (cc Core) InitDbTables() error{
	err := cc.db.CreateTables()
	if err != nil {
		return err
	}

	return nil
}

func NewCore() (*Core, error){
	db, err := NewDBBackend()
	if err != nil {
		return nil, err
	}

	return &Core{db:db}, nil
}

type Core struct {
	db *DBBackend
}
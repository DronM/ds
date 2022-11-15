package ds_test

import (
	"testing"
	"context"
	"fmt"
	
	"ds"
	"ds/pgds"
	
	"github.com/jackc/pgx/v4"
)

func TestPG(t *testing.T) {
	//get provider interface
	prov, err := ds.NewProvider("pg", "postgresql://tmnstroi:159753@192.168.1.77:5432/tmnstroi", nil, nil)
	if err != nil {
		panic(err)
	}
	
	//cast to pg
	pg := prov.(*pgds.PgProvider)
	//aquire connection
	pool_conn, conn_id, err := pg.GetPrimary()
	if err != nil {
		panic(err)
	}
	defer pg.Release(pool_conn, conn_id)
	
	q_res := pool_conn.Conn().QueryRow(context.Background(), "SELECT id, name FROM users LIMIT 1")
	res := struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}{}
	if err := q_res.Scan(&res.Id, &res.Name); err != nil && err != pgx.ErrNoRows {
		panic(err)
	}
	fmt.Println(res)
}


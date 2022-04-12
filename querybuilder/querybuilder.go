package querybuilder

import (
	"context"
	"errors"
	"strings"

	"github.com/takaaki-s/nextengine-sdk-go/entity"
	"github.com/takaaki-s/nextengine-sdk-go/nextengine"
)

func getAPI(e interface{}) (string, error) {
	uri := ""
	switch e.(type) {
	case *entity.LoginUser:
		uri = "/api_v1_login_user"
	case *entity.MasterShop:
		uri = "/api_v1_master_shop"
	case *entity.SystemPaymentMethod:
		uri = "/api_v1_system_paymentmethod"
	case *entity.ReceiveOrderBase:
		uri = "/api_v1_receiveorder_base"
	default:
		return "", errors.New("unknown entity")
	}

	return uri, nil
}

type QueryBuilder struct {
	fields string
	where  map[string]string
	cli    *nextengine.Client
}

func NewQueryBuilder(cli *nextengine.Client) *QueryBuilder {
	return &QueryBuilder{
		cli:   cli,
		where: map[string]string{},
	}
}

func (q *QueryBuilder) Select(args ...string) *QueryBuilder {
	q.fields = strings.Join(args, ",")
	return q
}

func (q *QueryBuilder) Where(target, operator, value string) *QueryBuilder {
	op := ""
	switch operator {
	case "=":
		op = "eq"
	case "!=":
		op = "neq"
	case ">=":
		op = "gte"
	case "<=":
		op = "lte"
	case "like":
		op = "like"
	}
	q.where[target+"-"+op] = value
	return q
}

func (q *QueryBuilder) Fetch(ctx context.Context, e nextengine.TokenGetter) error {
	uri, err := getAPI(e)
	if err != nil {
		return err
	}
	switch uri {
	case "/api_v1_login_user", "/api_v1_system_paymentmethod":
		uri += "/info"
	case "/api_v1_master_shop", "/api_v1_receiveorder_base":
		uri += "/search"
	}

	return q.do(ctx, e, uri)
}

func (q *QueryBuilder) Count(ctx context.Context, e nextengine.TokenGetter) error {
	uri, err := getAPI(e)
	if err != nil {
		return err
	}
	switch uri {
	default:
		uri += "/count"
	}
	return q.do(ctx, e, uri)
}

func (q *QueryBuilder) do(ctx context.Context, e nextengine.TokenGetter, uri string) error {

	params := make(map[string]string)
	if len(q.fields) > 0 {
		params["fields"] = q.fields
	}
	for k, v := range q.where {
		params[k] = v
	}

	err := q.cli.APIExecute(ctx, uri, params, e)
	if err != nil {
		return err
	}
	return nil
}

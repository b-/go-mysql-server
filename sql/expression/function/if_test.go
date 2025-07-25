// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package function

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/types"
)

func TestIf(t *testing.T) {
	testCases := []struct {
		expr     sql.Expression
		row      sql.Row
		expected interface{}
		type1    sql.Type
		type2    sql.Type
	}{
		{eq(lit(1, types.Int64), lit(1, types.Int64)), sql.Row{"a", "b"}, "a", types.Text, types.Text},
		{eq(lit(1, types.Int64), lit(0, types.Int64)), sql.Row{"a", "b"}, "b", types.Text, types.Text},
		{eq(lit(1, types.Int64), lit(1, types.Int64)), sql.Row{1, 2}, int64(1), types.Int64, types.Int64},
		{eq(lit(1, types.Int64), lit(0, types.Int64)), sql.Row{1, 2}, int64(2), types.Int64, types.Int64},
		{eq(lit(nil, types.Int64), lit(1, types.Int64)), sql.Row{"a", "b"}, "b", types.Text, types.Text},
		{eq(lit(1, types.Int64), lit(1, types.Int64)), sql.Row{nil, "b"}, nil, nil, types.Text},
	}

	for _, tc := range testCases {
		f := NewIf(
			tc.expr,
			expression.NewGetField(0, tc.type1, "true", true),
			expression.NewGetField(1, tc.type2, "false", true),
		)

		v, err := f.Eval(sql.NewEmptyContext(), tc.row)
		require.NoError(t, err)
		require.Equal(t, tc.expected, v)
	}
}

func eq(left, right sql.Expression) sql.Expression {
	return expression.NewEquals(left, right)
}

func lit(n interface{}, typ sql.Type) sql.Expression {
	return expression.NewLiteral(n, typ)
}

func col(idx int, typ sql.Type, db, table, col string) sql.Expression {
	return expression.NewGetFieldWithTable(idx, 0, typ, db, table, col, false)
}

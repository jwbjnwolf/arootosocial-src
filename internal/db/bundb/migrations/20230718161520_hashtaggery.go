// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			// Drop now unused columns from tags table.
			for _, column := range []string{
				"url",
				"first_seen_from_account_id",
				"last_status_at",
			} {
				if _, err := tx.
					NewDropColumn().
					Table("tags").
					Column(column).
					Exec(ctx); err != nil {
					return err
				}
			}

			// Index status_to_tags table properly.
			for index, columns := range map[string][]string{
				// Index for tag timeline paging.
				"status_to_tags_tag_timeline_idx": {"tag_id", "status_id"},
				// These indexes were only implicit
				// before, make them explicit now.
				"status_to_tags_tag_id_idx":    {"tag_id"},
				"status_to_tags_status_id_idx": {"status_id"},
			} {
				if _, err := tx.
					NewCreateIndex().
					Table("status_to_tags").
					Index(index).
					Column(columns...).
					Exec(ctx); err != nil {
					return err
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}

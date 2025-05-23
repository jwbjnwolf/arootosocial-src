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

//go:build !nowasmffmpeg && !nowasm

package ffmpeg

import (
	"context"
	"errors"

	"codeberg.org/gruf/go-ffmpreg/wasm"
)

// ffprobeRunner limits the number of
// ffprobe WebAssembly instances that
// may be concurrently running, in
// order to reduce memory usage.
var ffprobeRunner runner

// InitFfprobe precompiles the ffprobe WebAssembly source into memory and
// prepares the runner to only allow max given concurrent running instances.
func InitFfprobe(ctx context.Context, max int) error {
	ffprobeRunner.Init(max)
	return initWASM(ctx)
}

// Ffprobe runs the given arguments with an instance of ffprobe.
func Ffprobe(ctx context.Context, args Args) (uint32, error) {
	return ffprobeRunner.Run(ctx, func() (uint32, error) {

		// Load WASM rt and module.
		ffmpreg := ffmpreg.Load()
		if ffmpreg == nil {
			return 0, errors.New("wasm not initialized")
		}

		// Call into ffprobe.
		args.Name = "ffprobe"
		return wasm.Run(ctx,
			ffmpreg.run,
			ffmpreg.mod,
			args,
		)
	})
}

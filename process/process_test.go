// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package process

import (
	"testing"
)

// mkdir -p /tmp/test && chown -R nobody:nobody /tmp/test
// sudo go test
func TestInit(t *testing.T) {
	if err := Init("nobody nobody", "./", "/tmp/test/process_test.pid"); err != nil {
		t.Error(err)
	}
}

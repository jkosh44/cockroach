// Copyright 2017 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/testutils/buildutil"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
)

func TestNoLinkForbidden(t *testing.T) {
	defer leaktest.AfterTest(t)()
	// Verify that the cockroach floss binary doesn't depend on certain packages.
	buildutil.VerifyNoImports(t,
		"github.com/cockroachdb/cockroach/pkg/cmd/cockroach-oss",
		true,
		nil,
		[]string{
			"github.com/cockroachdb/cockroach/pkg/ccl",
			"github.com/cockroachdb/cockroach/pkg/ui/distccl",
		},
	)
}

/*
Copyright (c) 2021-2022 Progressive Casualty Insurance Company. All rights reserved.

Use of this source code is governed by an MIT license that can be found in
the LICENSE file at https://github.com/Progressive-Insurance/need-cla/blob/main/LICENSE.md
*/

package needcla

import (
	"context"
	"testing"

	"github.com/google/go-github/v43/github"
)

// TestUsesCLAAssistantActionNoWorkflowsDir ensures a repo without a
// .github/workflows directory returns (false, nil) rather than issuing a
// GetTree call with an empty SHA (which 404s) and surfacing a spurious
// ActionErr.
func TestUsesCLAAssistantActionNoWorkflowsDir(t *testing.T) {
	c := checker{
		branch: "main",
		repo:   "some-repo",
		owner:  "some-owner",
		// A tree with entries, but no .github/workflows directory.
		tree: &github.Tree{
			Entries: []*github.TreeEntry{
				{Path: github.String("README.md"), Type: github.String("blob"), SHA: github.String("abc")},
				{Path: github.String(".clabot"), Type: github.String("blob"), SHA: github.String("def")},
			},
		},
	}

	got, err := c.usesCLAAssistantAction(context.Background())
	if err != nil {
		t.Fatalf("expected no error for a repo without .github/workflows, got: %v", err)
	}
	if got {
		t.Errorf("expected Action=false when there is no .github/workflows dir, got true")
	}
}

package storage

import (
	"context"
	"gin_realworld/utils"
	"testing"
)

func TestListPopularTags(t *testing.T) {
	ctx := context.TODO()
	res, err := ListPopularTags(ctx)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("tags: %v\n", utils.JsonMarshal(res))
}

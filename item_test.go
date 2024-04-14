package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
    t.Run("Get Item", func(t *testing.T) {
        // Connect to the database
        item, err := getItem(db, "1")
        assert.Nil(t, err)
        assert.Equal(t, item.Name, "CIO")
        assert.Equal(t, item.ID, uint(1))
    })
}
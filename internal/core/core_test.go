package core

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestValidateApp(t *testing.T) {
	err := fx.ValidateApp(Core())
	require.NoError(t, err)
}

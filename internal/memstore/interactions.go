package memstore

import (
	"context"

	"github.com/Karitham/corde"
)

// IncrementInteractionCount increments the interaction count for the given channel/user combo
func (c Client) IncrementInteractionCount(ctx context.Context, channelID corde.Snowflake) error {
	return c.client.Incr(ctx, channelID.String()).Err()
}

// GetInteractionCount gets the interaction count for the given channel/user combo
func (c Client) GetInteractionCount(ctx context.Context, channelID corde.Snowflake) (int64, error) {
	return c.client.Get(ctx, channelID.String()).Int64()
}

// ResetInteractionCount resets the interaction count for the given channel/user combo
func (c Client) ResetInteractionCount(ctx context.Context, channelID corde.Snowflake) error {
	return c.client.Del(ctx, channelID.String()).Err()
}

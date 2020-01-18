package gauth

// ContextJWTMetaKey is the key within Context that holds the current UserID.
const ContextJWTMetaKey ContextKeyType = iota

// ContextKeyType is the key type for Contexts.
type ContextKeyType int

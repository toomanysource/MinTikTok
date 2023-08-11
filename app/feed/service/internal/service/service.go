package service

import "github.com/google/wire"

// ProviderSet is service providers.
// var ProviderSet = wire.NewSet(NewFeedService, NewUserServiceClient)
var ProviderSet = wire.NewSet(NewFeedService)
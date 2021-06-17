package internal

import "github.com/rs/zerolog/log"

var Logger = log.With().Str("revision", ShortRevision()).Logger()

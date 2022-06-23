package mparser

import (
	"github.com/Smaug6739/mparser/internal/logger"
	"github.com/Smaug6739/mparser/preprocessor"
	"github.com/Smaug6739/mparser/tokenizer"
)

// Execute preprocessor and tokenizer
func Tokenize(str string) *preprocessor.Markdown {
	log := logger.New()
	log.Info("START PREPROCESSOR.")
	markdown := preprocessor.New(str)
	log.Info("END PREPROCESSOR.")
	log.Info("START TOKENIZER.")
	tokenizer.New(markdown)
	log.Info("END TOKENIZER.")
	return markdown
}

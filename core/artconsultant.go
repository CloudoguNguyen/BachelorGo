package core

import (
	"github.com/cloudogu/BachelorGo/service"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

const (
	artSurrealism      = "surrealism"
	artComplex         = "complex art"
	artJapanese        = "japanese art"
	artNeutral         = "neutral_art"
	artNatural         = "natural_art"
	artRepresentative  = "representative art"
	artImpression      = "impressionism"
	artTradition       = "traditional art"
	artKubism          = "kubism"
	artNegEmotion      = "negatively emotional art"
	artPop             = "pop-art"
	artAbstract        = "abstract art"
	highIntensity      = "high"
	lowIntensity       = "low"
	middleIntensity    = "middle"
	artConsultantToken = "1fedc8b90ea54efc652b6a42c82de9f2"
)

type ArtConsultant struct {
	recastClient *service.RecastClient
}

func NewArtConsultant() *ArtConsultant {

	recast := service.NewRecastClient(artConsultantToken)

	return &ArtConsultant{recast}
}

func (ac *ArtConsultant) GetResponse(message string, conversationID string, watsonPI service.WatsonPI) (string, error) {

	if message == profile_not_valid {

		message, err := ac.getToKnowUser(conversationID)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get get to know user")
		}

		return message, nil
	}

	response := ""

	intent, err := ac.getIntent(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get intent")
	}

	if intent == "ask-art" {
		response = ac.recommendArt(watsonPI)
	}

	if response == "" {
		response = "We don't know what you want"
	}

	return response, nil
}

func (ac *ArtConsultant) getToKnowUser(conversationID string) (string, error) {

	reply, err := ac.recastClient.GetReplies(profile_not_valid, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply from %s", profile_not_valid)
	}

	return reply, nil

}

func getIntensity(value int) string {
	if value > 65 && value <= 100 {
		return highIntensity
	} else if value <= 65 && value >= 34 {
		return middleIntensity
	}

	return lowIntensity
}

func (ac *ArtConsultant) recommendArt(watsonPI service.WatsonPI) string {

	recommendableArts := ac.getRecommendableArts(watsonPI)

	if len(recommendableArts) > 0 {
		response := "You might like this direction of art: "

		for _, art := range recommendableArts {
			response += art + ", "
		}

		return response
	}

	return "We can't find any art which suit you"

}

func (ac *ArtConsultant) getRecommendableArts(watsonPI service.WatsonPI) []string {
	recommendableArts := []string{}

	if getIntensity(watsonPI.Openness()) == highIntensity {
		recommendableArts = append(recommendableArts, artSurrealism, artComplex, artJapanese)
	} else if getIntensity(watsonPI.Openness()) == lowIntensity {
		recommendableArts = append(recommendableArts, artNeutral, artNatural)
	}

	if getIntensity(watsonPI.Conscientiousness()) == highIntensity {
		recommendableArts = append(recommendableArts, artRepresentative)
	} else if getIntensity(watsonPI.Conscientiousness()) == lowIntensity {
		recommendableArts = append(recommendableArts, artImpression, artTradition)
	}

	if getIntensity(watsonPI.Extraversion()) == highIntensity {
		recommendableArts = append(recommendableArts, artKubism)
	}

	if getIntensity(watsonPI.Neuroticism()) == highIntensity {
		recommendableArts = append(recommendableArts, artNegEmotion, artPop, artAbstract)
	}

	return recommendableArts
}

func (ac *ArtConsultant) getIntent(message string, conversationID string) (string, error) {
	intent, err := ac.recastClient.GetIntent(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get intent from message %s", message)
	}

	return intent.Slug, err

}

func (ac *ArtConsultant) GetNewRandomConversationID() string {

	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

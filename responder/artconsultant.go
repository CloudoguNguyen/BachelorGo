package responder

import (
	"github.com/BachelorGo/service"
	"github.com/pkg/errors"
)

const (
	artSurrealism     = "surrealism"
	artComplex        = "complex art"
	artJapanese       = "japanese art"
	artNeutral        = "neutral art"
	artNatural        = "natural art"
	artRepresentative = "representative art"
	artImpression     = "impressionism"
	artTradition      = "traditional art"
	artKubism         = "kubism"
	artNegEmotion     = "negatively emotional art"
	artPop            = "pop-art"
	artAbstract       = "abstract art"

	highIntensity   = "high"
	lowIntensity    = "low"
	middleIntensity = "middle"

	artConsultantToken = "8fe1499c88ff54d49cc6ce5a8c549f28"
	ProfileNotValid    = "profile not valid"
)

type ArtConsultant struct {
	recastClient   *service.RecastClient
	isProfileKnown map[string]bool
}

func NewArtConsultant() *ArtConsultant {

	recast := service.NewRecastClient(artConsultantToken)
	isProfileKnown := make(map[string]bool)

	return &ArtConsultant{recast, isProfileKnown}
}

/**
1. If message = profile_not_valid -> A conversation for collection user info will be started.
This will be repeated until a user profile is created
2. Get intent of the message
3. If the intent = ask-art -> response is art recommendation based on the user profile
4. If intent = unknown, this information will be the response
*/
func (ac *ArtConsultant) GetResponse(message string, conversationID string, profile *service.UserProfile) (string, error) {

	response := ""

	if message == ProfileNotValid {

		response, err := ac.getToKnowUser(conversationID)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get to know user")
		}

		return response, nil
	}

	intent, err := ac.getIntent(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get intent")
	}

	if intent == "ask-art" {
		response = ac.recommendArt(*profile)
	}

	if response == "" {
		if ac.isProfileKnown[conversationID] == false {

			response = "We have enough information about you now. Please tell us what you want"
			ac.isProfileKnown[conversationID] = true

		} else {
			response = "We don't know what you want"
		}
	}

	return response, nil
}

func (ac *ArtConsultant) getToKnowUser(conversationID string) (string, error) {

	reply, err := ac.recastClient.GetReply(ProfileNotValid, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply from %s", ProfileNotValid)
	}

	return reply, nil

}

func getIntensity(value int) string {
	if value > 66 && value <= 100 {
		return highIntensity
	} else if value <= 66 && value >= 34 {
		return middleIntensity
	}

	return lowIntensity
}

func (ac *ArtConsultant) recommendArt(profile service.UserProfile) string {

	matchingArts := ac.getMatchingArts(profile)

	if len(matchingArts) > 0 {
		response := "You might like this direction of art: "

		for _, art := range matchingArts {
			response += "\n " + art
		}

		return response
	}

	return "We can't find any art which suit you"

}

func (ac *ArtConsultant) getMatchingArts(profile service.UserProfile) []string {
	matchingArts := []string{}

	// if openness's intensity is high
	if getIntensity(profile.Openness()) == highIntensity {
		matchingArts = append(matchingArts, artSurrealism, artComplex, artJapanese)
	} else if getIntensity(profile.Openness()) == lowIntensity {
		matchingArts = append(matchingArts, artNeutral, artNatural)
	}

	if getIntensity(profile.Conscientiousness()) == highIntensity {
		matchingArts = append(matchingArts, artRepresentative)
	} else if getIntensity(profile.Conscientiousness()) == lowIntensity {
		matchingArts = append(matchingArts, artImpression, artTradition)
	}

	if getIntensity(profile.Extraversion()) == highIntensity {
		matchingArts = append(matchingArts, artKubism)
	}

	if getIntensity(profile.Neuroticism()) == highIntensity {
		matchingArts = append(matchingArts, artNegEmotion, artPop, artAbstract)
	}

	return matchingArts
}

func (ac *ArtConsultant) getIntent(message string, conversationID string) (string, error) {
	intent, err := ac.recastClient.GetIntent(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get intent from message %s", message)
	}

	// return the intent only if recast is more than 94% sure about the correctness of intent
	if intent.Confidence >= 0.94 {
		return intent.Slug, nil
	}

	return "", nil

}

package core

import "time"

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type Payload struct {
	Message        Message `json:"message"`
	ConversationID string  `json:"conversation_id"`
}

type Conversation struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Memory   struct {
	} `json:"memory"`
	Skill           string `json:"skill"`
	SkillOccurences int    `json:"skill_occurences"`
}

type RecastResponse struct {
	Results struct {
		Messages     []Message    `json:"messages"`
		Conversation Conversation `json:"conversation"`
		Nlp          struct {
			UUID      string        `json:"uuid"`
			Source    string        `json:"source"`
			Intents   []interface{} `json:"intents"`
			Act       string        `json:"act"`
			Type      interface{}   `json:"type"`
			Sentiment string        `json:"sentiment"`
			Entities  struct {
				Datetime []struct {
					Formatted  string    `json:"formatted"`
					Iso        time.Time `json:"iso"`
					Accuracy   string    `json:"accuracy"`
					Chronology string    `json:"chronology"`
					State      string    `json:"state"`
					Raw        string    `json:"raw"`
					Confidence float64   `json:"confidence"`
				} `json:"datetime"`
			} `json:"entities"`
			Language           string    `json:"language"`
			ProcessingLanguage string    `json:"processing_language"`
			Version            string    `json:"version"`
			Timestamp          time.Time `json:"timestamp"`
			Status             int       `json:"status"`
		} `json:"nlp"`
	} `json:"results"`
	Message string `json:"message"`
}

type Profile struct {
	Tree struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Children []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Children []struct {
				ID         string  `json:"id"`
				Name       string  `json:"name"`
				Category   string  `json:"category"`
				Percentage float64 `json:"percentage"`
				Children   []struct {
					ID            string  `json:"id"`
					Name          string  `json:"name"`
					Category      string  `json:"category"`
					Percentage    float64 `json:"percentage"`
					SamplingError float64 `json:"sampling_error"`
					Children      []struct {
						ID            string  `json:"id"`
						Name          string  `json:"name"`
						Category      string  `json:"category"`
						Percentage    float64 `json:"percentage"`
						SamplingError float64 `json:"sampling_error"`
					} `json:"children"`
				} `json:"children"`
			} `json:"children"`
		} `json:"children"`
	} `json:"tree"`
	ID               string `json:"id"`
	Source           string `json:"source"`
	WordCount        int    `json:"word_count"`
	WordCountMessage string `json:"word_count_message"`
	ProcessedLang    string `json:"processed_lang"`
}

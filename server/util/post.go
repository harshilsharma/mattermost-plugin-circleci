package util

import (
	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/chetanyakan/mattermost-plugin-circleci/server/config"
)

func BaseSlackAttachment() *model.SlackAttachment {
	return &model.SlackAttachment{
		AuthorIcon: config.BotIconURL,
		AuthorName: config.BotDisplayName,
		Color:      "#7FC1EE",
		ThumbURL:   config.BotThumbnail,
	}
}

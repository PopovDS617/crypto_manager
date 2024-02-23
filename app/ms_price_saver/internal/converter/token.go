package converter

import "ms_price_saver/internal/model"

func ToRepoFromMessageQueue(incomingMessage *model.MessageQueueTokenDataList) []model.RepoTokenData {

	result := make([]model.RepoTokenData, 0, len(incomingMessage.TokenData))

	for _, v := range incomingMessage.TokenData {

		tokenData := model.RepoTokenData{TokenData: v, Timestamp: incomingMessage.Timestamp}
		result = append(result, tokenData)

	}

	return result
}

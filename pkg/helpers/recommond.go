package helpers

import (
	"math"
	"sort"
)

type BaseModel struct {
	ID int
}

type ResourceModel struct {
	ID          int
	Name        string
	Description string
	Url         string
	Tag         string
	Score       float64
}

type UserActionModel struct {
	ID         int
	UserID     int
	ResourceID int
	ActionType string
}

/*这段代码实现了一个基于用户交互历史和资源特征计算相似度的推荐系统。具体而言，它的过程分为以下几步：

1. 获取用户最近交互的资源：遍历给定的用户行为列表，构建一个包含最近访问过的资源ID的集合。

2. 计算资源之间的相似度：遍历所有资源，针对每个资源，提取其特征（如名称、描述、标签等），然后与用户访问过的资源进行比较，计算它们之间的相似度得分。相似度得分的计算方法是：统计两个资源特征的交集大小，并将其除以两者特征集合大小的平方根。这个值越大，则说明两个资源越相似。

3. 推荐资源，并按评分排序：遍历所有资源，对于每个资源，首先排除那些最近被用户访问过的资源；然后，针对每个用户对该资源的交互历史，给出一个初始的评分；最后，乘上前面计算得到的相似度得分，得到最终的评分。根据评分大小将所有资源排序，返回前N个资源作为推荐结果。

需要注意的是，在这个方法中，关于资源特征的处理方式比较简单，只是将名称、描述和标签拼接在一起后以字符串的方式进行比较。如果使用更加复杂的特征提取方法，如人工标注、文本分析或深度学习等，可能会得到更好的推荐效果。



*/

func GetRecommendations(userActions []UserActionModel, resources []ResourceModel, numRecs int) []ResourceModel {
	// 遍历用户与资源之间的交互记录，获取用户最近交互的资源，并保存在字典 recentResources 中。
	recentResources := make(map[int]bool)
	for _, action := range userActions {
		if _, ok := recentResources[action.ResourceID]; !ok {
			recentResources[action.ResourceID] = true
		}
	}

	// 对每个资源计算它们之间的相似度，并保存在字典 simScores 中。
	simScores := make(map[int]float64)
	resourceFeatures := make(map[int]map[string]bool) // 资源的特征，例如标签、描述等
	for _, r := range resources {
		resourceFeatures[r.ID] = make(map[string]bool)
		for _, f := range []string{r.Name, r.Description, r.Tag} {
			resourceFeatures[r.ID][f] = true
		}

		// 计算相似度
		simScore := 0.0
		for id := range recentResources {
			if id == r.ID {
				continue
			}
			featureCount := 0
			for f := range resourceFeatures[id] {
				if resourceFeatures[r.ID][f] {
					featureCount++
				}
			}
			simScore += float64(featureCount) / math.Sqrt(float64(len(resourceFeatures[id]))*float64(len(resourceFeatures[r.ID])))
		}
		simScores[r.ID] = simScore
	}

	// 对所有资源进行评分，并按照得分排序，返回前 numRecs 个推荐结果。
	recommendations := make([]ResourceModel, 0)
	for _, r := range resources {
		if _, ok := recentResources[r.ID]; ok {
			continue // 排除最近用户交互过的资源
		}
		score := 0.0
		for _, action := range userActions {
			if action.ResourceID == r.ID {
				score += 1.0 // 用户曾经对该资源进行过交互，给予较高评分
			}
		}
		score *= simScores[r.ID] // 加权得分
		recommendations = append(recommendations, ResourceModel{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Url:         r.Url,
			Tag:         r.Tag,
			Score:       score,
		})
	}
	sort.Slice(recommendations, func(i, j int) bool { return recommendations[i].Score > recommendations[j].Score })
	if numRecs < len(recommendations) {
		return recommendations[:numRecs]
	} else {
		return recommendations
	}
}

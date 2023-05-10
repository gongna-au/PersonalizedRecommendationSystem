// 计算两个向量之间的余弦相似度
package helpers

import (
	"math"
	"sort"
)

func cosineSimilarity(vec1, vec2 []float64) float64 {
	dotProduct := 0.0
	norm1 := 0.0
	norm2 := 0.0
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
		norm1 += vec1[i] * vec1[i]
		norm2 += vec2[i] * vec2[i]
	}
	return dotProduct / (math.Sqrt(norm1) * math.Sqrt(norm2))
}

// 计算用户之间的相似度矩阵
func computeUserSimilarityMatrix(data map[int]map[int]float64) map[int]map[int]float64 {
	simMatrix := make(map[int]map[int]float64)
	for uid1, items1 := range data {
		simMatrix[uid1] = make(map[int]float64)
		for uid2, items2 := range data {
			if uid1 != uid2 {
				vec1 := make([]float64, len(items1))
				vec2 := make([]float64, len(items1))
				k := 0
				for iid := range items1 {
					if rating2, ok := items2[iid]; ok {
						vec1[k] = items1[iid]
						vec2[k] = rating2
						k++
					}
				}
				vec1 = vec1[:k]
				vec2 = vec2[:k]
				simMatrix[uid1][uid2] = cosineSimilarity(vec1, vec2)
			}
		}
	}
	return simMatrix
}

// 计算物品之间的相似度矩阵
func computeItemSimilarityMatrix(data map[int]map[int]float64) map[int]map[int]float64 {
	simMatrix := make(map[int]map[int]float64)
	for iid1, users1 := range transposeData(data) {
		simMatrix[iid1] = make(map[int]float64)
		for iid2, users2 := range transposeData(data) {
			if iid1 != iid2 {
				vec1 := make([]float64, len(users1))
				vec2 := make([]float64, len(users1))
				k := 0
				for uid := range users1 {
					if rating2, ok := users2[uid]; ok {
						vec1[k] = users1[uid]
						vec2[k] = rating2
						k++
					}
				}
				vec1 = vec1[:k]
				vec2 = vec2[:k]
				simMatrix[iid1][iid2] = cosineSimilarity(vec1, vec2)
			}
		}
	}
	return simMatrix
}

// 转置数据集，将物品 ID 和用户 ID 进行交换
func transposeData(data map[int]map[int]float64) map[int]map[int]float64 {
	transposedData := make(map[int]map[int]float64)
	for uid, items := range data {
		for iid, rating := range items {
			if _, ok := transposedData[iid]; !ok {
				transposedData[iid] = make(map[int]float64)
			}
			transposedData[iid][uid] = rating
		}
	}
	return transposedData
}

// 为指定用户进行物品推荐
func recommendItems(uid int, data map[int]map[int]float64, simMatrix map[int]map[int]float64, K, N int) []int {
	if _, ok := data[uid]; !ok {
		return nil
	}
	neighbors := topKSimilarUsers(uid, data, simMatrix, K)
	itemScores := make(map[int]float64)
	for _, neighbor := range neighbors {
		for iid, rating := range data[neighbor] {
			if _, ok := data[uid][iid]; !ok {
				itemScores[iid] += simMatrix[uid][neighbor] * rating
			}
		}
	}
	items := rankByScore(itemScores)
	if len(items) > N {
		items = items[:N]
	}
	results := make([]int, len(items))
	for i, item := range items {
		results[i] = item.key
	}
	return results
}

// 找到与指定用户相似度最高的 K 个用户
func topKSimilarUsers(uid int, data map[int]map[int]float64, simMatrix map[int]map[int]float64, K int) []int {
	sims := make([]struct {
		uid int
		sim float64
	}, 0, len(data)-1)
	for otherUID, _ := range data {
		if otherUID != uid {
			sims = append(sims, struct {
				uid int
				sim float64
			}{otherUID, simMatrix[uid][otherUID]})
		}
	}
	sort.Slice(sims, func(i, j int) bool {
		return sims[i].sim > sims[j].sim
	})
	if len(sims) > K {
		sims = sims[:K]
	}
	results := make([]int, K)
	for i, item := range sims {
		results[i] = item.uid
	}
	return results
}

// 按照分数对物品进行排序，返回排名结果和分数
func rankByScore(scores map[int]float64) []struct {
	key   int
	value float64
} {
	items := make([]struct {
		key   int
		value float64
	}, len(scores))
	i := 0
	for k, v := range scores {
		items[i] = struct {
			key   int
			value float64
		}{k, v}
		i++
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})
	return items
}

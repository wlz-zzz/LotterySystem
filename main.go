package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Participant 代表一个参与者，包含ID和权重
type Participant struct {
	ID     int
	Weight int
}

// LotterySystem 代表抽奖系统，包含参与者列表
type LotterySystem struct {
	Participants []Participant
}

// NewLotterySystem 初始化抽奖系统，创建指定数量的参与者
func NewLotterySystem(n int) *LotterySystem {
	participants := make([]Participant, n)
	for i := 0; i < n; i++ {
		participants[i] = Participant{
			ID:     i + 1,
			Weight: 1, // 初始权重设为1
		}
	}
	return &LotterySystem{
		Participants: participants,
	}
}

// DrawWinners 随机抽取中奖者，根据权重选择并更新未中奖者的权重
func (ls *LotterySystem) DrawWinners(numWinners int) []Participant {
	var winners []Participant
	totalWeight := ls.getTotalWeight() // 计算当前总权重
	rand.Seed(time.Now().UnixNano())   // 设置随机种子

	// 记录已中奖的参与者，避免重复中奖
	winnerIDs := make(map[int]bool)

	for len(winners) < numWinners {
		randomValue := rand.Intn(totalWeight) // 生成随机值
		cumulativeWeight := 0

		for _, p := range ls.Participants {
			cumulativeWeight += p.Weight
			// 随机值落在当前参与者的权重区间，且参与者未中奖
			if randomValue < cumulativeWeight && !winnerIDs[p.ID] {
				winners = append(winners, p)
				winnerIDs[p.ID] = true  // 标记为已中奖
				totalWeight -= p.Weight // 更新总权重，避免重复中奖
				break
			}
		}
	}

	// 增加未中奖者的权重
	for i := range ls.Participants {
		if !winnerIDs[ls.Participants[i].ID] {
			ls.Participants[i].Weight++ // 增加未中奖者的权重
		}
	}

	return winners
}

// getTotalWeight 计算所有参与者的总权重
func (ls *LotterySystem) getTotalWeight() int {
	total := 0
	for _, p := range ls.Participants {
		total += p.Weight
	}
	return total
}

func main() {
	numParticipants := 10 // 参与者人数
	numWinners := 3       // 每次抽取的中奖人数
	lotterySystem := NewLotterySystem(numParticipants)
	for i := 0; i < 3; i++ {
		lottery(numParticipants, numWinners, lotterySystem)
	}
}

func lottery(numParticipants int, numWinners int, lotterySystem *LotterySystem) {
	fmt.Printf(
		"开始抽奖，%d 名参与者中抽取 %d 名中奖者。\n", numParticipants, numWinners,
	)

	winners := lotterySystem.DrawWinners(numWinners)

	fmt.Println("中奖者列表:")
	for _, winner := range winners {
		fmt.Printf("参与者ID: %d\n", winner.ID)
	}

	fmt.Println("\n下一轮参与者及其更新后的权重:")
	for _, p := range lotterySystem.Participants {
		fmt.Printf("参与者ID: %d, 权重: %d\n", p.ID, p.Weight)
	}
	fmt.Printf("抽奖结束\n\n")
}


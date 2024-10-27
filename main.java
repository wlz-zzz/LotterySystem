import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Random;
import java.util.Set;

class Participant {
    int id;
    int weight;

    Participant(int id, int weight) {
        this.id = id;
        this.weight = weight;
    }
}

class LotterySystem {
    List<Participant> participants;

    public LotterySystem(int n) {
        participants = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            participants.add(new Participant(i, 1)); // 初始权重设为1
        }
    }

    // 抽取指定数量的中奖者，并更新未中奖者的权重
    public List<Participant> drawWinners(int numWinners) {
        List<Participant> winners = new ArrayList<>();
        int totalWeight = getTotalWeight();
        Random random = new Random();

        // 记录已中奖的参与者，避免重复中奖
        Set<Integer> winnerIds = new HashSet<>();

        while (winners.size() < numWinners) {
            int randomValue = random.nextInt(totalWeight);
            int cumulativeWeight = 0;

            for (Participant p : participants) {
                cumulativeWeight += p.weight;
                if (randomValue < cumulativeWeight && !winnerIds.contains(p.id)) {
                    winners.add(p);
                    winnerIds.add(p.id);  // 标记为已中奖
                    totalWeight -= p.weight; // 更新总权重
                    break;
                }
            }
        }

        // 增加未中奖者的权重
        for (Participant p : participants) {
            if (!winnerIds.contains(p.id)) {
                p.weight++;
            }
        }

        return winners;
    }

    // 计算所有参与者的总权重
    private int getTotalWeight() {
        int total = 0;
        for (Participant p : participants) {
            total += p.weight;
        }
        return total;
    }
}

public class Main {
    public static void main(String[] args) {
        int numParticipants = 10; // 参与者人数
        int numWinners = 3;       // 每次抽取的中奖人数
        LotterySystem lotterySystem = new LotterySystem(numParticipants);

        for (int i = 0; i < 3; i++) {  // 模拟多轮抽奖
            runLottery(numParticipants, numWinners, lotterySystem);
        }
    }

    public static void runLottery(int numParticipants, int numWinners, LotterySystem lotterySystem) {
        System.out.printf("开始抽奖，%d 名参与者中抽取 %d 名中奖者。\n", numParticipants, numWinners);

        List<Participant> winners = lotterySystem.drawWinners(numWinners);

        System.out.println("中奖者列表:");
        for (Participant winner : winners) {
            System.out.printf("参与者ID: %d\n", winner.id);
        }

        System.out.println("\n下一轮参与者及其更新后的权重:");
        for (Participant p : lotterySystem.participants) {
            System.out.printf("参与者ID: %d, 权重: %d\n", p.id, p.weight);
        }
        System.out.println("抽奖结束\n");
    }
}


package GA

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	groupSize      int      //种群大小
	chromosomeSize int      //染色体长度
	selectRand     float64  //轮盘选择概率
	crossRand      float64  //交叉概率
	mutationRand   float64  //变异概率
	group          []Person //种群
	bestPerson     Person   //当前最好的个体
	r              *rand.Rand
)

//Person 个体
type Person struct {
	chromosome []int   //染色体
	value      float64 //适应值
}

//Init 初始化函数
//初始化设置种群大小、轮盘选择概率、交叉概率已经变异的概率
func Init(GroupSize, ChromosomeSize int, SelectRand, CrossRand, MutationRand float64) {
	groupSize = GroupSize
	crossRand = CrossRand
	selectRand = SelectRand
	mutationRand = MutationRand
	chromosomeSize = ChromosomeSize
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	bestPerson.chromosome = make([]int, chromosomeSize)
}

//InitGroup 初始化种群
//根据种群大小随机产生一些个体填充
func InitGroup() {
	group = make([]Person, groupSize)
	for i := 0; i < groupSize; i++ {
		group[i].chromosome = make([]int, chromosomeSize)
		for j := 0; j < chromosomeSize; j++ {
			if r.Float64() > selectRand {
				group[i].chromosome[j] = 1
			}
		}
	}
}

//Fitness 计算适应值
func Fitness(person Person) float64 {
	x := decode(person)
	return x + 10*math.Sin(5*x) + 7*math.Cos(4*x)
}

//解码
func decode(person Person) float64 {
	var sum float64
	//解码
	for i := 0; i < chromosomeSize; i++ {
		//二进制染色体转十进制值
		if person.chromosome[i] == 1 {
			sum = sum + math.Pow(2.0, float64(i))
		}
	}
	return sum * 9 / (math.Pow(2.0, 14.0) - 1)
}

//Select 选择
func Select() {
	newGroup := make([]Person, groupSize)
	for i := 0; i < groupSize; i++ {
		newGroup[i].chromosome = make([]int, chromosomeSize)
		rnd := r.Float64()

	A:
		for j := 0; j < groupSize; j++ {
			if group[j].value > rnd*bestPerson.value {
				copy(newGroup[i].chromosome, group[j].chromosome)
				break A
			}
			if j == groupSize-1 {
				copy(newGroup[i].chromosome, bestPerson.chromosome)
			}
		}
	}
	group = newGroup
	newGroup = nil
}

//Cross 交叉
func Cross() {
	for i := 0; i < groupSize; i = i + 2 {
		if r.Float64() < crossRand {
			crossPosition := r.Intn(chromosomeSize - 1)
			if crossPosition == 0 || crossPosition == 1 {
				continue
			}
			//交叉
			for j := crossPosition; j < chromosomeSize; j++ {
				tmp := group[i].chromosome[j]
				group[i].chromosome[j] = group[i+1].chromosome[j]
				group[i+1].chromosome[j] = tmp
			}
		}
	}
}

//Mutation 变异
func Mutation() {
	for i := 0; i < groupSize; i++ {
		if r.Float64() < mutationRand {
			mutationPosition := r.Intn(chromosomeSize - 1)

			//单点变异
			if group[i].chromosome[mutationPosition] == 0 {
				group[i].chromosome[mutationPosition] = 1
			} else {
				group[i].chromosome[mutationPosition] = 0
			}
		}
	}
}

//GA 遗传算法
func GA() {
	//初始化
	Init(100, 14, 0.5, 0.6, 0.05)
	//初始化种群
	InitGroup()

	//遗传循环
	for i := 0; i < 1000; i++ {

		//计算适应值
		for j := 0; j < groupSize; j++ {
			group[j].value = Fitness(group[j])

			//保存当前最好的个体
			if group[j].value > bestPerson.value {
				copy(bestPerson.chromosome, group[j].chromosome)
				bestPerson.value = group[j].value
			}
		}

		fmt.Println("第", i, "代最好个体：", bestPerson.value, " ", decode(bestPerson))

		//选择
		Select()

		//交叉
		Cross()

		//变异
		Mutation()
	}
}

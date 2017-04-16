package GA

var (
	groupSize    int //种群大小
	selectRand   int //轮盘选择概率
	crossRand    int //交叉概率
	mutationRand int //变异概率
)

//Person 个体
type Person struct {
	chromosomes []int //染色体
}

//Init 初始化函数
//初始化设置种群大小、轮盘选择概率、交叉概率已经变异的概率
func Init(GroupSize, SelectRand, CrossRand, MutationRand int) {
	groupSize = GroupSize
	crossRand = CrossRand
	selectRand = SelectRand
	mutationRand = MutationRand
}

//InitGroup 初始化种群
//根据种群大小随机产生一些个体填充
func InitGroup() {

}

//Fitness 计算适应值
func Fitness() {

}

//Select 选择
func Select() {

}

//Cross 交叉
func Cross() {

}

//Mutation 变异
func Mutation() {

}

//GA 遗传算法
func GA() {

}

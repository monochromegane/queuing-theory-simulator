package queuing_theory_simulator

type Model interface {
	Progress(int) (int, int, int, []int)
	History() [][]int
}
